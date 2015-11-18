#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

void CEF_CALLBACK string_visitor_visit(struct _cef_string_visitor_t* self,//struct
    const cef_string_t* string) {

  //DCHECK(self);
  if (!self)
    return;
  // Unverified params: string

  // Execute
  //CefStringVisitorCppToC::Get(self)->Visit(
  //    CefString(string));
  go_StringVisitorVisit(self, string);
}

void initialize_cef_string_visitor(struct _cef_string_visitor_t* visitor) {
    //cef_string_visitor_t* visitor = (cef_string_visitor_t*)calloc(1, sizeof(cef_string_visitor_t));
    visitor->base.size = sizeof(cef_string_visitor_t);
    initialize_cef_base((cef_base_t*) visitor);
    // callbacks
    visitor->visit = string_visitor_visit;
}
