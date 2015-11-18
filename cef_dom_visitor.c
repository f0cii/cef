#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_v8_capi.h"
#include "cef_app.h"
#include "cef_client.h"
#include "cef_helpers.h"

void CEF_CALLBACK cef_dom_visit(struct _cef_domvisitor_t* self,
    struct _cef_domdocument_t* document)
{
    go_Log("cef_dom_visit");
    go_DOMVisitorVisit(self, document);
}

void initialize_cef_dom_visitor(struct _cef_domvisitor_t* visitor) {
    //cef_domvisitor_t* visitor = (cef_domvisitor_t*)calloc(1, sizeof(cef_domvisitor_t));
    visitor->base.size = sizeof(cef_domvisitor_t);
    initialize_cef_base((cef_base_t*) visitor);
    // callbacks
    visitor->visit = cef_dom_visit;
}

void VisitDOM(cef_browser_t* browser, cef_domvisitor_t* visitor)
{
	cef_frame_t * frame = browser->get_main_frame(browser);
	frame->visit_dom(frame, visitor);
}