#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

extern cef_browser_t* cef_browser_host_t_get_browser(cef_browser_host_t *self);

extern void cef_browser_host_t_close_browser(cef_browser_host_t *self, int force_close);

extern void cef_browser_host_t_set_focus(cef_browser_host_t *self, int focus);

extern void cef_browser_host_t_set_window_visibility(struct _cef_browser_host_t* self,
    int visible);

extern cef_window_handle_t cef_browser_host_t_get_window_handle(
    struct _cef_browser_host_t* self);

extern cef_window_handle_t cef_browser_host_t_get_opener_window_handle(
    struct _cef_browser_host_t* self);

extern struct _cef_client_t* cef_browser_host_t_get_client(
  struct _cef_browser_host_t* self);

extern struct _cef_request_context_t* cef_browser_host_t_get_request_context(
  struct _cef_browser_host_t* self);

extern double cef_browser_host_t_get_zoom_level(struct _cef_browser_host_t* self);

extern void cef_browser_host_t_set_zoom_level(struct _cef_browser_host_t* self,
  double zoomLevel);

extern void cef_browser_host_t_run_file_dialog(struct _cef_browser_host_t* self,
  cef_file_dialog_mode_t mode, const cef_string_t* title,
  const cef_string_t* default_file_name, cef_string_list_t accept_types,
  struct _cef_run_file_dialog_callback_t* callback);

extern void cef_browser_host_t_start_download(struct _cef_browser_host_t* self,
  const cef_string_t* url);

extern void cef_browser_host_t_print(struct _cef_browser_host_t* self);

extern void cef_browser_host_t_find(struct _cef_browser_host_t* self, int identifier,
  const cef_string_t* searchText, int forward, int matchCase,
  int findNext);

extern void cef_browser_host_t_stop_finding(struct _cef_browser_host_t* self,
  int clearSelection);

extern void cef_browser_host_t_show_dev_tools(struct _cef_browser_host_t* self,
  const struct _cef_window_info_t* windowInfo,
  struct _cef_client_t* client,
  const struct _cef_browser_settings_t* settings,
  const cef_point_t* inspect_element_at);

extern void cef_browser_host_t_close_dev_tools(struct _cef_browser_host_t* self);

extern void cef_browser_host_t_set_mouse_cursor_change_disabled(
  struct _cef_browser_host_t* self, int disabled);

extern int cef_browser_host_t_is_mouse_cursor_change_disabled(
  struct _cef_browser_host_t* self);

extern void cef_browser_host_t_replace_misspelling(struct _cef_browser_host_t* self,
  const cef_string_t* word);

extern void cef_browser_host_t_add_word_to_dictionary(struct _cef_browser_host_t* self,
  const cef_string_t* word);

extern int cef_browser_host_t_is_window_rendering_disabled(
  struct _cef_browser_host_t* self);

extern void cef_browser_host_t_was_resized(struct _cef_browser_host_t* self);

extern void cef_browser_host_t_was_hidden(struct _cef_browser_host_t* self, int hidden);

extern void cef_browser_host_t_notify_screen_info_changed(
  struct _cef_browser_host_t* self);

extern void cef_browser_host_t_invalidate(struct _cef_browser_host_t* self,
  cef_paint_element_type_t type);

extern void cef_browser_host_t_send_key_event(struct _cef_browser_host_t* self,
  const struct _cef_key_event_t* event);

extern void cef_browser_host_t_send_mouse_click_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event, cef_mouse_button_type_t type,
  int mouseUp, int clickCount);

extern void cef_browser_host_t_send_mouse_move_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event, int mouseLeave);

extern void cef_browser_host_t_send_mouse_wheel_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event, int deltaX, int deltaY);

extern void cef_browser_host_t_send_focus_event(struct _cef_browser_host_t* self,
  int setFocus);

extern void cef_browser_host_t_send_capture_lost_event(
  struct _cef_browser_host_t* self);

extern cef_text_input_context_t cef_browser_host_t_get_nstext_input_context(
  struct _cef_browser_host_t* self);

extern void cef_browser_host_t_handle_key_event_before_text_input_client(
  struct _cef_browser_host_t* self, cef_event_handle_t keyEvent);

extern void cef_browser_host_t_handle_key_event_after_text_input_client(
  struct _cef_browser_host_t* self, cef_event_handle_t keyEvent);

extern void cef_browser_host_t_drag_target_drag_enter(struct _cef_browser_host_t* self,
  struct _cef_drag_data_t* drag_data,
  const struct _cef_mouse_event_t* event,
  cef_drag_operations_mask_t allowed_ops);

extern void cef_browser_host_t_drag_target_drag_over(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event,
  cef_drag_operations_mask_t allowed_ops);

extern void cef_browser_host_t_drag_target_drag_leave(struct _cef_browser_host_t* self);

extern void cef_browser_host_t_drag_target_drop(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event);

extern void cef_browser_host_t_drag_source_ended_at(struct _cef_browser_host_t* self,
  int x, int y, cef_drag_operations_mask_t op);

extern void cef_browser_host_t_drag_source_system_drag_ended(
  struct _cef_browser_host_t* self);