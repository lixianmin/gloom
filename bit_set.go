package gloom

import "github.com/bits-and-blooms/bitset"

/********************************************************************
created:    2022-03-28
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type BitSet struct {
	bitSet *bitset.BitSet
}

func NewBitSet(bitSize uint) *BitSet {
	return &BitSet{bitSet: bitset.New(bitSize)}
}

func (b *BitSet) Set(offsets []uint) error {
	for _, offset := range offsets {
		b.bitSet.Set(offset)
	}
	return nil
}

func (b *BitSet) Test(offsets []uint) (bool, error) {
	for _, offset := range offsets {
		if !b.bitSet.Test(offset) {
			return false, nil
		}
	}

	return true, nil
}
