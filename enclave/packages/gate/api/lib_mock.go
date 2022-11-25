//go:build secretcli
// +build secretcli

package api

//
//// #include <stdlib.h>
//// #include "bindings.h"
//import "C"

// nice aliases to the rust names
//type i32 = C.int32_t
//type i64 = C.int64_t
//type u64 = C.uint64_t
//type u32 = C.uint32_t
//type u8 = C.uint8_t
//type u8_ptr = *C.uint8_t
//type usize = C.uintptr_t
//type cint = C.int
//
//type Cache struct {
//	ptr *C.cache_t
//}

type Cache struct{}

func GetRandom() ([]byte, error) {
	return nil, nil
}
