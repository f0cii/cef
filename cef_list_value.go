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
#include "cef_list_value.h"
*/
import "C"

type CefListValue struct {
	CStruct *C.struct__cef_list_value_t
}

func (l *CefListValue) IsValid() bool {
	return (C.cef_list_value_t_is_valid(l.CStruct) == 1)
}

func (l *CefListValue) IsOwned() bool {
	return (C.cef_list_value_t_is_owned(l.CStruct) == 1)
}

func (l *CefListValue) IsReadOnly() bool {
	return (C.cef_list_value_t_is_read_only(l.CStruct) == 1)
}

func (l *CefListValue) Copy() (o *CefListValue) {
	o.CStruct = C.cef_list_value_t_copy(l.CStruct)
	return
}

func (l *CefListValue) SetSize(size int) {
	C.cef_list_value_t_set_size(l.CStruct, C.size_t(size))
}

func (l *CefListValue) GetSize() int {
	return int(C.cef_list_value_t_get_size(l.CStruct))
}

func (l *CefListValue) Clear() {
	C.cef_list_value_t_clear(l.CStruct)
}

func (l *CefListValue) Remove(index int) {
	C.cef_list_value_t_remove(l.CStruct, C.int(index))
}

func (l *CefListValue) GetType(index int) C.cef_value_type_t {
	return C.cef_list_value_t_get_type(l.CStruct, C.int(index))
}

func (l *CefListValue) GetBool(index int) bool {
	return (C.cef_list_value_t_get_bool(l.CStruct, C.int(index)) == 1)
}

func (l *CefListValue) GetInt(index int) int {
	return int(C.cef_list_value_t_get_int(l.CStruct, C.int(index)))
}

func (l *CefListValue) GetDouble(index int) float64 {
	return float64(C.cef_list_value_t_get_double(l.CStruct, C.int(index)))
}

func (l *CefListValue) GetString(index int) string {
	return CEFToGoString(C.cef_list_value_t_get_string(l.CStruct, C.int(index)))
}

func (l *CefListValue) SetNull(index int) bool {
	return (C.cef_list_value_t_set_null(l.CStruct, C.int(index)) == 1)
}

func (l *CefListValue) SetBool(index int, value bool) bool {
	var v = 0
	if value {
		v = 1
	}
	return (C.cef_list_value_t_set_bool(l.CStruct, C.int(index), C.int(v)) == 1)
}

func (l *CefListValue) SetInt(index int, value int) bool {
	return (C.cef_list_value_t_set_int(l.CStruct, C.int(index), C.int(value)) == 1)
}

func (l *CefListValue) SetDouble(index int, value float64) bool {
	return (C.cef_list_value_t_set_double(l.CStruct, C.int(index), C.double(value)) == 1)
}

func (l *CefListValue) SetString(index int, value string) bool {
	s := CEFString(value)
	return (C.cef_list_value_t_set_string(l.CStruct, C.int(index), s) == 1)
}
