// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "runtime"

// extern void callGoWithVariousStack();
import "C"

func main() {}

//export GoF
func GoF() { runtime.GC() }

//export callGoWithVariousStackAndGoFrame
func callGoWithVariousStackAndGoFrame() {
	C.callGoWithVariousStack();
}
