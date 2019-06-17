package main

import (
	// #cgo LDFLAGS: -framework CoreGraphics
	// #include <CoreGraphics/CoreGraphics.h>
	"C"
	"fmt"
	"unsafe"
)

func main() {
	display := C.CGMainDisplayID
	x := (C.CGFloat)(1)
	y := (C.CGFloat)(2)
	point := C.CGPointMake(x, y)

	// displayid := unsafe.Int(display)
	d := *(*C.uint)(unsafe.Pointer(&display))
	p := *(*C.CGPoint)(unsafe.Pointer(&point))

	res := C.CGDisplayMoveCursorToPoint(d, p)
	fmt.Println(display, point, res)
}
