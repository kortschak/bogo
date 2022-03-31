// Copyright ©2022 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bogo_test

import (
	"math"
	"sort"
	"testing"

	"github.com/kortschak/bogo"
)

var (
	// Copyright 2009 The Go Authors. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	// https://cs.opensource.google/go/go/+/master:LICENSE
	ints     = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
	strings  = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}
)

func TestBogoIntSlice(t *testing.T) {
	data := ints[:11] // Limit input to feasible length.
	a := sort.IntSlice(data[0:])
	bogo.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestBogoFloat64Slice(t *testing.T) {
	data := float64s[:11] // Limit input to feasible length.
	a := sort.Float64Slice(data[0:])
	bogo.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestBogoStringSlice(t *testing.T) {
	data := strings
	a := sort.StringSlice(data[0:])
	bogo.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}
