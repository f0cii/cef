package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

extern void initialize_cef_string_visitor(struct _cef_string_visitor_t* visitor);

*/
import "C"

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

var (
	stringVisitorMap  = make(map[unsafe.Pointer]*CefStringVisitor)
	stringVisitorLock sync.Mutex
)

type CefStringVisitor struct {
	CStruct *C.struct__cef_string_visitor_t
	data    string
	//wg      sync.WaitGroup
	done chan error
}

func (v *CefStringVisitor) Close() {
	go_Release(unsafe.Pointer(v.CStruct))
}

func (v *CefStringVisitor) Add(n int) {
	//v.wg.Add(n)
}

func (v *CefStringVisitor) Done(err error) {
	//v.wg.Done()
	go func() { v.done <- err }()
}

func (v *CefStringVisitor) Wait(timeout time.Duration) (error, bool) {
	var err error
	select {
	case err = <-v.done:
		return err, false
	}
}

func NewCefStringVisitor() *CefStringVisitor {
	v := &CefStringVisitor{}

	var visitor *C.struct__cef_string_visitor_t
	visitor = (*C.struct__cef_string_visitor_t)(C.calloc(1, C.sizeof_struct__cef_string_visitor_t))
	go_AddRef(unsafe.Pointer(visitor))

	C.initialize_cef_string_visitor(visitor)

	v.CStruct = visitor
	v.done = make(chan error)

	return v
}

//export go_StringVisitorVisit
func go_StringVisitorVisit(self unsafe.Pointer, s *C.cef_string_t) {
	//fmt.Printf("go_StringVisitorVisit self=%v\n", self)

	stringVisitorLock.Lock()
	visitor, ok := stringVisitorMap[unsafe.Pointer(self)]
	if ok {
		delete(stringVisitorMap, unsafe.Pointer(self))
	}
	stringVisitorLock.Unlock()
	if ok {
		//v := (*CefStringVisitorT)(self)
		v := visitor
		v.data = CEFToGoString(s)
		v.Done(errors.New("ok"))
	} else {
		fmt.Printf("go_StringVisitorVisit fail.\n")
	}
}
