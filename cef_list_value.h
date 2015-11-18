#include "include/internal/cef_string.h"
#include "include/capi/cef_v8_capi.h"

extern int CEF_CALLBACK cef_list_value_t_is_valid(struct _cef_list_value_t* self);

extern int CEF_CALLBACK cef_list_value_t_is_owned(struct _cef_list_value_t* self);

extern int CEF_CALLBACK cef_list_value_t_is_read_only(struct _cef_list_value_t* self);

extern struct _cef_list_value_t* CEF_CALLBACK cef_list_value_t_copy(
  struct _cef_list_value_t* self);

extern int CEF_CALLBACK cef_list_value_t_set_size(struct _cef_list_value_t* self, size_t size);

extern size_t CEF_CALLBACK cef_list_value_t_get_size(struct _cef_list_value_t* self);

extern int CEF_CALLBACK cef_list_value_t_clear(struct _cef_list_value_t* self);

extern int CEF_CALLBACK cef_list_value_t_remove(struct _cef_list_value_t* self, int index);

extern cef_value_type_t CEF_CALLBACK cef_list_value_t_get_type(struct _cef_list_value_t* self,
  int index);

extern int CEF_CALLBACK cef_list_value_t_get_bool(struct _cef_list_value_t* self, int index);

extern int CEF_CALLBACK cef_list_value_t_get_int(struct _cef_list_value_t* self, int index);

extern double CEF_CALLBACK cef_list_value_t_get_double(struct _cef_list_value_t* self, int index);

extern cef_string_userfree_t CEF_CALLBACK cef_list_value_t_get_string(
  struct _cef_list_value_t* self, int index);

extern struct _cef_binary_value_t* CEF_CALLBACK cef_list_value_t_get_binary(
  struct _cef_list_value_t* self, int index);

extern struct _cef_dictionary_value_t* CEF_CALLBACK cef_list_value_t_get_dictionary(
  struct _cef_list_value_t* self, int index);

extern struct _cef_list_value_t* CEF_CALLBACK cef_list_value_t_get_list(
  struct _cef_list_value_t* self, int index);

extern int CEF_CALLBACK cef_list_value_t_set_null(struct _cef_list_value_t* self, int index);

extern int CEF_CALLBACK cef_list_value_t_set_bool(struct _cef_list_value_t* self, int index,
  int value);

extern int CEF_CALLBACK cef_list_value_t_set_int(struct _cef_list_value_t* self, int index,
  int value);

extern int CEF_CALLBACK cef_list_value_t_set_double(struct _cef_list_value_t* self, int index,
  double value);

extern int CEF_CALLBACK cef_list_value_t_set_string(struct _cef_list_value_t* self, int index,
  const cef_string_t* value);

extern int CEF_CALLBACK cef_list_value_t_set_binary(struct _cef_list_value_t* self, int index,
  struct _cef_binary_value_t* value);

extern int CEF_CALLBACK cef_list_value_t_set_dictionary(struct _cef_list_value_t* self, int index,
  struct _cef_dictionary_value_t* value);

extern int CEF_CALLBACK cef_list_value_t_set_list(struct _cef_list_value_t* self, int index,
  struct _cef_list_value_t* value);