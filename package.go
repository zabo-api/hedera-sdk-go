package hedera

// #cgo CFLAGS: -I ./include
// #cgo LDFLAGS: -l hedera
// #cgo darwin LDFLAGS: -L libs/x86_64-apple-darwin -framework Security -l c++
// #cgo linux LDFLAGS: -L libs/x86_64-unknown-linux-musl
// #cgo windows LDFLAGS: -L libs/x86_64-pc-windows-gnu -l ws2_32 -l iphlpapi -l dbghelp -l userenv
// #include <stdlib.h>
// #include "hedera.h"
import "C"
import "unsafe"
import "errors"

func hederaLastError() error {
	return errors.New(hederaString(C.hedera_last_error()))
}

func hederaString(bytes *C.char) string {
	defer C.free(unsafe.Pointer(bytes))
	return C.GoString(bytes)
}
