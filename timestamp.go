package hedera

// #include "hedera-timestamp.h"
import "C"
import "unsafe"

type Timestamp struct {
	Seconds int64
	Nanos   int32
	// Added by `go tool cgo` to make this struct C compatible
	padding [4]byte
}

func Now() Timestamp {
	response := C.hedera_timestamp_now()
	return *((*Timestamp)(unsafe.Pointer(&response)))
}
