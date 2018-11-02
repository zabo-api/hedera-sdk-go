package hedera

// #cgo CFLAGS: -I ./include
// #cgo LDFLAGS: -l hedera
// #cgo darwin LDFLAGS: -L libs/x86_64-apple-darwin -framework Security -l c++
// #cgo linux LDFLAGS: -L libs/x86_64-unknown-linux-gnu -l stdc++ -l m -l dl
// #include <stdlib.h>
// #include "hedera-error.h"
import "C"
import (
	"errors"
	"unsafe"
)

func hederaError(err C.int64_t) error {
	if err == 0 {
		return nil
	}

	messageBytes := C.hedera_error_message(err)
	defer C.free(unsafe.Pointer(messageBytes))

	return errors.New(C.GoString(messageBytes))
}
