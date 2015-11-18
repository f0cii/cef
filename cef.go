// Website: https://github.com/CzarekTomczak/cef2go
// Website: https://github.com/lroc/cef

package cef

/*
CEF capi fixes
--------------
1. In cef_string.h:
    this => typedef cef_string_utf16_t cef_string_t;
    to => #define cef_string_t cef_string_utf16_t
2. In cef_export.h:
    #elif defined(COMPILER_GCC)
    #define CEF_EXPORT __attribute__ ((visibility("default")))
    #ifdef OS_WIN
    #define CEF_CALLBACK __stdcall
    #else
    #define CEF_CALLBACK
    #endif
*/

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./Release -lcef_sandbox
#include <stdlib.h>
#include "string.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"
*/
import "C"
import (
	"log"
	"os"
	"time"
	"unsafe"
	//"fmt"
)

var logger *log.Logger = log.New(os.Stdout, "[cef] ", log.Lshortfile)

var _MainArgs *C.struct__cef_main_args_t
var _AppHandler *C.cef_app_t               // requires reference counting
var _ClientHandler *C.struct__cef_client_t // requires reference counting

// Set up the js console handlers
type ConsoleHandlerFunc func(message, source string, line int)

var DefaultConsoleHandler ConsoleHandlerFunc = ConsoleHandlerFunc(func(message, source string, line int) {
	logger.Printf("[console:%s %d] %s\n", source, line, message)
})
var consoleHandler ConsoleHandlerFunc = DefaultConsoleHandler

// Sandbox is disabled. Including the "cef_sandbox.lib"
// library results in lots of GCC warnings/errors. It is
// compatible only with VS 2010. It would be required to
// build it using GCC. Add -lcef_sandbox to LDFLAGS.
// capi doesn't expose sandbox functions, you need do add
// these before import "C":
// void* cef_sandbox_info_create();
// void cef_sandbox_info_destroy(void* sandbox_info);
var _SandboxInfo unsafe.Pointer

type Settings struct {
	SingleProcess           int
	BrowserSubprocessPath   string
	CommandLineArgsDisabled int
	UserAgent               string
	CachePath               string
	LogSeverity             int
	LogFile                 string
	Locale                  string
	ResourcesDirPath        string
	LocalesDirPath          string
	JavaScriptFlags         string
	RemoteDebuggingPort     int
}

const (
	LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
	LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO    = C.LOGSEVERITY_INFO
	LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
	LOGSEVERITY_ERROR   = C.LOGSEVERITY_ERROR
	//LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
	LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE

	kV8EvalMessage       = "ClientRenderer.V8Eval"
	kV8EvalResultMessage = "ClientRenderer.V8EvalResult"
	kV8CallbackMessage   = "ClientRenderer.V8Callback"
	kV8OpenWindowMessage = "ClientRenderer.V8OpenWindow"
)

func init() {

}

func CEFString(original string) (final *C.cef_string_t) {
	final = (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	charString := C.CString(original)
	defer C.free(unsafe.Pointer(charString))
	C.cef_string_from_utf8(charString, C.strlen(charString), final)
	return final
}

func CEFToGoString(source *C.cef_string_t) string {
	utf8string := C.cefSourceToString(source)
	defer C.cef_string_userfree_utf8_free(utf8string)
	return C.GoString(utf8string.str)
}

func _InitializeGlobalCStructures() {
	_MainArgs = (*C.struct__cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	go_AddRef(unsafe.Pointer(_MainArgs))

	_AppHandler = (*C.cef_app_t)(C.calloc(1, C.sizeof_cef_app_t))
	go_AddRef(unsafe.Pointer(_AppHandler))
	C.initialize_app_handler(_AppHandler)

	_ClientHandler = (*C.struct__cef_client_t)(C.calloc(1, C.sizeof_struct__cef_client_t))
	go_AddRef(unsafe.Pointer(_ClientHandler))
	C.initialize_client_handler(_ClientHandler)
}

func ExecuteProcess(appHandle unsafe.Pointer) int {
	logger.Printf("ExecuteProcess, args=%v\n", os.Args)

	_InitializeGlobalCStructures()
	FillMainArgs(_MainArgs, appHandle)

	// Sandbox info needs to be passed to both cef_execute_process()
	// and cef_initialize().
	// OFF: _SandboxInfo = C.cef_sandbox_info_create()

	var exitCode C.int = C.cef_execute_process(_MainArgs, _AppHandler, _SandboxInfo)
	if exitCode >= 0 {
		os.Exit(int(exitCode))
	}
	logger.Printf("Finished ExecuteProcess, args=%v %d %d\n", os.Args, os.Getpid(), exitCode)
	return int(exitCode)
}

func cefStateFromBool(state bool) C.cef_state_t {
	if state == true {
		return C.STATE_ENABLED
	} else {
		return C.STATE_DISABLED
	}
}

func (settings *Settings) ToCStruct() (cefSettings *C.struct__cef_settings_t) {
	// Initialize cef_settings_t structure.
	cefSettings = (*C.struct__cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))
	cefSettings.size = C.sizeof_struct__cef_settings_t
	cefSettings.single_process = C.int(settings.SingleProcess)
	cefSettings.cache_path = *CEFString(settings.CachePath)
	cefSettings.log_severity = (C.cef_log_severity_t)(C.int(settings.LogSeverity))
	cefSettings.log_file = *CEFString(settings.LogFile)
	cefSettings.resources_dir_path = *CEFString(settings.ResourcesDirPath)
	cefSettings.locales_dir_path = *CEFString(settings.LocalesDirPath)
	if len(settings.BrowserSubprocessPath) > 0 {
		cefSettings.browser_subprocess_path = *CEFString(settings.BrowserSubprocessPath)
	}
	cefSettings.command_line_args_disabled = C.int(settings.CommandLineArgsDisabled)
	if len(settings.UserAgent) > 0 {
		cefSettings.user_agent = *CEFString(settings.UserAgent)
	}
	cefSettings.remote_debugging_port = C.int(settings.RemoteDebuggingPort)
	cefSettings.javascript_flags = *CEFString(settings.JavaScriptFlags)
	cefSettings.locale = *CEFString(settings.Locale)
	cefSettings.no_sandbox = C.int(1)
	return
}

func Initialize(settings Settings) int {
	logger.Println("Initialize")

	if _MainArgs == nil {
		// _MainArgs structure is initialized and filled in ExecuteProcess.
		// If cef_execute_process is not called, and there is a call
		// to cef_initialize, then it would result in creation of infinite
		// number of processes. See Issue 1199 in CEF:
		// https://code.google.com/p/chromiumembedded/issues/detail?id=1199
		logger.Println("ERROR: missing a call to ExecuteProcess")
		return 0
	}

	globalLifespanHandler = &LifeSpanHandler{} //make(chan *Browser)
	go_AddRef(unsafe.Pointer(_AppHandler))
	ret := C.cef_initialize(_MainArgs, settings.ToCStruct(), _AppHandler, _SandboxInfo)
	// Sleep for 1500ms to let cef _really_ initialize
	// https://code.google.com/p/cefpython/issues/detail?id=131#c2
	time.Sleep(1500 * time.Millisecond)

	return int(ret)
}

func RunMessageLoop() {
	logger.Println("RunMessageLoop")
	C.cef_run_message_loop()
}

func QuitMessageLoop() {
	logger.Println("QuitMessageLoop")
	C.cef_quit_message_loop()
}

func Shutdown() {
	logger.Println("Shutdown")
	C.cef_shutdown()
}

func SetConsoleHandler(handler ConsoleHandlerFunc) {
	consoleHandler = handler
}

func OnUIThread() bool {
	return C.cef_currently_on(C.TID_UI) == 1
}
