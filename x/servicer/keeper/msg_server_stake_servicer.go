package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/servicer/types"
)

func (k msgServer) StakeServicer(goCtx context.Context, msg *types.MsgStakeServicer) (*types.MsgStakeServicerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	specName := msg.Spec
	err := specName.ValidateBasic() //TODO: basic validation, we dont want to read the entire spec list here
	if err != nil {
		return nil, err
	}

	foundAndActive, _ := k.Keeper.specKeeper.IsSpecFoundAndActive(ctx, specName.Name)
	if foundAndActive != true {
		return nil, errors.New("spec not found or not enabled")
	}
	//if we get here, the spec is active and supported
	if msg.Amount.IsLT(k.Keeper.GetMinStake(ctx)) {
		return nil, errors.New(fmt.Sprintf("insufficient stake amount, amount must be %s or greater", k.Keeper.GetMinStake(ctx)))
	}
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid creator address %s error: %s", msg.Creator, err))
	}
	//define the function here for later use
	verifySufficientAmountAndSendToModule := func(ctx sdk.Context, k msgServer, addr sdk.AccAddress, neededAmount sdk.Coin) (bool, error) {
		if k.Keeper.bankKeeper.GetBalance(ctx, addr, "stake").IsLT(neededAmount) {
			return false, errors.New(fmt.Sprintf("insufficient balance for staking %s current balance: %s", neededAmount, k.Keeper.bankKeeper.GetBalance(ctx, addr, "stake")))
		}
		err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, "servicer", []sdk.Coin{neededAmount})
		if err != nil {
			return false, errors.New(fmt.Sprintf("invalid transfer coins to module, %s", err))
		}
		return true, nil
	}

	// }
	specStakeStorage, found := k.Keeper.GetSpecStakeStorage(ctx, specName.Name)
	if found != true {

		//this is the first servicer for the supported spec
		stakeStorage := types.StakeStorage{
			Staked:    []types.StakeMap{},
			Unstaking: []types.StakeMap{},
		}
		specStakeStorage = types.SpecStakeStorage{
			Index:        specName.Name,
			StakeStorage: &stakeStorage,
		}
		// newSpecStakeStorage :=
		// k.Keeper.SetSpecStakeStorage(ctx, newSpecStakeStorage)
	}
	stakeStorage := specStakeStorage.StakeStorage
	//TODO: improve the finding logic and the way its saved looping a list is slow and bad
	for _, storageMap := range stakeStorage.Staked {
		if storageMap.Index == msg.Creator {
			// already exists
			if storageMap.Stake.IsLT(msg.Amount) {
				// increasing stake is allowed
				if storageMap.Deadline.Num >= msg.Deadline.Num {
					//lowering the deadline is allowed
					valid, err := verifySufficientAmountAndSendToModule(ctx, k, senderAddr, msg.Amount.Sub(storageMap.Stake))
					if !valid {
						return nil, errors.New(fmt.Sprintf("error updating stake: %s", err))
					}
					storageMap.Stake = msg.Amount
					storageMap.Deadline = msg.Deadline
				}
				return nil, errors.New("can't increase deadline for existing servicer")
			}
			return nil, errors.New("can't decrease stake for existing servicer")
		}
	}
	// servicer isn't staked so add him
	valid, err := verifySufficientAmountAndSendToModule(ctx, k, senderAddr, msg.Amount)
	if !valid {
		return nil, errors.New(fmt.Sprintf("error staking: %s", err))
	}

	stakeStorage.Staked = append(stakeStorage.Staked, types.StakeMap{
		Index:    msg.Creator,
		Stake:    msg.Amount,
		Deadline: msg.Deadline,
	})
	k.Keeper.SetSpecStakeStorage(ctx, specStakeStorage)
	_ = ctx

	return &types.MsgStakeServicerResponse{}, nil
}
