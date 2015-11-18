package cef

/*
#include <stdlib.h>
#include "string.h"
#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "cef_helpers.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type V8CallbackHandler func(*Browser, *CefProcessMessage) interface{}

var V8CallbackHandlers map[string]V8CallbackHandler

func init() {
	V8CallbackHandlers = make(map[string]V8CallbackHandler)
}

func RegisterV8CallbackHandler(name string, callback V8CallbackHandler) {
	V8CallbackHandlers[name] = callback
}

//export go_OnBeforePopup
func go_OnBeforePopup(
	self *C.struct__cef_life_span_handler_t,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	target_url *C.cef_string_t,
	target_frame_name *C.cef_string_t,
	popupFeatures *C.struct__cef_popup_features_t,
	windowInfo *C.struct__cef_window_info_t,
	client **C.struct__cef_client_t,
	settings *C.struct__cef_browser_settings_t,
	no_javascript_access *C.int) int {

	//C.releaseVoid(unsafe.Pointer(browser))*
	//C.releaseVoid(unsafe.Pointer(frame))*
	//C.releaseVoid(unsafe.Pointer(popupFeatures))*
	//C.releaseVoid(unsafe.Pointer(windowInfo))*
	//C.releaseVoid(unsafe.Pointer(client))
	//C.releaseVoid(unsafe.Pointer(settings))*
	logger.Println("go_OnBeforePopup")
	if createBrowserWindowCallback != nil {
		createBrowserWindowCallback(CEFToGoString(target_url))
		return 1
	}
	return 0
}

//export go_OnAfterCreated
func go_OnAfterCreated(self *C.struct__cef_life_span_handler_t, browserId int, browser *C.cef_browser_t) {
	//logger.Println("go_OnAfterCreated")
	if globalLifespanHandler != nil {
		globalLifespanHandler.OnAfterCreated(&Browser{browserId, browser, nil})
	}
}

//export go_RunModal
func go_RunModal(
	self *C.struct__cef_life_span_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t) int {
	logger.Println("go_RunModal")
	if globalLifespanHandler != nil {
		return globalLifespanHandler.RunModal(&Browser{browserId, browser, nil})
	} else {
		logger.Println("go_RunModal releaseVoid")
		C.releaseVoid(unsafe.Pointer(browser))
	}
	return 0
}

//export go_DoClose
func go_DoClose(
	self *C.struct__cef_life_span_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t) int {
	logger.Println("go_DoClose")
	if globalLifespanHandler != nil {
		globalLifespanHandler.DoClose(&Browser{browserId, browser, nil})
	} else {
		C.releaseVoid(unsafe.Pointer(browser))
	}
	return 0
}

//export go_BeforeClose
func go_BeforeClose(
	self *C.struct__cef_life_span_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t) {
	logger.Println("go_BeforeClose")
	if globalLifespanHandler != nil {
		globalLifespanHandler.BeforeClose(&Browser{browserId, browser, nil})
	} else {
		C.releaseVoid(unsafe.Pointer(browser))
	}
}

//export go_OnLoadingStateChange
func go_OnLoadingStateChange(
	self *C.struct__cef_load_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t,
	isLoading, canGoBack,
	canGoForward int) {
}

//export go_OnLoadStart
func go_OnLoadStart(
	self *C.struct__cef_load_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t) {
	//fmt.Println("go_OnLoadStart")
}

//export go_OnLoadEnd
func go_OnLoadEnd(
	self *C.struct__cef_load_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	httpStatusCode int) {
}

//export go_OnLoadError
func go_OnLoadError(
	self *C.struct__cef_load_handler_t,
	browserId int,
	browser *C.struct__cef_browser_t,
	frame *C.struct__cef_frame_t,
	errorCode C.cef_errorcode_t,
	errorText *C.cef_string_t,
	failedUrl *C.cef_string_t) {
	fmt.Println("go_OnLoadError")
}

//export go_Log
func go_Log(str *C.char) {
	logger.Println(C.GoString(str))
}

//export go_LogPointer
func go_LogPointer(str *C.char, p unsafe.Pointer) {
	logger.Printf(C.GoString(str)+" %p\n", p)
}

//export go_OnTitleChange
func go_OnTitleChange(browserId int, browser *C.cef_browser_t, title *C.cef_string_t) {
	logger.Println("go_OnTitleChange.")
	if b, ok := BrowserById(browserId); ok {
		hWnd := b.GetRootWindowHandle()
		SetWindowText(hWnd, CEFToGoString(title))
	}
}

//export go_OnConsoleMessage
func go_OnConsoleMessage(browser *C.cef_browser_t, message *C.cef_string_t, source *C.cef_string_t, line int) {
	logger.Println("go_OnConsoleMessage.")
	consoleHandler(CEFToGoString(message), CEFToGoString(source), line)
}

//export go_OnProcessMessageReceived
func go_OnProcessMessageReceived(self *C.cef_client_t, browserId int, browser *C.cef_browser_t, source_process C.cef_process_id_t,
	message *C.cef_process_message_t) int {
	if b, ok := BrowserById(browserId); ok {
		m := NewCefProcessMessage(message)
		name := m.GetName()
		if name == kV8EvalResultMessage {
			b.onExecuteJavaScriptResult(m)
		}
		if name == kV8OpenWindowMessage {
			// 打开Window
			url := m.GetArgumentList().GetString(0)
			if OnOpenWindow != nil {
				OnOpenWindow(url)
			}
		}
		if name == kV8CallbackMessage {
			callbackName := m.GetArgumentList().GetString(0)
			h, ok := V8CallbackHandlers[callbackName]
			if ok {
				h(b, m)
			} else {
				fmt.Printf("V8CallbackMessage not matched. [%v]\n", callbackName)
			}
		}
	} else {
		logger.Fatalln("go_OnProcessMessageReceived fail.")
	}
	return 1
}
