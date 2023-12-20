package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/rewards/types"
	subscriptionTypes "github.com/lavanet/lava/x/subscription/types"
)

func (k Keeper) AggregateRewards(ctx sdk.Context, provider, chainid string, adjustmentDenom uint64, rewards math.Int) {
	index := types.BasePayIndex{Provider: provider, ChainID: chainid}
	basepay, found := k.getBasePay(ctx, index)
	adjustedPay := sdk.NewDecFromInt(rewards).QuoInt64(int64(adjustmentDenom))
	if !found {
		basepay = types.BasePay{Total: rewards, TotalAdjusted: adjustedPay}
	} else {
		basepay.Total = basepay.Total.Add(rewards)
		basepay.TotalAdjusted = basepay.TotalAdjusted.Add(adjustedPay)
	}

	k.setBasePay(ctx, index, basepay)
}

func (k Keeper) distributeMonthlyBonusRewards(ctx sdk.Context) {
	total := k.TotalPoolTokens(ctx, types.ProviderRewardsDistributionPool)
	totalRewarded := sdk.ZeroInt()
	totalToSendRewarded := sdk.ZeroInt()
	specs := k.specEmissionParts(ctx)
	for _, spec := range specs {
		basepays, totalbasepay := k.specProvidersBasePay(ctx, spec.ChainID)
		if totalbasepay.IsZero() {
			continue
		}

		specTotalPayout := k.specTotalPayout(ctx, total, sdk.NewDecFromInt(totalbasepay), spec)
		for _, basepay := range basepays {
			reward := specTotalPayout.Mul(basepay.TotalAdjusted).QuoInt(totalbasepay).TruncateInt()
			totalRewarded = totalRewarded.Add(reward)
			if totalRewarded.GT(total) {
				utils.LavaFormatError("provider rewards are larger than the distribution pool balance", nil,
					utils.LogAttr("distribution_pool_balance", total.String()),
					utils.LogAttr("provider_reward", totalRewarded.String()))
				return
			}
			// now give the reward the provider contributor and delegators
			providerAddr, err := sdk.AccAddressFromBech32(basepay.Provider)
			if err != nil {
				continue
			}
			_, totalSent, err := k.dualstakingKeeper.RewardProvidersAndDelegators(ctx, providerAddr, basepay.ChainID, reward, string(types.ProviderRewardsDistributionPool), false, false, false)
			if err != nil {
				utils.LavaFormatError("Failed to send bonus rewards to provider", err, utils.LogAttr("provider", basepay.Provider))
			}
			totalToSendRewarded = totalToSendRewarded.Add(totalSent)
		}
	}
	k.removeAllBasePay(ctx)
	if !totalToSendRewarded.IsZero() {
		err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, string(types.ProviderRewardsDistributionPool), subscriptionTypes.ModuleName, sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), totalToSendRewarded)))
		if err != nil {
			utils.LavaFormatError("failed to send bonus rewards to subscription module", err, utils.LogAttr("amount", totalRewarded.String()))
		}
	}

	tokensToBurn := total.Sub(totalRewarded)
	err := k.BurnPoolTokens(ctx, types.ProviderRewardsDistributionPool, tokensToBurn)
	if err != nil {
		utils.LavaFormatError("Failed to burn left over bonus rewards", err, utils.LogAttr("amount", tokensToBurn.String()))
	}
}

func (k Keeper) specTotalPayout(ctx sdk.Context, totalMonthlyPayout math.Int, totalProvidersBaseRewards sdk.Dec, spec types.SpecEmmisionPart) math.LegacyDec {
	specPayoutAllocation := spec.Emission.MulInt(totalMonthlyPayout)
	rewardBoost := totalProvidersBaseRewards.MulInt64(int64(k.MaxRewardBoost(ctx)))
	diminishingRewards := sdk.MaxDec(sdk.ZeroDec(), (sdk.NewDecWithPrec(15, 1).Mul(specPayoutAllocation)).Sub(sdk.NewDecWithPrec(5, 1).Mul(totalProvidersBaseRewards)))
	return sdk.MinDec(sdk.MinDec(specPayoutAllocation, rewardBoost), diminishingRewards)
}

func (k Keeper) specEmissionParts(ctx sdk.Context) (emisions []types.SpecEmmisionPart) {
	chainIDs := k.specKeeper.GetAllChainIDs(ctx)
	totalStake := sdk.ZeroDec()
	chainStake := map[string]sdk.Dec{}
	for _, chainID := range chainIDs {
		spec, found := k.specKeeper.GetSpec(ctx, chainID)
		if !found {
			continue
		}

		stakeStorage, found := k.epochstorage.GetStakeStorageCurrent(ctx, chainID)
		if !found {
			continue
		}
		chainStake[chainID] = sdk.ZeroDec()
		for _, entry := range stakeStorage.StakeEntries {
			chainStake[chainID] = chainStake[chainID].Add(sdk.NewDecFromInt(entry.EffectiveStake()))
		}

		chainStake[chainID] = chainStake[chainID].MulInt64(int64(spec.Shares))
		totalStake = totalStake.Add(chainStake[chainID])
	}

	for _, chainID := range chainIDs {
		if stake, ok := chainStake[chainID]; ok {
			if stake.IsZero() || totalStake.IsZero() {
				continue
			}

			emisions = append(emisions, types.SpecEmmisionPart{ChainID: chainID, Emission: stake.Quo(totalStake)})
		}
	}

	return emisions
}

func (k Keeper) specProvidersBasePay(ctx sdk.Context, chainID string) ([]types.BasePayWithIndex, math.Int) {
	basepays := k.popAllBasePayForChain(ctx, chainID)
	totalBasePay := math.ZeroInt()
	for _, basepay := range basepays {
		totalBasePay = totalBasePay.Add(basepay.Total)
	}
	return basepays, totalBasePay
}

// ContributeToValidatorsAndCommunityPool transfers some of the providers' rewards to the validators and community pool
// the function return the updated reward after the participation deduction
func (k Keeper) ContributeToValidatorsAndCommunityPool(ctx sdk.Context, reward math.Int, senderModule string) (updatedReward math.Int, err error) {
	// calculate validators and community participation fractions
	validatorsParticipation, communityParticipation, err := k.CalculateValidatorsAndCommunityContribution(ctx, reward, senderModule)
	if err != nil {
		return reward, err
	}

	if communityParticipation.Equal(sdk.OneDec()) {
		err := k.fundCommunityPoolFromModule(ctx, reward, senderModule)
		if err != nil {
			return reward, utils.LavaFormatError("failed funding the community pool with whole reward", err,
				utils.Attribute{Key: "reward", Value: reward.String()},
				utils.Attribute{Key: "community_participation", Value: communityParticipation.String()},
			)
		}
		return sdk.ZeroInt(), nil
	}

	// send validators participation and update reward
	validatorsParticipationReward := validatorsParticipation.MulInt(reward).TruncateInt()
	coins := sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), validatorsParticipationReward))
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, senderModule, k.feeCollectorName, coins)
	if err != nil {
		return reward, utils.LavaFormatError("sending validators participation failed", err,
			utils.Attribute{Key: "validators_participation_reward", Value: coins.String()},
			utils.Attribute{Key: "validators_participation", Value: validatorsParticipation.String()},
			utils.Attribute{Key: "reward", Value: reward.String()},
		)
	}
	reward = reward.Sub(validatorsParticipationReward)

	// send community participation and update reward
	communityParticipationReward := communityParticipation.MulInt(reward).TruncateInt()
	coins = sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), communityParticipationReward))
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, senderModule, k.feeCollectorName, coins)
	if err != nil {
		return reward, utils.LavaFormatError("sending community participation failed", err,
			utils.Attribute{Key: "community_participation_reward", Value: coins.String()},
			utils.Attribute{Key: "community_participation", Value: communityParticipation.String()},
			utils.Attribute{Key: "reward", Value: reward.String()},
		)
	}
	reward = reward.Sub(communityParticipationReward)

	return reward, nil
}

func (k Keeper) CalculateValidatorsAndCommunityContribution(ctx sdk.Context, reward math.Int, senderModule string) (validatorsParticipation math.LegacyDec, communityParticipation math.LegacyDec, err error) {
	communityTax := k.distributionKeeper.GetParams(ctx).CommunityTax
	if communityTax.Equal(sdk.OneDec()) {
		return sdk.ZeroDec(), sdk.OneDec(), nil
	}

	// validators_participation = validators_participation_param / (1-community_tax)
	validatorsParticipationParam := k.GetParams(ctx).ValidatorsSubscriptionParticipation
	validatorsParticipation = validatorsParticipationParam.Quo((sdk.OneDec().Sub(communityTax)))
	if validatorsParticipation.GT(sdk.OneDec()) {
		return sdk.ZeroDec(), sdk.ZeroDec(), utils.LavaFormatError("validators participation bigger than 100%", fmt.Errorf("validators participation calc failed"),
			utils.Attribute{Key: "validators_participation", Value: validatorsParticipation.String()},
			utils.Attribute{Key: "validators_subscription_participation_param", Value: validatorsParticipationParam.String()},
			utils.Attribute{Key: "community_tax", Value: communityTax.String()},
		)
	}

	// community_participation = (community_tax + validators_participation_param) - validators_participation
	communityParticipation = communityTax.Add(validatorsParticipationParam).Sub(validatorsParticipation)
	if communityParticipation.IsNegative() || communityParticipation.GT(sdk.OneDec()) {
		return sdk.ZeroDec(), sdk.ZeroDec(), utils.LavaFormatError("community participation is negative or bigger than 100%", fmt.Errorf("community participation calc failed"),
			utils.Attribute{Key: "community_participation", Value: communityParticipation.String()},
			utils.Attribute{Key: "validators_participation", Value: validatorsParticipation.String()},
			utils.Attribute{Key: "validators_subscription_participation_param", Value: validatorsParticipationParam.String()},
			utils.Attribute{Key: "community_tax", Value: communityTax.String()},
		)
	}

	// check the participation rewards are not more than 100%
	if validatorsParticipation.Add(communityParticipation).GT(sdk.OneDec()) {
		return sdk.ZeroDec(), sdk.ZeroDec(), utils.LavaFormatError("validators and community participation parts are bigger than 100%", fmt.Errorf("validators and community participation aborted"),
			utils.Attribute{Key: "community_participation", Value: communityParticipation.String()},
			utils.Attribute{Key: "validators_participation", Value: validatorsParticipation.String()},
		)
	}

	return validatorsParticipation, communityParticipation, nil
}

func (k Keeper) fundCommunityPoolFromModule(ctx sdk.Context, amount math.Int, senderModule string) error {
	coins := sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), amount))
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, senderModule, distributiontypes.ModuleName, coins); err != nil {
		return err
	}

	feePool := k.distributionKeeper.GetFeePool(ctx)
	feePool.CommunityPool = feePool.CommunityPool.Add(sdk.NewDecCoinsFromCoins(coins...)...)
	k.distributionKeeper.SetFeePool(ctx, feePool)

	return nil
}
