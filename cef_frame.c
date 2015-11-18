#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

void cef_frame_t_execute_java_script(cef_frame_t* frame, const char* code, const char* script_url, int start_line)
{
    cef_string_t * codeCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(code, strlen(code), codeCef);
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(script_url, strlen(script_url), urlVal);

    frame->execute_java_script(frame, codeCef, urlVal, start_line);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(codeCef);
}

int cef_frame_t_is_valid(cef_frame_t* self)
{
    return self->is_valid(self);
}

cef_string_t * cef_frame_t_get_url(cef_frame_t* frame)
{
    return frame->get_url(frame);
}

void cef_frame_t_get_source(cef_frame_t* frame, cef_string_visitor_t* visitor)
{
	frame->get_source(frame, visitor);
}

void cef_frame_t_get_text(cef_frame_t* frame, cef_string_visitor_t* visitor)
{
    frame->get_text(frame, visitor);
}

cef_browser_t* cef_frame_t_get_browser(cef_frame_t* self)
{
    return self->get_browser(self);
}