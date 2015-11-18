// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package cef

/*
#include <stdlib.h>
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
*/
import "C"

import (
	"fmt"
)

type LifeSpanHandler struct {
	browser *Browser
}

func (l *LifeSpanHandler) RegisterAndWaitForBrowser() (browser *Browser, err error) {
	browserWaitGroup.Wait()
	browser = l.browser
	return
}

func (l *LifeSpanHandler) RunModal(browser *Browser) int {
	url := browser.GetURL()
	logger.Printf("RunModal, %v, url %s\n", browser, url)
	return 0
}

func (l *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	url := browser.GetURL()
	fmt.Printf("LifeSpanHandler::OnAfterCreated\n")
	logger.Printf("Browser: %v Url %s\n", browser, url)

	hwnd := browser.GetRootWindowHandle()
	fmt.Printf("RootWindowHandle=%v\n", hwnd)
	l.browser = browser

	browsers[browser.Id] = browser
	browserHandles[hwnd] = browser

	if OnBrowserCreated != nil {
		OnBrowserCreated(hwnd, browser)
	}
}

func (l *LifeSpanHandler) DoClose(browser *Browser) int {
	url := browser.GetURL()
	logger.Printf("DoClose, %v, url %s\n", browser, url)
	return 0
}

func (l *LifeSpanHandler) BeforeClose(browser *Browser) {
	url := browser.GetURL()
	logger.Printf("BeforeClose, %v, url %s\n", browser, url)
}

var _LifeSpanHandler *C.struct__cef_life_span_handler_t // requires reference counting
var globalLifespanHandler *LifeSpanHandler
