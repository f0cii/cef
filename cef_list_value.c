#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

///
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
///
int CEF_CALLBACK cef_list_value_t_is_valid(struct _cef_list_value_t* self)
{
    return self->is_valid(self);
}

///
// Returns true (1) if this object is currently owned by another object.
///
int CEF_CALLBACK cef_list_value_t_is_owned(struct _cef_list_value_t* self)
{
    return self->is_owned(self);
}

///
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
///
int CEF_CALLBACK cef_list_value_t_is_read_only(struct _cef_list_value_t* self)
{
    return self->is_read_only(self);
}

///
// Returns a writable copy of this object.
///
struct _cef_list_value_t* CEF_CALLBACK cef_list_value_t_copy(
  struct _cef_list_value_t* self)
{
    return self->copy(self);
}

///
// Sets the number of values. If the number of values is expanded all new
// value slots will default to type null. Returns true (1) on success.
///
int CEF_CALLBACK cef_list_value_t_set_size(struct _cef_list_value_t* self, size_t size)
{
    return self->set_size(self, size);
}

///
// Returns the number of values.
///
size_t CEF_CALLBACK cef_list_value_t_get_size(struct _cef_list_value_t* self)
{
    return self->get_size(self);
}

///
// Removes all values. Returns true (1) on success.
///
int CEF_CALLBACK cef_list_value_t_clear(struct _cef_list_value_t* self)
{
    return self->clear(self);
}

///
// Removes the value at the specified index.
///
int CEF_CALLBACK cef_list_value_t_remove(struct _cef_list_value_t* self, int index)
{
    return self->remove(self, index);
}

///
// Returns the value type at the specified index.
///
cef_value_type_t CEF_CALLBACK cef_list_value_t_get_type(struct _cef_list_value_t* self,
  int index)
{
    return self->get_type(self, index);
}

///
// Returns the value at the specified index as type bool.
///
int CEF_CALLBACK cef_list_value_t_get_bool(struct _cef_list_value_t* self, int index)
{
    return self->get_bool(self, index);
}

///
// Returns the value at the specified index as type int.
///
int CEF_CALLBACK cef_list_value_t_get_int(struct _cef_list_value_t* self, int index)
{
    return self->get_int(self, index);
}

///
// Returns the value at the specified index as type double.
///
double CEF_CALLBACK cef_list_value_t_get_double(struct _cef_list_value_t* self, int index)
{
    return self->get_double(self, index);
}

///
// Returns the value at the specified index as type string.
///
// The resulting string must be freed by calling cef_string_userfree_free().
cef_string_userfree_t CEF_CALLBACK cef_list_value_t_get_string(
  struct _cef_list_value_t* self, int index)
{
    return self->get_string(self, index);
}

///
// Returns the value at the specified index as type binary.
///
struct _cef_binary_value_t* CEF_CALLBACK cef_list_value_t_get_binary(
  struct _cef_list_value_t* self, int index)
{
    return self->get_binary(self, index);
}

///
// Returns the value at the specified index as type dictionary.
///
struct _cef_dictionary_value_t* CEF_CALLBACK cef_list_value_t_get_dictionary(
  struct _cef_list_value_t* self, int index)
{
    return self->get_dictionary(self, index);
}

///
// Returns the value at the specified index as type list.
///
struct _cef_list_value_t* CEF_CALLBACK cef_list_value_t_get_list(
  struct _cef_list_value_t* self, int index)
{
    return self->get_list(self, index);
}

///
// Sets the value at the specified index as type null. Returns true (1) if the
// value was set successfully.
///
int CEF_CALLBACK cef_list_value_t_set_null(struct _cef_list_value_t* self, int index)
{
    return self->set_null(self, index);
}

///
// Sets the value at the specified index as type bool. Returns true (1) if the
// value was set successfully.
///
int CEF_CALLBACK cef_list_value_t_set_bool(struct _cef_list_value_t* self, int index,
  int value)
{
    return self->set_bool(self, index, value);
}

///
// Sets the value at the specified index as type int. Returns true (1) if the
// value was set successfully.
///
int CEF_CALLBACK cef_list_value_t_set_int(struct _cef_list_value_t* self, int index,
  int value)
{
    return self->set_int(self, index, value);
}

///
// Sets the value at the specified index as type double. Returns true (1) if
// the value was set successfully.
///
int CEF_CALLBACK cef_list_value_t_set_double(struct _cef_list_value_t* self, int index,
  double value)
{
    return self->set_double(self, index, value);
}

///
// Sets the value at the specified index as type string. Returns true (1) if
// the value was set successfully.
///
int CEF_CALLBACK cef_list_value_t_set_string(struct _cef_list_value_t* self, int index,
  const cef_string_t* value)
{
    return self->set_string(self, index, value);
}

///
// Sets the value at the specified index as type binary. Returns true (1) if
// the value was set successfully. After calling this function the |value|
// object will no longer be valid. If |value| is currently owned by another
// object then the value will be copied and the |value| reference will not
// change. Otherwise, ownership will be transferred to this object and the
// |value| reference will be invalidated.
///
int CEF_CALLBACK cef_list_value_t_set_binary(struct _cef_list_value_t* self, int index,
  struct _cef_binary_value_t* value)
{
    return self->set_binary(self, index, value);
}

///
// Sets the value at the specified index as type dict. Returns true (1) if the
// value was set successfully. After calling this function the |value| object
// will no longer be valid. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
///
int CEF_CALLBACK cef_list_value_t_set_dictionary(struct _cef_list_value_t* self, int index,
  struct _cef_dictionary_value_t* value)
{
    return self->set_dictionary(self, index, value);
}

///
// Sets the value at the specified index as type list. Returns true (1) if the
// value was set successfully. After calling this function the |value| object
// will no longer be valid. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
///
int CEF_CALLBACK cef_list_value_t_set_list(struct _cef_list_value_t* self, int index,
  struct _cef_list_value_t* value)
{
    return self->set_list(self, index, value);
}