package maps

import (
	"encoding/binary"

	"github.com/lavanet/lava/v2/utils/lavaslices"
	"golang.org/x/exp/constraints"
)

func FindLargestIntValueInMap[K comparable](myMap map[K]int) (K, int) {
	var maxVal int
	var maxKey K
	firstIteration := true

	for key, val := range myMap {
		if firstIteration || val > maxVal {
			maxVal = val
			maxKey = key
			firstIteration = false
		}
	}

	return maxKey, maxVal
}

// GetIDBytes returns the byte representation of the uint64 ID
func GetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIDFromBytes returns ID in uint64 format from a byte array
func GetIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func StableSortedKeys[T constraints.Ordered, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	lavaslices.SortStable(keys)
	return keys
}

func GetMaxKey[T constraints.Ordered, V any](m map[T]V) T {
	var maxKey T
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}
	return maxKey
}
