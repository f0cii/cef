package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "cef_helpers.h"
*/
import "C"

import (
	"unsafe"
)

type CefRect C.cef_rect_t
type CefScreenInfo C.cef_screen_info_t
type CefPaintElementType C.cef_paint_element_type_t
type CefCursorHandle C.cef_cursor_handle_t

func (r *CefRect) SetDimensions(x, y, width, height int) {
	C.setCefRectDimensions((*C.cef_rect_t)(r), (C.int)(x), (C.int)(y), (C.int)(width), (C.int)(height))
}

type RenderHandler interface {
	GetRootScreenRect(*CefRect) int
	GetViewRect(*CefRect) int
	GetScreenPoint(int, int, *int, *int) int
	GetScreenInfo(*CefScreenInfo) int
	OnPopupShow(int)
	OnPopupSize(*CefRect)
	OnPaint(CefPaintElementType, int, unsafe.Pointer, unsafe.Pointer, int, int)
	OnCursorChange(CefCursorHandle)
	OnScrollOffsetChanged()
}

//export go_RenderHandlerGetRootScreenRect
func go_RenderHandlerGetRootScreenRect(browserId int, rect *C.cef_rect_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetRootScreenRect((*CefRect)(rect))
	}
	return 0
}

//export go_RenderHandlerGetViewRect
func go_RenderHandlerGetViewRect(browserId int, rect *C.cef_rect_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetViewRect((*CefRect)(rect))
	}
	return 0
}

//export go_RenderHandlerGetScreenPoint
func go_RenderHandlerGetScreenPoint(browserId, x, y int, screenX *int, screenY *int) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetScreenPoint(x, y, screenX, screenY)
	}
	return 0
}

//export go_RenderHandlerGetScreenInfo
func go_RenderHandlerGetScreenInfo(browserId int, info *C.cef_screen_info_t) int {
	if b, ok := BrowserById(browserId); ok {
		return b.RenderHandler.GetScreenInfo((*CefScreenInfo)(info))
	}
	return 0
}

//export go_RenderHandlerOnPopupShow
func go_RenderHandlerOnPopupShow(browserId int, show int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPopupShow(show)
	}
}

//export go_RenderHandlerOnPopupSize
func go_RenderHandlerOnPopupSize(browserId int, size *C.cef_rect_t) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPopupSize((*CefRect)(size))
	}
}

//export go_RenderHandlerOnPaint
func go_RenderHandlerOnPaint(browserId int, paintType C.cef_paint_element_type_t, dirtyRectsCount C.size_t, dirtyRects unsafe.Pointer, buffer unsafe.Pointer, width, height int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnPaint((CefPaintElementType)(paintType), int(dirtyRectsCount), dirtyRects, buffer, width, height)
	}
}

//export go_RenderHandlerOnCursorChange
func go_RenderHandlerOnCursorChange(browserId int, cursor C.cef_cursor_handle_t) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnCursorChange((CefCursorHandle)(cursor))
	}
}

//export go_RenderHandlerOnScrollOffsetChanged
func go_RenderHandlerOnScrollOffsetChanged(browserId int) {
	if b, ok := BrowserById(browserId); ok {
		b.RenderHandler.OnScrollOffsetChanged()
	}
}

type DefaultRenderHandler struct {
	Browser *Browser
}

func (d *DefaultRenderHandler) GetRootScreenRect(rect *CefRect) int {
	return 0
}

func (d *DefaultRenderHandler) GetViewRect(rect *CefRect) int {
	return 0
}

func (d *DefaultRenderHandler) GetScreenPoint(x, y int, screenX, screenY *int) int {
	return 0
}

func (d *DefaultRenderHandler) GetScreenInfo(info *CefScreenInfo) int {
	return 0
}

func (d *DefaultRenderHandler) OnPopupShow(show int) {
}

func (d *DefaultRenderHandler) OnPopupSize(size *CefRect) {
}

func (d *DefaultRenderHandler) OnPaint(paintType CefPaintElementType, dirtyRectsCount int, dirtyRects unsafe.Pointer, buffer unsafe.Pointer, width, height int) {
}

func (d *DefaultRenderHandler) OnCursorChange(cursor CefCursorHandle) {
}

func (d *DefaultRenderHandler) OnScrollOffsetChanged() {
}
