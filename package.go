package hedera

// #cgo CFLAGS: -I ./include
// #cgo LDFLAGS: -l hedera
// #cgo darwin LDFLAGS: -L libs/x86_64-apple-darwin -framework Security -l c++
// #include <stdlib.h>
// #include "hedera-error.h"
import "C"
import (
	"errors"
	"unsafe"
)

func hederaError(err C.longlong) error {
	messageBytes := C.hedera_error_message(err)
	defer C.free(unsafe.Pointer(messageBytes))

	return errors.New(C.GoString(messageBytes))
}
