package cef

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "browser.h"
#include "cef_frame.h"
*/
import "C"

import (
	"sync"
	"time"
	"unsafe"
)

type CefFrame struct {
	CStruct *C.struct__cef_frame_t
}

func (self CefFrame) IsValid() bool {
	return (C.cef_frame_t_is_valid(self.CStruct) == 1)
}

func (self CefFrame) GetURL() string {
	return CEFToGoString(C.cef_frame_t_get_url(self.CStruct))
}

func (self CefFrame) GetSource() string {
	visitor := NewCefStringVisitor()
	defer visitor.Close()

	stringVisitorLock.Lock()
	stringVisitorMap[unsafe.Pointer(visitor.CStruct)] = visitor
	stringVisitorLock.Unlock()

	visitor.Add(1)
	C.cef_frame_t_get_source(self.CStruct, visitor.CStruct)
	visitor.Wait(10 * time.Second)
	return visitor.data
}

func (self CefFrame) GetText() string {
	visitor := NewCefStringVisitor()
	defer visitor.Close()

	stringVisitorLock.Lock()
	stringVisitorMap[unsafe.Pointer(visitor.CStruct)] = visitor
	stringVisitorLock.Unlock()

	visitor.Add(1)
	C.cef_frame_t_get_text(self.CStruct, visitor.CStruct)
	visitor.Wait(10 * time.Second)
	return visitor.data
}

// 执行JS
func (self CefFrame) ExecuteJavaScript(code, url string, startLine int) {
	codeCString := C.CString(code)
	defer C.free(unsafe.Pointer(codeCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.cef_frame_t_execute_java_script(self.CStruct, codeCString, urlCString, C.int(startLine))
}

// `var cefEvalFunc = function() { return 1; };var cef_result = cefEvalFunc();app.cefResult(cef_result);`
func (self CefFrame) ExecuteJavaScriptWithResult(code string) string {
	browser := self.GetBrowser()
	var wg sync.WaitGroup
	wg.Add(1)
	evalJsLock.Lock()
	evalJsMap[unsafe.Pointer(browser.GetCStruct())] = &wg
	evalJsLock.Unlock()
	self.ExecuteJavaScript(code, ``, 1)
	wg.Wait()
	return evalJavaScriptResult
}

func (self CefFrame) GetBrowser() (browser *Browser) {
	b := C.cef_frame_t_get_browser(self.CStruct)
	identifier := C.cef_browser_t_get_identifier(b)
	browser_, ok := BrowserById(int(identifier))
	if ok {
		browser = browser_
	}
	return
}
