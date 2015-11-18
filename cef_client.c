#include <string.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "include/capi/cef_load_handler_capi.h"
#include "include/capi/cef_render_handler_capi.h"
#include "include/capi/cef_request_handler_capi.h"
#include "include/capi/cef_resource_handler_capi.h"
#include "cef_base.h"

typedef struct cef_go_client {
    
  cef_display_handler_t *display_handler;
  cef_life_span_handler_t *life_span_handler;
  cef_load_handler_t *load_handler;
  cef_render_handler_t *render_handler;
  cef_request_handler_t *request_handler;
  cef_resource_handler_t *resource_handler;

} cef_go_client;

cef_go_client * go_client;

void CEF_CALLBACK cef_display_handler_t_on_address_change(struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, struct _cef_frame_t* frame,
      const cef_string_t* url) {
}

void CEF_CALLBACK cef_display_handler_t_on_title_change(struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, const cef_string_t* title) {
    go_OnTitleChange(browser->get_identifier(browser), browser, title);
}

int CEF_CALLBACK cef_display_handler_t_on_tooltip(struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, cef_string_t* text) {
      return 0;
}

void CEF_CALLBACK cef_display_handler_t_on_status_message(struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, const cef_string_t* value) {
}

int CEF_CALLBACK cef_display_handler_t_on_console_message(
      struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, const cef_string_t* message,
      const cef_string_t* source, int line) {
    go_OnConsoleMessage(browser, message, source, line);
    return 1;
}

int CEF_CALLBACK cef_life_span_handler_t_on_before_popup(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser,
        struct _cef_frame_t* frame,
        const cef_string_t* target_url,
        const cef_string_t* target_frame_name,
        const struct _cef_popup_features_t* popupFeatures,
        struct _cef_window_info_t* windowInfo,
        struct _cef_client_t** client,
        struct _cef_browser_settings_t* settings,
        int* no_javascript_access) {
    return go_OnBeforePopup(
        self,
        browser,
        frame,
        target_url,
        target_frame_name,
        popupFeatures,
        windowInfo,
        client,
        settings,
        no_javascript_access);
}

void CEF_CALLBACK cef_life_span_handler_t_on_after_created(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    //DEBUG_CALLBACK("client->LifeSpanHandler->on_after_created\n");
    go_OnAfterCreated(self, browser->get_identifier(browser), browser);
}

int CEF_CALLBACK cef_life_span_handler_t_run_modal(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    return go_RunModal(self, browser->get_identifier(browser), browser);
}

int CEF_CALLBACK cef_life_span_handler_t_do_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    return go_DoClose(self, browser->get_identifier(browser), browser);
}

void CEF_CALLBACK cef_life_span_handler_t_on_before_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    go_BeforeClose(self, browser->get_identifier(browser), browser);
}

void CEF_CALLBACK cef_load_handler_on_loading_state_change(struct _cef_load_handler_t* self,
    struct _cef_browser_t* browser, int isLoading, int canGoBack,
    int canGoForward) {
    go_OnLoadingStateChange(self, browser->get_identifier(browser), browser, isLoading, canGoBack, canGoForward);
}

void CEF_CALLBACK cef_load_handler_on_load_start(struct _cef_load_handler_t* self,
    struct _cef_browser_t* browser, struct _cef_frame_t* frame) {
    go_OnLoadStart(self, browser->get_identifier(browser), browser, frame);
}

void CEF_CALLBACK cef_load_handler_on_load_end(struct _cef_load_handler_t* self,
    struct _cef_browser_t* browser, struct _cef_frame_t* frame,
    int httpStatusCode) {
    go_OnLoadEnd(self, browser->get_identifier(browser), browser, frame, httpStatusCode);
}

void CEF_CALLBACK cef_load_handler_on_load_error(struct _cef_load_handler_t* self,
    struct _cef_browser_t* browser, struct _cef_frame_t* frame,
    cef_errorcode_t errorCode, const cef_string_t* errorText,
    const cef_string_t* failedUrl) {
    go_OnLoadError(self, browser->get_identifier(browser), browser, frame, errorCode, errorText, failedUrl);
}

int CEF_CALLBACK cef_render_handler_t_get_root_screen_rect(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->get_root_screen_rect");
      return go_RenderHandlerGetRootScreenRect(browser->get_identifier(browser), rect);
}

int CEF_CALLBACK cef_render_handler_t_get_view_rect(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->get_view_rect");
      return go_RenderHandlerGetViewRect(browser->get_identifier(browser), rect);
}

int CEF_CALLBACK cef_render_handler_t_get_screen_point(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, int viewX, int viewY, int* screenX, int* screenY) {
      //DEBUG_CALLBACK("render_handler->get_screen_point");
      return go_RenderHandlerGetScreenPoint(browser->get_identifier(browser), viewX, viewY, screenX, screenY);
}

int CEF_CALLBACK cef_render_handler_t_get_screen_info(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, struct _cef_screen_info_t* info) {
      //DEBUG_CALLBACK("render_handler->get_screen_info");
      return go_RenderHandlerGetScreenInfo(browser->get_identifier(browser), info);
}

void CEF_CALLBACK cef_render_handler_t_on_popup_show(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, int show) {
      //DEBUG_CALLBACK("render_handler->on_popup_show");
      go_RenderHandlerOnPopupShow(browser->get_identifier(browser), show);
}

void CEF_CALLBACK cef_render_handler_t_on_popup_size(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, const cef_rect_t* rect) {
      //DEBUG_CALLBACK("render_handler->on_popup_size");
      go_RenderHandlerOnPopupSize(browser->get_identifier(browser), rect);
}

void CEF_CALLBACK cef_render_handler_t_on_paint(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_paint_element_type_t type,
      size_t dirtyRectsCount, cef_rect_t const* dirtyRects, const void* buffer,
      int width, int height) {
      //DEBUG_CALLBACK("render_handler->on_paint");
      go_RenderHandlerOnPaint(browser->get_identifier(browser), type, dirtyRectsCount, dirtyRects, buffer, width, height);
}

void CEF_CALLBACK cef_render_handler_t_on_cursor_change(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser, cef_cursor_handle_t cursor) {
      //DEBUG_CALLBACK("render_handler->on_cursor_change");
      go_RenderHandlerOnCursorChange(browser->get_identifier(browser), cursor);
}

void CEF_CALLBACK cef_render_handler_t_on_scroll_offset_changed(struct _cef_render_handler_t* self,
      struct _cef_browser_t* browser) {
      DEBUG_CALLBACK("render_handler->on_scroll_offset_changed");
      go_RenderHandlerOnScrollOffsetChanged(browser->get_identifier(browser));
}

int CEF_CALLBACK cef_resource_handler_t_process_request(struct _cef_resource_handler_t* self,
      struct _cef_request_t* request, struct _cef_callback_t* callback) {
      return 0;
}

int CEF_CALLBACK cef_request_handler_t_on_before_browse(struct _cef_request_handler_t* self,
      struct _cef_browser_t* browser, struct _cef_frame_t* frame,
      struct _cef_request_t* request, int is_redirect)
{
    return 0;
}

int CEF_CALLBACK cef_request_handler_t_on_before_resource_load(
  struct _cef_request_handler_t* self, struct _cef_browser_t* browser,
  struct _cef_frame_t* frame, struct _cef_request_t* request)
{
    return go_RequestHandlerOnBeforeResourceLoad(self, browser, frame, request);
}

struct _cef_resource_handler_t* CEF_CALLBACK cef_request_handler_t_get_resource_handler(
  struct _cef_request_handler_t* self, struct _cef_browser_t* browser,
  struct _cef_frame_t* frame, struct _cef_request_t* request)
{
    return NULL;
}

void CEF_CALLBACK cef_request_handler_t_on_resource_redirect(struct _cef_request_handler_t* self,
  struct _cef_browser_t* browser, struct _cef_frame_t* frame,
  const cef_string_t* old_url, cef_string_t* new_url)
{
}

int CEF_CALLBACK cef_request_handler_t_get_auth_credentials(struct _cef_request_handler_t* self,
  struct _cef_browser_t* browser, struct _cef_frame_t* frame, int isProxy,
  const cef_string_t* host, int port, const cef_string_t* realm,
  const cef_string_t* scheme, struct _cef_auth_callback_t* callback)
{
    return 0;
}

int CEF_CALLBACK cef_request_handler_t_on_quota_request(struct _cef_request_handler_t* self,
  struct _cef_browser_t* browser, const cef_string_t* origin_url,
  int64 new_size, struct _cef_quota_callback_t* callback)
{
    return 0;
}

void CEF_CALLBACK cef_request_handler_t_on_protocol_execution(
  struct _cef_request_handler_t* self, struct _cef_browser_t* browser,
  const cef_string_t* url, int* allow_os_execution)
{
}

int CEF_CALLBACK cef_request_handler_t_on_certificate_error(struct _cef_request_handler_t* self,
  cef_errorcode_t cert_error, const cef_string_t* request_url,
  struct _cef_allow_certificate_error_callback_t* callback)
{
    return 0;
}

int CEF_CALLBACK cef_request_handler_t_on_before_plugin_load(struct _cef_request_handler_t* self,
  struct _cef_browser_t* browser, const cef_string_t* url,
  const cef_string_t* policy_url, struct _cef_web_plugin_info_t* info)
{
    return 0;
}

void CEF_CALLBACK cef_request_handler_t_on_plugin_crashed(struct _cef_request_handler_t* self,
  struct _cef_browser_t* browser, const cef_string_t* plugin_path)
{
}

void CEF_CALLBACK cef_request_handler_t_on_render_process_terminated(
  struct _cef_request_handler_t* self, struct _cef_browser_t* browser,
  cef_termination_status_t status)
{
}

int CEF_CALLBACK on_process_message_received(
        struct _cef_client_t* self,
        struct _cef_browser_t* browser, cef_process_id_t source_process,
        struct _cef_process_message_t* message) {
    //DEBUG_CALLBACK("on_process_message_received");
    return go_OnProcessMessageReceived(self, browser->get_identifier(browser), browser, source_process, message);
}

void initialize_display_handler() {
    cef_display_handler_t* displayHandler = (cef_display_handler_t*)calloc(1, sizeof(cef_display_handler_t));
    displayHandler->base.size = sizeof(cef_display_handler_t);
    initialize_cef_base((cef_base_t*) displayHandler);
    // callbacks
    displayHandler->on_address_change = cef_display_handler_t_on_address_change;
    displayHandler->on_title_change = cef_display_handler_t_on_title_change;
    displayHandler->on_tooltip =  cef_display_handler_t_on_tooltip;
    displayHandler->on_status_message = cef_display_handler_t_on_status_message;
    displayHandler->on_console_message = cef_display_handler_t_on_console_message;
    go_client->display_handler = displayHandler;
}

void initialize_life_span_handler() {
    cef_life_span_handler_t* lifeHandler = (cef_life_span_handler_t*)calloc(1, sizeof(cef_life_span_handler_t));
    //DEBUG_CALLBACK("client->initialize_life_span_handler\n");
    lifeHandler->base.size = sizeof(cef_life_span_handler_t);
    initialize_cef_base((cef_base_t*) lifeHandler);
    // callbacks
    lifeHandler->on_before_popup = cef_life_span_handler_t_on_before_popup; //*
    lifeHandler->on_after_created = cef_life_span_handler_t_on_after_created;
    lifeHandler->run_modal = cef_life_span_handler_t_run_modal; //*
    lifeHandler->do_close = cef_life_span_handler_t_do_close;//*
    lifeHandler->on_before_close = cef_life_span_handler_t_on_before_close;//*
    go_client->life_span_handler = lifeHandler;
}

void initialize_load_handler() {
    cef_load_handler_t* loadHandler = (cef_load_handler_t*)calloc(1, sizeof(cef_load_handler_t));
    //DEBUG_CALLBACK("client->initialize_load_handler\n");
    loadHandler->base.size = sizeof(cef_load_handler_t);
    initialize_cef_base((cef_base_t*) loadHandler);
    // callbacks
    loadHandler->on_loading_state_change = cef_load_handler_on_loading_state_change;
    loadHandler->on_load_start = cef_load_handler_on_load_start;
    loadHandler->on_load_end = cef_load_handler_on_load_end;
    loadHandler->on_load_error = cef_load_handler_on_load_error;
    go_client->load_handler = loadHandler;
}

void initialize_render_handler() {
    //DEBUG_CALLBACK("initialize_render_handler");
    cef_render_handler_t* renderHandler = (cef_render_handler_t*)calloc(1, sizeof(cef_render_handler_t));
    renderHandler->base.size = sizeof(cef_render_handler_t);
    initialize_cef_base((cef_base_t*) renderHandler);
    // callbacks
    renderHandler->get_root_screen_rect = cef_render_handler_t_get_root_screen_rect;
    renderHandler->get_view_rect = cef_render_handler_t_get_view_rect;
    renderHandler->get_screen_point = cef_render_handler_t_get_screen_point;
    renderHandler->get_screen_info = cef_render_handler_t_get_screen_info;
    renderHandler->on_popup_show = cef_render_handler_t_on_popup_show;
    renderHandler->on_popup_size = cef_render_handler_t_on_popup_size;
    renderHandler->on_paint = cef_render_handler_t_on_paint;
    renderHandler->on_cursor_change = cef_render_handler_t_on_cursor_change;
    renderHandler->on_scroll_offset_changed = cef_render_handler_t_on_scroll_offset_changed;
    DEBUG_POINTER("render_handler", renderHandler);
    //go_AddRef((void *) renderHandler);
    go_client->render_handler = renderHandler;
}

void initialize_cef_request_handler() {
    cef_request_handler_t* requestHandler = (cef_request_handler_t*)calloc(1, sizeof(cef_request_handler_t));
    requestHandler->base.size = sizeof(cef_request_handler_t);
    initialize_cef_base((cef_base_t*) requestHandler);
    // callbacks
    requestHandler->on_before_browse = cef_request_handler_t_on_before_browse;
    requestHandler->on_before_resource_load = cef_request_handler_t_on_before_resource_load;
    requestHandler->get_resource_handler = cef_request_handler_t_get_resource_handler;
    requestHandler->on_resource_redirect = cef_request_handler_t_on_resource_redirect;
    requestHandler->get_auth_credentials = cef_request_handler_t_get_auth_credentials;
    requestHandler->on_quota_request = cef_request_handler_t_on_quota_request;
    requestHandler->on_protocol_execution = cef_request_handler_t_on_protocol_execution;
    requestHandler->on_certificate_error = cef_request_handler_t_on_certificate_error;
    requestHandler->on_before_plugin_load = cef_request_handler_t_on_before_plugin_load;
    requestHandler->on_plugin_crashed = cef_request_handler_t_on_plugin_crashed;
    requestHandler->on_render_process_terminated = cef_request_handler_t_on_render_process_terminated;
    go_client->request_handler = requestHandler;
}

void initialize_cef_response_handler() {
    //DEBUG_CALLBACK("initialize_render_handler");
    cef_resource_handler_t* resourceHandler = (cef_resource_handler_t*)calloc(1, sizeof(cef_resource_handler_t));
    resourceHandler->base.size = sizeof(cef_resource_handler_t);
    initialize_cef_base((cef_base_t*) resourceHandler);
    // callbacks
    //resourceHandler->get_root_screen_rect = cef_render_handler_t_get_root_screen_rect;
    //resourceHandler->get_view_rect = cef_render_handler_t_get_view_rect;
    //resourceHandler->get_screen_point = cef_render_handler_t_get_screen_point;
    //resourceHandler->get_screen_info = cef_render_handler_t_get_screen_info;
    //resourceHandler->on_popup_show = cef_render_handler_t_on_popup_show;
    //resourceHandler->on_popup_size = cef_render_handler_t_on_popup_size;
    //resourceHandler->on_paint = cef_render_handler_t_on_paint;
    //resourceHandler->on_cursor_change = cef_render_handler_t_on_cursor_change;
    //resourceHandler->on_scroll_offset_changed = cef_render_handler_t_on_scroll_offset_changed;
    DEBUG_POINTER("resource_handler", resourceHandler);
    //go_AddRef((void *) resourceHandler);
    go_client->resource_handler = resourceHandler;
}

struct _cef_display_handler_t* CEF_CALLBACK get_display_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_display_handler");
    go_AddRef((void *) go_client->display_handler);
    return go_client->display_handler;
}

struct _cef_life_span_handler_t* CEF_CALLBACK get_life_span_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_life_span_handler");
    go_AddRef((void *) go_client->life_span_handler);
    return go_client->life_span_handler;
}

struct _cef_load_handler_t* CEF_CALLBACK get_load_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_load_handler");
    go_AddRef((void *) go_client->load_handler);
    return go_client->load_handler;
}

struct _cef_render_handler_t* CEF_CALLBACK get_render_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_render_handler");
    go_AddRef((void *) go_client->render_handler);
    return go_client->render_handler;
}

struct _cef_request_handler_t* CEF_CALLBACK get_request_handler(
        struct _cef_client_t* self) {
    //DEBUG_CALLBACK("get_request_handler");
    go_AddRef((void *) go_client->request_handler);
    return go_client->request_handler;
}

void initialize_client_handler(struct _cef_client_t* client) {
    DEBUG_POINTER("initialize_client_handler", client);
    go_client = (cef_go_client*)calloc(1, sizeof(cef_go_client));
    initialize_display_handler();
    initialize_life_span_handler();
    initialize_load_handler();
    //initialize_render_handler();
    initialize_cef_request_handler();

    client->base.size = sizeof(cef_client_t);
    initialize_cef_base((cef_base_t*)client);
    // callbacks
    //DEBUG_CALLBACK("set_display_handler");
    client->get_display_handler = get_display_handler;
    client->get_life_span_handler = get_life_span_handler;
    client->get_load_handler = get_load_handler;
    //client->get_render_handler = get_render_handler;
    client->get_request_handler = get_request_handler;
    client->on_process_message_received = on_process_message_received;
}
