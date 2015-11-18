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

extern void initialize_cef_dom_visitor(struct _cef_domvisitor_t* visitor);

extern void VisitDOM(cef_browser_t* browser, cef_domvisitor_t* visitor);

*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

var (
	domVisitorMap  = make(map[unsafe.Pointer]*CefDOMVisitor)
	domVisitorLock sync.Mutex
)

type CefDOMVisitor struct {
	CStruct *C.struct__cef_domvisitor_t
	wg      sync.WaitGroup
}

func (v *CefDOMVisitor) Close() {
	go_Release(unsafe.Pointer(v.CStruct))
}

func (v *CefDOMVisitor) Add(n int) {
	v.wg.Add(n)
}

func (v *CefDOMVisitor) Done() {
	v.wg.Done()
}

func (v *CefDOMVisitor) Wait() {
	v.wg.Wait()
}

func NewCefDOMVisitor() *CefDOMVisitor {
	v := &CefDOMVisitor{}

	var visitor *C.struct__cef_domvisitor_t
	visitor = (*C.struct__cef_domvisitor_t)(C.calloc(1, C.sizeof_struct__cef_domvisitor_t))
	go_AddRef(unsafe.Pointer(visitor))

	C.initialize_cef_dom_visitor(visitor)

	v.CStruct = visitor

	return v
}

//export go_DOMVisitorVisit
func go_DOMVisitorVisit(self unsafe.Pointer, document *C.cef_domdocument_t) {
	fmt.Println("go_DOMVisitorVisit")

	domVisitorLock.Lock()
	visitor, ok := domVisitorMap[unsafe.Pointer(self)]
	if ok {
		delete(domVisitorMap, unsafe.Pointer(self))
	}
	domVisitorLock.Unlock()
	if ok {
		v := visitor
		v.Done()
	}
}
