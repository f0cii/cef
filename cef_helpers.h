#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

extern int cef_v8value_t_is_valid(cef_v8value_t* self);
extern int cef_v8value_t_is_undefined(cef_v8value_t* self);
extern int cef_v8value_t_is_null(cef_v8value_t* self);
extern int cef_v8value_t_is_bool(cef_v8value_t* self);
extern int cef_v8value_t_is_int(cef_v8value_t* self);
extern int cef_v8value_t_is_uint(cef_v8value_t* self);
extern int cef_v8value_t_is_double(cef_v8value_t* self);
extern int cef_v8value_t_is_date(cef_v8value_t* self);
extern int cef_v8value_t_is_string(cef_v8value_t* self);
extern int cef_v8value_t_is_object(cef_v8value_t* self);
extern int cef_v8value_t_is_array(cef_v8value_t* self);
extern int cef_v8value_t_is_function(cef_v8value_t* self);
extern int cef_v8value_t_is_same(cef_v8value_t* self, cef_v8value_t* that);

extern cef_string_utf8_t * cefSourceToString(cef_string_t * source);
extern cef_string_userfree_t v8ValueToString(cef_v8value_t * str);
extern int32 v8ValueToInt32(cef_v8value_t * i);
extern uint32 v8ValueToUInt32(cef_v8value_t* v);
extern int v8ValueToBool(cef_v8value_t * b);
extern int v8ValueToDouble(cef_v8value_t * b);
extern int v8ValueGetArrayLength(cef_v8value_t * v);
extern cef_v8value_t* v8ValueGetValueByIndex(cef_v8value_t* self, int index);
extern void setCefRectDimensions(cef_rect_t * rect, int x, int y, int width, int height);

extern struct _cef_browser_host_t* cef_browser_t_get_host(struct _cef_browser_t* self);

extern cef_window_handle_t cef_browser_host_t_get_window_handle(
  struct _cef_browser_host_t* self);

extern cef_window_handle_t cef_browser_host_t_get_opener_window_handle(
  struct _cef_browser_host_t* self);

extern cef_window_handle_t GetWindowHandle(cef_browser_t* browser);

extern void ExecuteJavaScript(cef_browser_t* browser, const char* code, const char* script_url, int start_line);

extern void cef_cookie_manager_t_delete_cookies(const char* url, const char* cookie_name);

extern void cef_cookie_manager_t_delete_all_cookies();

extern int64 int64_array_get(int64* aa, size_t i);

extern void cef_request_t_get_header_map(struct _cef_request_t* self, cef_string_multimap_t headerMap);

extern void cef_request_t_set_header_map(struct _cef_request_t* self, cef_string_multimap_t headerMap);

