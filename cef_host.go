package cef

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "cef_host.h"
*/
import "C"

import (
//"unsafe"
)

type BrowserHost struct {
	cbrowserhost *C.struct__cef_browser_host_t
}

func (h *BrowserHost) SetFocus(focus bool) {
	var nFocus C.int = 0
	if focus {
		nFocus = 1
	}
	C.cef_browser_host_t_set_focus(h.cbrowserhost, nFocus)
}

func (h *BrowserHost) SendKeyEvent(event *CefKeyEvent) {
	var keyEvent C.cef_key_event_t
	switch event.Type {
	case KEYEVENT_RAWKEYDOWN:
		keyEvent._type = C.KEYEVENT_RAWKEYDOWN
	case KEYEVENT_KEYDOWN:
		keyEvent._type = C.KEYEVENT_KEYDOWN
	case KEYEVENT_KEYUP:
		keyEvent._type = C.KEYEVENT_KEYUP
	case KEYEVENT_CHAR:
		keyEvent._type = C.KEYEVENT_CHAR
	}
	keyEvent.windows_key_code = C.int(event.WindowsKeyCode)
	keyEvent.native_key_code = C.int(event.NativeKeyCode)
	if event.IsSystemKey {
		keyEvent.is_system_key = 1
	} else {
		keyEvent.is_system_key = 0
	}
	keyEvent.character = C.char16(event.Character)
	keyEvent.unmodified_character = C.char16(event.UnmodifiedCharacter)
	keyEvent.modifiers = C.uint32(event.Modifiers)
	keyEvent.focus_on_editable_field = 0
	C.cef_browser_host_t_send_key_event(h.cbrowserhost, &keyEvent)
}

/*
    // 鼠标左键点击
	var event CefMouseEvent
	event.X = 120
	event.Y = 160
	event.Modifiers = 0
	h.SendMouseClickEvent(&event, MBT_LEFT, false, 1)
	h.SendMouseClickEvent(&event, MBT_LEFT, true, 1)
*/
func (h *BrowserHost) SendMouseClickEvent(event *CefMouseEvent, buttonType int, mouseUp bool, clickCount int) {
	var evt C.cef_mouse_event_t
	var button_type C.cef_mouse_button_type_t
	var mouse_up int

	evt.x = C.int(event.X)
	evt.y = C.int(event.Y)
	evt.modifiers = C.uint32(event.Modifiers)

	switch buttonType {
	case MBT_LEFT:
		button_type = C.MBT_LEFT
	case MBT_MIDDLE:
		button_type = C.MBT_MIDDLE
	case MBT_RIGHT:
		button_type = C.MBT_RIGHT
	}

	if mouseUp {
		mouse_up = 1
	} else {
		mouse_up = 0
	}
	C.cef_browser_host_t_send_mouse_click_event(h.cbrowserhost, &evt, button_type, C.int(mouse_up), C.int(clickCount))
}

func (h *BrowserHost) SendMouseMoveEvent(event *CefMouseEvent) {
	var evt C.cef_mouse_event_t

	evt.x = C.int(event.X)
	evt.y = C.int(event.Y)

	C.cef_browser_host_t_send_mouse_move_event(h.cbrowserhost, &evt, C.int(0))
}
