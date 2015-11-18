package cef

/*
#include <stdlib.h>
#include "string.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"
*/
import "C"
import (
//"log"
//"os"
//"time"
//"unsafe"
)

// C.cef_key_event_type_t
const (
	KEYEVENT_RAWKEYDOWN = 0
	KEYEVENT_KEYDOWN    = 1
	KEYEVENT_KEYUP      = 2
	KEYEVENT_CHAR       = 3
)

const (
	MBT_LEFT   = 0
	MBT_MIDDLE = 1
	MBT_RIGHT  = 2
)

type MouseButtonType C.cef_mouse_button_type_t

type CefKeyEvent struct {
	Type                int // C.cef_key_event_type_t
	Modifiers           uint32
	WindowsKeyCode      int
	NativeKeyCode       int
	Character           int
	UnmodifiedCharacter int
	IsSystemKey         bool
}

type CefMouseEvent struct {
	X         int
	Y         int
	Modifiers uint32
}
