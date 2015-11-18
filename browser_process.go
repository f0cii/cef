package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "cef_helpers.h"
*/
import "C"

import "fmt"

//export go_BrowserProcessHandlerOnContextInitialized
func go_BrowserProcessHandlerOnContextInitialized(self *C.cef_browser_process_handler_t) {
	fmt.Println("go_BrowserProcessHandlerOnContextInitialized")
	if OnContextInitialized != nil {
		OnContextInitialized()
	}
}
