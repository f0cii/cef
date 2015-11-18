#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

extern void cef_frame_t_execute_java_script(cef_frame_t* frame, const char* code, const char* script_url, int start_line);
extern int cef_frame_t_is_valid(cef_frame_t* self);
extern cef_string_t * cef_frame_t_get_url(cef_frame_t* frame);
extern void cef_frame_t_get_source(cef_frame_t* frame, cef_string_visitor_t* visitor);
extern void cef_frame_t_get_text(cef_frame_t* frame, cef_string_visitor_t* visitor);
extern cef_browser_t* cef_frame_t_get_browser(cef_frame_t* self);

