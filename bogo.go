// Copyright ©2022 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bogo provides near pessimal sort functionality.
package bogo

import (
	"math/rand"
	"sort"
)

// Sort sorts data. It makes on average O((n+1)!) calls to data.Len, data.Less and data.Swap.
// The sort is not guaranteed to be stable.
func Sort(data sort.Interface) {
	// Implementation is according to Gruber, Holzer and Ruepp, doi:10.1007/978-3-540-72914-3_17.
	for !sort.IsSorted(data) {
		rand.Shuffle(data.Len(), func(i, j int) {
			data.Swap(i, j)
		})
	}
}
