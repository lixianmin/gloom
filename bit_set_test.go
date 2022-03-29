package gloom

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/********************************************************************
created:    2022-03-28
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func TestBitSetBloomFilter(t *testing.T) {
	const estimatedKeys = 1000
	var bitSize, locationNum = EstimateParameters(estimatedKeys, 0.001)
	var bloomFilter = New(bitSize, locationNum, NewBitSet(bitSize))

	fmt.Printf("estimatedKeys=%d, bitSize=%d, locationNum=%d\n", estimatedKeys, bitSize, locationNum)
	testBloomFilter(t, estimatedKeys, bloomFilter)
}

func testBloomFilter(t *testing.T, estimatedKeys int, bloomFilter *BloomFilter) {
	var startTime = time.Now()
	var falseCounter = 0
	for i := 0; i < estimatedKeys; i++ {
		var data = []byte(strconv.Itoa(i * 100))

		var existsBefore, err = bloomFilter.Exists(data)
		if err != nil {
			t.Logf("Error checking for existence in bloom filter, err=%v", err)
		}

		if existsBefore {
			falseCounter++
			t.Logf("i=%d, falseCounter=%d", i, falseCounter)
		}

		err = bloomFilter.Add(data)
		if err != nil {
			t.Logf("add err=%v", err)
		}
	}

	var costTime = time.Now().Sub(startTime)
	fmt.Printf("estimatedKeys=%d, costTime=%s\n", estimatedKeys, costTime.String())
}

func Benchmark_BitSet(b *testing.B) {
	var estimatedKeys = 100000
	var bitSize, locationNum = EstimateParameters(estimatedKeys, 0.001)
	var bloomFilter = New(bitSize, locationNum, NewBitSet(bitSize))

	b.Logf("estimatedKeys=%d, bitSize=%d, locationNum=%d", estimatedKeys, bitSize, locationNum)
	benchmarkBloomFilter(b, estimatedKeys, bloomFilter)
}

func benchmarkBloomFilter(b *testing.B, estimatedKeys int, bloomFilter *BloomFilter) {
	b.ResetTimer()

	var falseCounter = 0
	for i := 0; i < estimatedKeys; i++ {
		var data = []byte(strconv.Itoa(i * 100))

		var existsBefore, err = bloomFilter.Exists(data)
		if err != nil {
			b.Logf("Error checking for existence in bloom filter, err=%v", err)
		}

		if existsBefore {
			falseCounter++
			b.Logf("i=%d, falseCounter=%d", i, falseCounter)
		}

		err = bloomFilter.Add(data)
		if err != nil {
			b.Logf("add err=%v", err)
		}
	}
}
