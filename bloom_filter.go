package gloom

import (
	"errors"
	"github.com/spaolacci/murmur3"
	"math"
)

/********************************************************************
created:    2022-03-28
author:     lixianmin

参考: https://github.com/bculberson/bloom

Copyright (C) - All Rights Reserved
*********************************************************************/

var ErrDataIsNil = errors.New("data is nil")

type BitSetProvider interface {
	Set([]uint) error
	Test([]uint) (bool, error)
}

type BloomFilter struct {
	bitSize     uint
	locationNum uint
	bitSet      BitSetProvider
}

func New(bitSize uint, locationNum uint, bitSet BitSetProvider) *BloomFilter {
	if bitSet == nil {
		panic("bitset should not be nil")
	}

	var filter = &BloomFilter{bitSize: bitSize, locationNum: locationNum, bitSet: bitSet}
	return filter
}

func EstimateParameters(estimatedKeys int, falsePositiveRate float64) (uint, uint) {
	var bitSize = math.Ceil(float64(estimatedKeys) * math.Log(falsePositiveRate) / math.Log(1.0/math.Pow(2.0, math.Ln2)))
	var locationNum = math.Ln2*bitSize/float64(estimatedKeys) + 0.5
	return uint(bitSize), uint(locationNum)
}

func (my *BloomFilter) Add(data []byte) error {
	if len(data) == 0 {
		return ErrDataIsNil
	}

	var locations = my.getLocations(data)
	err := my.bitSet.Set(locations)
	if err != nil {
		return err
	}

	return nil
}

func (my *BloomFilter) Exists(data []byte) (bool, error) {
	if len(data) == 0 {
		return false, ErrDataIsNil
	}

	var locations = my.getLocations(data)
	var isSet, err = my.bitSet.Test(locations)

	if err != nil {
		return false, err
	}

	if !isSet {
		return false, nil
	}

	return true, nil
}

func (my *BloomFilter) getLocations(data []byte) []uint {
	var size = my.locationNum
	var locations = make([]uint, size)

	for i := uint(0); i < size; i++ {
		data = append(data, byte(i))
		var hashValue = bloomHash(data)
		locations[i] = uint(hashValue % uint64(my.bitSize))
	}

	return locations
}

func bloomHash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
