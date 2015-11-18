#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

void ExecuteJavaScript(cef_browser_t* browser, const char* code, const char* script_url, int start_line)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * codeCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(code, strlen(code), codeCef);
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(script_url, strlen(script_url), urlVal);

    frame->execute_java_script(frame, codeCef, urlVal, start_line);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(codeCef);
}

void LoadURL(cef_browser_t* browser, const char* url)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * urlCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), urlCef);
    frame->load_url(frame, urlCef);
    cef_string_userfree_utf16_free(urlCef);
}

void _LoadString(cef_browser_t* browser, const char* string_val, const char* url)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    cef_string_t * valCef = cef_string_userfree_utf16_alloc();
    cef_string_t * urlCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(string_val), valCef);
    cef_string_from_utf8(url, strlen(url), urlCef);
    frame->load_string(frame, valCef, urlCef);
    cef_string_userfree_utf16_free(valCef);
    cef_string_userfree_utf16_free(urlCef);
}

void BrowserWasResized(cef_browser_t* browser)
{
    cef_browser_host_t * host = browser->get_host(browser);
    host->was_resized(host);
}

cef_window_handle_t GetWindowHandle(cef_browser_t* browser)
{
    cef_browser_host_t * host = browser->get_host(browser);
    return host->get_window_handle(host);
}

cef_window_handle_t GetRootWindowHandle(cef_browser_t* browser)
{
     cef_browser_host_t * host = browser->get_host(browser);
     return GetAncestor(host->get_window_handle(host), GA_ROOT);
}

// Force close the browser
void CloseBrowser(cef_browser_t* browser)
{
    cef_browser_host_t * host = browser->get_host(browser);
    host->close_browser(host, 1);
}

cef_string_utf8_t * cefStringToUtf8(cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    if (source == 0) {
        return output;
    }
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}

int cef_browser_t_get_identifier(cef_browser_t* browser)
{
    return browser->get_identifier(browser);
}

cef_frame_t * cef_browser_t_get_main_frame(cef_browser_t* browser)
{
    return browser->get_main_frame(browser);
}

cef_frame_t * cef_browser_t_get_focused_frame(cef_browser_t* browser)
{
    return browser->get_focused_frame(browser);
}

cef_frame_t * cef_browser_t_get_frame_byident(cef_browser_t* browser, int64 identifier)
{
    return browser->get_frame_byident(browser, identifier);
}

cef_frame_t * cef_browser_t_get_frame(cef_browser_t* browser, const cef_string_t* name)
{
    return browser->get_frame(browser, name);
}

size_t cef_browser_t_get_frame_count(cef_browser_t* browser)
{
    return browser->get_frame_count(browser);
}

void cef_browser_t_get_frame_identifiers(struct _cef_browser_t* self,
    size_t* identifiersCount, int64* identifiers)
{
    self->get_frame_identifiers(self, identifiersCount, identifiers);
}

void cef_browser_t_get_frame_names(struct _cef_browser_t* self,
    cef_string_list_t names)
{
    self->get_frame_names(self, names);
}

cef_string_t * GetURL(cef_browser_t* browser)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    return frame->get_url(frame);
}

void GetSource(cef_browser_t* browser, cef_string_visitor_t* visitor)
{
	cef_frame_t * frame = browser->get_main_frame(browser);
	frame->get_source(frame, visitor);
}

void GetText(cef_browser_t* browser, cef_string_visitor_t* visitor)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    frame->get_text(frame, visitor);
}

struct _cef_v8context_t* GetV8Context(cef_browser_t* browser)
{
    cef_frame_t * frame = browser->get_main_frame(browser);
    return frame->get_v8context(frame);
}

int V8Eval(struct _cef_v8context_t* self,
      const cef_string_t* code, struct _cef_v8value_t** retval,
      struct _cef_v8exception_t** exception)
{
    return self->eval(self, code, retval, exception);
}

int SendProcessMessage(struct _cef_browser_t* self,
      cef_process_id_t target_process,
      struct _cef_process_message_t* message)
{
    return self->send_process_message(self, target_process, message);
}