//go:build !secretcli
// +build !secretcli

package api

// #include <stdlib.h>
// #include "bindings.h"
import "C"

import (
	"fmt"
	"syscall"
)

// nice aliases to the rust names
type i32 = C.int32_t

type (
	i64    = C.int64_t
	u64    = C.uint64_t
	u32    = C.uint32_t
	u8     = C.uint8_t
	u8_ptr = *C.uint8_t
	usize  = C.uintptr_t
	cint   = C.int
)

func GetRandom() ([]byte, error) {
	errmsg := C.Buffer{}

	res, err := C.get_health_check(&errmsg)
	if err != nil {
		return nil, errorWithMessage(err, errmsg)
	}
	return receiveVector(res), nil
}

/**** To error module ***/

func errorWithMessage(err error, b C.Buffer) error {
	// this checks for out of gas as a special case
	if errno, ok := err.(syscall.Errno); ok && int(errno) == 2 {
		panic("Wtf please go away")
	}
	msg := receiveVector(b)
	if msg == nil {
		return err
	}
	return fmt.Errorf("%s", string(msg))
}
