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
#include "browser.h"
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	//"os"
	"sync"
	"time"
	"unsafe"
)

var browsers map[int]*Browser
var browserHandles map[unsafe.Pointer]*Browser
var mainBrowserId = -1
var browserWaitGroup sync.WaitGroup
var browserCreating bool
var createBrowserWindowCallback CreateBrowserWindowCallback

var evalJsLock sync.Mutex
var evalJsMap = make(map[unsafe.Pointer]*sync.WaitGroup)
var evalJavaScriptResult string

type CreateBrowserWindowCallback func( /*url*/ string) /*hwnd*/ unsafe.Pointer

func init() {
	browsers = make(map[int]*Browser)
	browserHandles = make(map[unsafe.Pointer]*Browser)
}

func SetupCreateBrowserWindowCallback(handler CreateBrowserWindowCallback) {
	createBrowserWindowCallback = handler
}

func CreateBrowser(hwnd unsafe.Pointer, browserSettings *BrowserSettings, url string, offscreenRendering bool) (browser *Browser) {
	logger.Printf("CreateBrowser, url=%s\n", url)

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)
	if offscreenRendering {
		windowInfo.windowless_rendering_enabled = 1
		windowInfo.transparent_painting_enabled = 1
	}

	browserCreating = true

	var clientHandler *C.struct__cef_client_t
	clientHandler = (*C.struct__cef_client_t)(C.calloc(1, C.sizeof_struct__cef_client_t))
	go_AddRef(unsafe.Pointer(clientHandler))
	C.initialize_client_handler(clientHandler)

	cRet := C.cef_browser_host_create_browser(windowInfo, clientHandler, CEFString(url), browserSettings.ToCStruct(), nil) //_ClientHandler
	fmt.Printf("CreateBrowser Ret=%v\n", cRet)

	return nil
}

func SafeExecuteJavaScript(browser *C.cef_browser_t, code, url string, startLine int) {
	codeCString := C.CString(code)
	defer C.free(unsafe.Pointer(codeCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.ExecuteJavaScript(browser, codeCString, urlCString, C.int(startLine))
}

type Browser struct {
	Id            int
	cbrowser      *C.cef_browser_t
	RenderHandler RenderHandler
}

func MainBrowser() (browser *Browser, ok bool) {
	browser, ok = BrowserById(mainBrowserId)
	return
}

func BrowserById(id int) (browser *Browser, ok bool) {
	browser, ok = browsers[id]
	return
}

func BrowserByHandle(handle unsafe.Pointer) (browser *Browser, ok bool) {
	for _, v := range browsers {
		if v.GetRootWindowHandle() == handle {
			return v, true
		}
	}
	return nil, false
}

func (b Browser) GetCStruct() *C.cef_browser_t {
	return b.cbrowser
}

func (b Browser) GetHost() (host *BrowserHost) {
	h := C.cef_browser_t_get_host(b.cbrowser)
	host = &BrowserHost{cbrowserhost: h}
	return
}

func (b Browser) GetMainFrame() (frame *CefFrame) {
	f := C.cef_browser_t_get_main_frame(b.cbrowser)
	frame = &CefFrame{CStruct: f}
	return
}

func (b Browser) GetFocusedFrame() (frame *CefFrame) {
	f := C.cef_browser_t_get_focused_frame(b.cbrowser)
	frame = &CefFrame{CStruct: f}
	return
}

func (b Browser) GetFrameByIdent(identifier int64) (frame *CefFrame) {
	f := C.cef_browser_t_get_frame_byident(b.cbrowser, C.int64(identifier))
	frame = &CefFrame{CStruct: f}
	return
}

func (b Browser) GetFrame(name string) (frame *CefFrame) {
	cefString := CEFString(name)
	f := C.cef_browser_t_get_frame(b.cbrowser, cefString)
	frame = &CefFrame{CStruct: f}
	return
}

func (b Browser) GetFrameCount() int {
	count := C.cef_browser_t_get_frame_count(b.cbrowser)
	return int(count)
}

func (b Browser) GetFrameIdentifiers() []int64 {
	var count C.size_t = C.size_t(b.GetFrameCount())
	var ids *C.int64 = (*C.int64)(C.calloc(count, C.sizeof_int64))
	C.cef_browser_t_get_frame_identifiers(b.cbrowser, &count, ids)
	rCount := int64(count)
	result := make([]int64, rCount)
	var i int64
	for i = 0; i < rCount; i++ {
		result[i] = int64(C.int64_array_get(ids, C.size_t(i)))
	}
	C.free(unsafe.Pointer(ids))
	return result
}

func (b Browser) GetFrameNames() []string {
	nameList := C.cef_string_list_alloc()
	C.cef_browser_t_get_frame_names(b.cbrowser, nameList)
	length := int(C.cef_string_list_size(nameList))
	goList := make([]string, length)
	for i := range goList {
		var cefName *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
		didRet := C.cef_string_list_value(nameList, C.int(i), (*C.cef_string_t)(cefName))
		if didRet == C.int(1) {
			nameUtf8 := C.cefStringToUtf8((*C.cef_string_t)(cefName))
			goList[i] = C.GoString(nameUtf8.str)
			C.cef_string_userfree_utf8_free(nameUtf8)
		}
		C.cef_string_userfree_utf16_free(cefName)
	}
	return goList
}

// 执行JS
func (b Browser) ExecuteJavaScript(code, url string, startLine int) {
	codeCString := C.CString(code)
	defer C.free(unsafe.Pointer(codeCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.ExecuteJavaScript(b.cbrowser, codeCString, urlCString, C.int(startLine))
}

// `var cefEvalFunc = function() { return 1; };var cef_result = cefEvalFunc();app.cefResult(cef_result);`
func (b Browser) ExecuteJavaScriptWithResult(code string) string {
	var wg sync.WaitGroup
	wg.Add(1)
	evalJsLock.Lock()
	evalJsMap[unsafe.Pointer(b.cbrowser)] = &wg
	evalJsLock.Unlock()
	b.ExecuteJavaScript(code, ``, 1)
	wg.Wait()
	return evalJavaScriptResult
}

func (b Browser) onExecuteJavaScriptResult(message *CefProcessMessage) {
	evalJavaScriptResult = message.GetArgumentList().GetString(0)
	evalJsLock.Lock()
	wg, ok := evalJsMap[unsafe.Pointer(b.cbrowser)]
	if ok {
		delete(evalJsMap, unsafe.Pointer(b.cbrowser))
	}
	evalJsLock.Unlock()
	if ok {
		wg.Done()
	} else {
		fmt.Printf("onExecuteJavaScriptResult fail.\n")
	}
}

func (b Browser) LoadURL(url string) {
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C.LoadURL(b.cbrowser, urlCString)
}

func (b Browser) LoadString(string_val, url string) {
	valCString := C.CString(string_val)
	defer C.free(unsafe.Pointer(valCString))
	urlCString := C.CString(url)
	defer C.free(unsafe.Pointer(urlCString))
	C._LoadString(b.cbrowser, valCString, urlCString)
}

func (b Browser) InjectJs(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("InjectJs %v\n", err)
		return
	}
	b.ExecuteJavaScript(string(content), "", 1)
}

func (b Browser) TriggerPaint() {
	C.BrowserWasResized(b.cbrowser)
}

func (b Browser) GetWindowHandle() unsafe.Pointer {
	return unsafe.Pointer(C.GetWindowHandle(b.cbrowser))
}

func (b Browser) GetOpenerWindowHandle() unsafe.Pointer {
	h := C.cef_browser_t_get_host(b.cbrowser)
	return unsafe.Pointer(C.cef_browser_host_t_get_opener_window_handle(h))
}

func (b Browser) GetRootWindowHandle() unsafe.Pointer {
	return unsafe.Pointer(C.GetRootWindowHandle(b.cbrowser))
}

func (b Browser) GetURL() string {
	return CEFToGoString(C.GetURL(b.cbrowser))
}

func (b Browser) GetSource() string {
	visitor := NewCefStringVisitor()
	defer visitor.Close()

	stringVisitorLock.Lock()
	stringVisitorMap[unsafe.Pointer(visitor.CStruct)] = visitor
	stringVisitorLock.Unlock()

	visitor.Add(1)
	C.GetSource(b.cbrowser, visitor.CStruct)
	visitor.Wait(10 * time.Second)
	return visitor.data
}

func (b Browser) GetText() string {
	visitor := NewCefStringVisitor()
	defer visitor.Close()

	stringVisitorLock.Lock()
	stringVisitorMap[unsafe.Pointer(visitor.CStruct)] = visitor
	stringVisitorLock.Unlock()

	visitor.Add(1)
	C.GetText(b.cbrowser, visitor.CStruct)
	visitor.Wait(10 * time.Second)
	return visitor.data
}

// This method can only used in Renderer process
func (b Browser) VisitDOM() {
	visitor := NewCefDOMVisitor()

	domVisitorLock.Lock()
	domVisitorMap[unsafe.Pointer(visitor.CStruct)] = visitor
	domVisitorLock.Unlock()

	visitor.Add(1)
	//C.VisitDOM(b.cbrowser, visitor.CStruct)
	//PostTask(TID_RENDERER, VisitA)
	visitor.Wait()
}

// This method can only used in Renderer process
func (b Browser) GetV8Context() *C.cef_v8context_t {
	return C.GetV8Context(b.cbrowser)
}

// This method can only used in Renderer process
func (b Browser) Eval(string_val string) string {
	v8context := C.cef_v8context_get_current_context()
	val := CEFString(string_val)
	retval := C.cef_v8value_create_null()
	go_AddRef(unsafe.Pointer(retval))
	var exception *C.cef_v8exception_t
	exception = (*C.struct__cef_v8exception_t)(C.calloc(1, C.sizeof_struct__cef_v8exception_t))
	go_AddRef(unsafe.Pointer(exception))
	C.V8Eval(v8context, val, &retval, &exception)
	ret := V8ValueToString(retval)
	return ret
}

func (b Browser) SendProcessMessage(target_process C.cef_process_id_t,
	message *C.struct__cef_process_message_t) int {
	return int(C.SendProcessMessage(b.cbrowser, target_process, message))
}

func (b Browser) SendProcessMessageTest() {
	// PID_BROWSER, PID_RENDERER
	message := CefProcessMessageCreate(kV8EvalMessage)
	valueList := message.GetArgumentList()
	s := `function() { return 10; }();`
	//valueList.SetInt(0, 10)
	valueList.SetString(0, s)
	b.SendProcessMessage(C.PID_RENDERER, message.CStruct)
}

var dacCond = sync.NewCond(new(sync.Mutex))

func DeleteAllCookiesExec() {
	C.cef_cookie_manager_t_delete_all_cookies()
	dacCond.L.Lock()
	dacCond.Broadcast()
	dacCond.L.Unlock()
	fmt.Printf("DeleteAllCookiesExec.\n")
}

func (b Browser) DeleteAllCookies() {
	PostTask(TID_IO, DeleteAllCookiesExec)
	dacCond.L.Lock()
	dacCond.Wait()
	dacCond.L.Unlock()
	fmt.Printf("DeleteAllCookies.\n")
}

func (b Browser) Close() {
	C.CloseBrowser(b.cbrowser)
	delete(browsers, b.Id)
}

type BrowserSettings struct {
	///
	// Controls whether file URLs will have access to all URLs. Also configurable
	// using the "allow-universal-access-from-files" command-line switch.
	///
	UniversalAccessFromFileUrls bool

	///
	// Controls whether file URLs will have access to other file URLs. Also
	// configurable using the "allow-access-from-files" command-line switch.
	///
	FileAccessFromFileUrls bool

	///
	// Controls whether web security restrictions (same-origin policy) will be
	// enforced. Disabling this setting is not recommend as it will allow risky
	// security behavior such as cross-site scripting (XSS). Also configurable
	// using the "disable-web-security" command-line switch.
	///
	WebSecurity bool
	///
	// Controls whether WebGL can be used. Note that WebGL requires hardware
	// support and may not work on all systems even when enabled. Also
	// configurable using the "disable-webgl" command-line switch.
	///
	Webgl bool

	///
	// Controls whether content that depends on accelerated compositing can be
	// used. Note that accelerated compositing requires hardware support and may
	// not work on all systems even when enabled. Also configurable using the
	// "disable-accelerated-compositing" command-line switch.
	///
	AcceleratedCompositing bool
}

func (b BrowserSettings) ToCStruct() (cefBrowserSettings *C.struct__cef_browser_settings_t) {
	// Initialize cef_browser_settings_t structure.
	cefBrowserSettings = (*C.struct__cef_browser_settings_t)(C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	cefBrowserSettings.size = C.sizeof_struct__cef_browser_settings_t

	cefBrowserSettings.universal_access_from_file_urls = cefStateFromBool(b.UniversalAccessFromFileUrls)
	cefBrowserSettings.file_access_from_file_urls = cefStateFromBool(b.FileAccessFromFileUrls)
	cefBrowserSettings.web_security = cefStateFromBool(b.WebSecurity)
	cefBrowserSettings.webgl = cefStateFromBool(b.Webgl)
	//cefBrowserSettings.browser_subprocess_path = *CEFString(b.BrowserSubprocessPath)
	//cefBrowserSettings.accelerated_compositing = cefStateFromBool(b.AcceleratedCompositing)
	//cefBrowserSettings.background_color = 255
	return cefBrowserSettings
}
