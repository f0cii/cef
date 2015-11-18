// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/nvsoft/cef

package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

var (
	memoryBridge = make(map[unsafe.Pointer]int)
	refCountLock sync.Mutex
)

//export go_AddRef
func go_AddRef(it unsafe.Pointer) int {
	refCountLock.Lock()
	defer refCountLock.Unlock()
	m, ok := memoryBridge[it]
	if !ok {
		m = 0
	}
	m++
	memoryBridge[it] = m
	return m
}

//export go_Release
func go_Release(it unsafe.Pointer) int {
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if m, ok := memoryBridge[it]; ok {
		m--
		memoryBridge[it] = m
		if m == 0 {
			C.free(it)
			delete(memoryBridge, it)
		}
		return m
	}
	return 1
}

//export go_GetRefCount
func go_GetRefCount(it unsafe.Pointer) int {
	refCountLock.Lock()
	defer refCountLock.Unlock()

	if m, ok := memoryBridge[it]; ok {
		return m
	}
	return 1
}
