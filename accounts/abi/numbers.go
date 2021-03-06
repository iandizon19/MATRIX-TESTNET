// Copyright 2018 The MATRIX Authors as well as Copyright 2014-2017 The go-ethereum Authors
// This file is consisted of the MATRIX library and part of the go-ethereum library.
//
// The MATRIX-ethereum library is free software: you can redistribute it and/or modify it under the terms of the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, 
//and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject tothe following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, 
//WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISINGFROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
//OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package abi

import (
	"math/big"
	"reflect"

	"github.com/matrix/go-matrix/common"
	"github.com/matrix/go-matrix/common/math"
)

var (
	bigT      = reflect.TypeOf(&big.Int{})
	derefbigT = reflect.TypeOf(big.Int{})
	uint8T    = reflect.TypeOf(uint8(0))
	uint16T   = reflect.TypeOf(uint16(0))
	uint32T   = reflect.TypeOf(uint32(0))
	uint64T   = reflect.TypeOf(uint64(0))
	intT      = reflect.TypeOf(int(0))
	int8T     = reflect.TypeOf(int8(0))
	int16T    = reflect.TypeOf(int16(0))
	int32T    = reflect.TypeOf(int32(0))
	int64T    = reflect.TypeOf(int64(0))
	addressT  = reflect.TypeOf(common.Address{})
	intTS     = reflect.TypeOf([]int(nil))
	int8TS    = reflect.TypeOf([]int8(nil))
	int16TS   = reflect.TypeOf([]int16(nil))
	int32TS   = reflect.TypeOf([]int32(nil))
	int64TS   = reflect.TypeOf([]int64(nil))
)

// U256 converts a big Int into a 256bit EVM number.
func U256(n *big.Int) []byte {
	return math.PaddedBigBytes(math.U256(n), 32)
}

// checks whether the given reflect value is signed. This also works for slices with a number type
func isSigned(v reflect.Value) bool {
	switch v.Type() {
	case intTS, int8TS, int16TS, int32TS, int64TS, intT, int8T, int16T, int32T, int64T:
		return true
	}
	return false
}
