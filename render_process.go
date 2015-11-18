package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "cef_helpers.h"
*/
import "C"

//export go_RenderProcessHandlerOnProcessMessageReceived
func go_RenderProcessHandlerOnProcessMessageReceived(self *C.cef_render_process_handler_t,
	browser *C.cef_browser_t, source_process C.cef_process_id_t,
	message *C.cef_process_message_t) int {
	logger.Println("go_RenderProcessHandlerOnProcessMessageReceived")
	m := NewCefProcessMessage(message)
	name := m.GetName()
	logger.Printf("go_RenderProcessHandlerOnProcessMessageReceived name=%v\n", name)
	switch name {
	case kV8EvalMessage:
		on_V8EvalMessage_received(browser, m)
	}
	return 1
}

func on_V8EvalMessage_received(browser *C.cef_browser_t, m *CefProcessMessage) {
	s := m.GetArgumentList().GetString(0)
	logger.Printf("go_RenderProcessHandlerOnProcessMessageReceived s=%v\n", s)
	b := &Browser{0, browser, nil}
	c := b.GetV8Context()
	logger.Printf("go_RenderProcessHandlerOnProcessMessageReceived c=%v\n", c)
	//b.Eval(s)
	logger.Printf("go_RenderProcessHandlerOnProcessMessageReceived d=0")

	message := CefProcessMessageCreate(kV8EvalResultMessage)
	valueList := message.GetArgumentList()
	result := `10`
	//valueList.SetInt(0, 10)
	valueList.SetString(0, result)
	b.SendProcessMessage(C.PID_BROWSER, message.CStruct)
}
