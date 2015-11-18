#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

int cef_process_message_is_valid(struct _cef_process_message_t* self)
{
    return self->is_valid(self);
}

int cef_process_message_is_read_only(struct _cef_process_message_t* self)
{
    return self->is_read_only(self);
}

struct _cef_process_message_t* cef_process_message_copy(
      struct _cef_process_message_t* self)
{
    return self->copy(self);
}

cef_string_userfree_t cef_process_message_get_name(
      struct _cef_process_message_t* self)
{
    return self->get_name(self);
}

struct _cef_list_value_t* cef_process_message_get_argument_list(
      struct _cef_process_message_t* self)
{
    return self->get_argument_list(self);
}