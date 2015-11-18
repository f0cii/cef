#include <string.h>
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_browser_process_handler_capi.h"
#include "include/capi/cef_render_process_handler_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_base.h"

// ----------------------------------------------------------------------------
// cef_app_t
// ----------------------------------------------------------------------------


int CEF_CALLBACK cef_v8handler_execute(struct _cef_v8handler_t* self,
      const cef_string_t* name, struct _cef_v8value_t* object,
      size_t argumentsCount, struct _cef_v8value_t* const* arguments,
      struct _cef_v8value_t** retval, cef_string_t* exception) {
    DEBUG_CALLBACK("v8handler->execute");
    cef_browser_t* browser = NULL;
    cef_v8context_t* context = cef_v8context_get_current_context();
    if (context) {
        browser = context->get_browser(context);
        DEBUG_CALLBACK("cef_v8handler_execute");
    }
    //return go_V8HandlerExecute(name, object, argumentsCount, arguments, retval, exception);
    return go_V8HandlerExecute(browser->get_identifier(browser), browser, name, object, argumentsCount, arguments, retval, exception);
}

// Set up the javascript cef extensions
void CEF_CALLBACK cef_render_process_handler_t_on_webkit_initialized(struct _cef_render_process_handler_t* self) {
    cef_v8handler_t* goV8Handler = (cef_v8handler_t*)calloc(1, sizeof(cef_v8handler_t));
    goV8Handler->base.size = sizeof(cef_v8handler_t);
    initialize_cef_base((cef_base_t*) goV8Handler);
    goV8Handler->execute = cef_v8handler_execute;
    go_RenderProcessHandlerOnWebKitInitialized(goV8Handler);
}

///
// Called when a new message is received from a different process. Return true
// (1) if the message was handled or false (0) otherwise. Do not keep a
// reference to or attempt to access the message outside of this callback.
///
int CEF_CALLBACK cef_render_process_handler_t_on_process_message_received(
    struct _cef_render_process_handler_t* self,
    struct _cef_browser_t* browser, cef_process_id_t source_process,
    struct _cef_process_message_t* message) {
    DEBUG_CALLBACK("cef_render_process_handler_t_on_process_message_received");
    return go_RenderProcessHandlerOnProcessMessageReceived(self, browser, source_process, message);
}

///
// Called on the browser process UI thread immediately after the CEF context
// has been initialized.
///
void CEF_CALLBACK cef_browser_process_handler_t_on_context_initialized(
    struct _cef_browser_process_handler_t* self)
{
    go_BrowserProcessHandlerOnContextInitialized(self);
}

///
// Called before a child process is launched. Will be called on the browser
// process UI thread when launching a render process and on the browser
// process IO thread when launching a GPU or plugin process. Provides an
// opportunity to modify the child process command line. Do not keep a
// reference to |command_line| outside of this function.
///
void CEF_CALLBACK cef_browser_process_handler_t_on_before_child_process_launch(
    struct _cef_browser_process_handler_t* self,
    struct _cef_command_line_t* command_line) {

}

///
// Called on the browser process IO thread after the main thread has been
// created for a new render process. Provides an opportunity to specify extra
// information that will be passed to
// cef_render_process_handler_t::on_render_thread_created() in the render
// process. Do not keep a reference to |extra_info| outside of this function.
///
void CEF_CALLBACK cef_browser_process_handler_t_on_render_process_thread_created(
    struct _cef_browser_process_handler_t* self,
    struct _cef_list_value_t* extra_info) {

}

///
// Return the handler for printing on Linux. If a print handler is not
// provided then printing will not be supported on the Linux platform.
///
struct _cef_print_handler_t* CEF_CALLBACK cef_browser_process_handler_t_get_print_handler(
    struct _cef_browser_process_handler_t* self) {
    return 0;
}

struct _cef_browser_process_handler_t*
      CEF_CALLBACK get_browser_process_handler(struct _cef_app_t* self) {
    //DEBUG_POINTER("get_browser_process_handler", self);
    cef_browser_process_handler_t* browserProcessHandler = (cef_browser_process_handler_t*)calloc(1, sizeof(cef_browser_process_handler_t));
    browserProcessHandler->base.size = sizeof(cef_browser_process_handler_t);
    initialize_cef_base((cef_base_t*) browserProcessHandler);
    browserProcessHandler->on_context_initialized = cef_browser_process_handler_t_on_context_initialized;
    browserProcessHandler->on_before_child_process_launch = cef_browser_process_handler_t_on_before_child_process_launch;
    browserProcessHandler->on_render_process_thread_created = cef_browser_process_handler_t_on_render_process_thread_created;
    browserProcessHandler->get_print_handler = cef_browser_process_handler_t_get_print_handler;
    return browserProcessHandler;
}

///
// Return the handler for functionality specific to the render process. This
// function is called on the render process main thread.
///
struct _cef_render_process_handler_t*
        CEF_CALLBACK get_render_process_handler(struct _cef_app_t* self) {
    //DEBUG_POINTER("get_render_process_handler", self);
    cef_render_process_handler_t* renderProcessHandler = (cef_render_process_handler_t*)calloc(1, sizeof(cef_render_process_handler_t));
    renderProcessHandler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base((cef_base_t*) renderProcessHandler);
    renderProcessHandler->on_web_kit_initialized = cef_render_process_handler_t_on_webkit_initialized;
    renderProcessHandler->on_process_message_received = cef_render_process_handler_t_on_process_message_received;
    return renderProcessHandler;
}

void initialize_app_handler(cef_app_t* app) {
    DEBUG_POINTER("initialize_app_handler", app);
    app->base.size = sizeof(cef_app_t);
    initialize_cef_base((cef_base_t*)app);
    // callbacks
    app->get_browser_process_handler = get_browser_process_handler;
    app->get_render_process_handler = get_render_process_handler;
}
