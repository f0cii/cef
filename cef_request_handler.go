package cef

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "include/capi/cef_load_handler_capi.h"
#include "include/capi/cef_render_handler_capi.h"
#include "include/capi/cef_request_handler_capi.h"
#include "include/capi/cef_resource_handler_capi.h"
#include "include/internal/cef_string_map.h"
#include "cef_helpers.h"
*/
import "C"

import (
	//"fmt"
	"unsafe"
)

var Header map[string]string

func init() {
	Header = make(map[string]string)
}

//export go_RequestHandlerOnBeforeResourceLoad
func go_RequestHandlerOnBeforeResourceLoad(
	self *C.cef_request_handler_t, browser *C.cef_browser_t,
	frame *C.cef_frame_t, request *C.cef_request_t) int {
	var headerMap C.cef_string_map_t = C.cef_string_map_alloc()
	var newHeaderMap C.cef_string_map_t = C.cef_string_map_alloc()

	C.cef_request_t_get_header_map(request, C.cef_string_multimap_t(headerMap))
	//fmt.Printf("go_RequestHandlerOnBeforeResourceLoad %v\n", headerMap)

	headers := TransferStringMapContents(headerMap)
	headers["Accept-Language"] = "zh-CN,zh;q=0.8"

	for k, v := range Header {
		headers[k] = v
	}
	//fmt.Printf("headers=%v\n", headers)

	TransferStringMapContentsToC(headers, newHeaderMap)
	C.cef_request_t_set_header_map(request, C.cef_string_multimap_t(newHeaderMap))

	C.cef_string_map_free(headerMap)
	C.cef_string_map_free(newHeaderMap)
	return 0
}

func TransferStringMapContents(fromMap C.cef_string_map_t) map[string]string {
	output := make(map[string]string)
	size := C.cef_string_map_size(fromMap)

	key := (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	value := (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	defer C.free(unsafe.Pointer(key))
	defer C.free(unsafe.Pointer(value))

	for i := 0; i < int(size); i++ {
		C.cef_string_map_key(fromMap, C.int(i), key)
		C.cef_string_map_value(fromMap, C.int(i), value)
		k := CEFToGoString(key)
		v := CEFToGoString(value)
		//fmt.Printf("~~~~~~~~~~~~~~~key=%v value=%v\n", k, v)
		output[k] = v
	}

	return output
}

func TransferStringMapContentsToC(fromMap map[string]string, toMap C.cef_string_map_t) {
	for k, v := range fromMap {
		key := CEFString(k)
		value := CEFString(v)
		C.cef_string_map_append(toMap, key, value)
	}
}
