#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

int cef_v8value_t_is_valid(cef_v8value_t* self)
{
    return self->is_valid(self);
}

int cef_v8value_t_is_undefined(cef_v8value_t* self)
{
    return self->is_undefined(self);
}

int cef_v8value_t_is_null(cef_v8value_t* self)
{
    return self->is_null(self);
}

int cef_v8value_t_is_bool(cef_v8value_t* self)
{
    return self->is_bool(self);
}

int cef_v8value_t_is_int(cef_v8value_t* self)
{
    return self->is_int(self);
}

int cef_v8value_t_is_uint(cef_v8value_t* self)
{
    return self->is_uint(self);
}

int cef_v8value_t_is_double(cef_v8value_t* self)
{
    return self->is_double(self);
}

int cef_v8value_t_is_date(cef_v8value_t* self)
{
    return self->is_date(self);
}

int cef_v8value_t_is_string(cef_v8value_t* self)
{
    return self->is_string(self);
}

int cef_v8value_t_is_object(cef_v8value_t* self)
{
    return self->is_object(self);
}

int cef_v8value_t_is_array(cef_v8value_t* self)
{
    return self->is_array(self);
}

int cef_v8value_t_is_function(cef_v8value_t* self)
{
    return self->is_function(self);
}

int cef_v8value_t_is_same(cef_v8value_t* self, cef_v8value_t* that)
{
    return self->is_same(self, that);
}

//int (CEF_CALLBACK *get_bool_value)(struct _cef_v8value_t* self);

//int32 (CEF_CALLBACK *get_int_value)(struct _cef_v8value_t* self);

//uint32 (CEF_CALLBACK *get_uint_value)(struct _cef_v8value_t* self);

//double (CEF_CALLBACK *get_double_value)(struct _cef_v8value_t* self);

//cef_time_t (CEF_CALLBACK *get_date_value)(struct _cef_v8value_t* self);

///
// Return a string value.  The underlying data will be converted to if
// necessary.
///
// The resulting string must be freed by calling cef_string_userfree_free().
//cef_string_userfree_t (CEF_CALLBACK *get_string_value)(
//  struct _cef_v8value_t* self);

cef_string_utf8_t * cefSourceToString(cef_string_t * source) {
      cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
      if (source == 0) {
          return output;
      }
      cef_string_to_utf8(source->str, source->length, output);
      return output;
}

cef_string_userfree_t v8ValueToString(cef_v8value_t * str) {
      return str->get_string_value(str);
}

int32 v8ValueToInt32(cef_v8value_t * i) {
      return i->get_int_value(i);
}

uint32 v8ValueToUInt32(cef_v8value_t* v) {
   return v->get_uint_value(v);
}

int v8ValueToBool(cef_v8value_t * b) {
      return b->get_bool_value(b);
}

double v8ValueToDouble(cef_v8value_t * b) {
    return b->get_double_value(b);
}

int v8ValueGetArrayLength(cef_v8value_t * v) {
    return v->get_array_length(v);
}

cef_v8value_t* v8ValueGetValueByIndex(cef_v8value_t* self, int index)
{
    return self->get_value_byindex(self, index);
}

void setCefRectDimensions(cef_rect_t * rect, int x, int y, int width, int height) {
      rect->x = x;
      rect->y = y;
      rect->width = width;
      rect->height = height;
}

struct _cef_browser_host_t* cef_browser_t_get_host(struct _cef_browser_t* self) {
    return self->get_host(self);
}

/*cef_window_handle_t cef_browser_host_t_get_window_handle(
  struct _cef_browser_host_t* self)
{
    return self->get_window_handle(self);
}*/

/*cef_window_handle_t cef_browser_host_t_get_opener_window_handle(
  struct _cef_browser_host_t* self)
{
    return self->get_opener_window_handle(self);
}*/

void cef_cookie_manager_t_delete_cookies(const char* url, const char* cookie_name)
{
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_t * cookieNameVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), urlVal);
    cef_string_from_utf8(cookie_name, strlen(cookie_name), cookieNameVal);

    cef_cookie_manager_t* manager = cef_cookie_manager_get_global_manager();
    manager->delete_cookies(manager, urlVal, cookieNameVal);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(cookieNameVal);
}

void cef_cookie_manager_t_delete_all_cookies()
{
    cef_cookie_manager_t_delete_cookies("", "");
}

int64 int64_array_get(int64* aa, size_t i) {
    return aa[i];
}

void cef_request_t_get_header_map(struct _cef_request_t* self, cef_string_multimap_t headerMap) {
    self->get_header_map(self, headerMap);
}

void cef_request_t_set_header_map(struct _cef_request_t* self, cef_string_multimap_t headerMap) {
    self->set_header_map(self, headerMap);
}