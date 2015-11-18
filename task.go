package cef

/*
#include <stdlib.h>
#include "include/capi/cef_task_capi.h"
extern void initialize_task(struct _cef_task_t* self);
extern void helper_cef_post_task(cef_thread_id_t threadId, struct _cef_task_t* what);
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

type ThreadId int

const (
	TID_UI ThreadId = iota
	TID_DB
	TID_FILE
	TID_FILE_USER_BLOCKING
	TID_PROCESS_LAUNCHER
	TID_CACHE
	TID_IO
	TID_RENDERER
)

var (
	taskMap      = make(map[unsafe.Pointer]TaskToExecute)
	postTaskLock sync.Mutex
)

type TaskToExecute func()

//export go_TaskExecute
func go_TaskExecute(self *C.struct__cef_task_t) {
	fmt.Printf("go_TaskExecute\n")
	postTaskLock.Lock()
	toExecute, ok := taskMap[unsafe.Pointer(self)]
	if ok {
		delete(taskMap, unsafe.Pointer(self))
	}
	postTaskLock.Unlock()
	if ok {
		toExecute()
	}
}

// allows you to execute a task on the specified thread
func PostTask(thread ThreadId, t TaskToExecute) {
	taskT := (*C.struct__cef_task_t)(
		C.calloc(1, C.sizeof_struct__cef_task_t))
	C.initialize_task(taskT)
	go_AddRef(unsafe.Pointer(taskT))

	// not defering the unlock, since i don't know if cef might immedialty execute it,
	// and thus we deadlock ourselves.
	postTaskLock.Lock()
	taskMap[unsafe.Pointer(taskT)] = t
	postTaskLock.Unlock()

	C.helper_cef_post_task(C.cef_thread_id_t(thread), taskT)
}
