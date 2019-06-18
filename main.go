package main

import (
	// #cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
	// #include <CoreGraphics/CoreGraphics.h>
	// static void releaseCGEvent(CGEventRef o) {
	// 	CFRelease(o);
	// }
	"C"
	"flag"
)

func moveMouse(x, y int) {
	point := C.CGPointMake(C.CGFloat(x), C.CGFloat(y))

	move := C.CGEventCreateMouseEvent(
		0,
		C.kCGEventMouseMoved,
		point,
		C.kCGMouseButtonLeft,
	)

	defer C.releaseCGEvent(move)
	C.CGEventPost(C.kCGHIDEventTap, move)
}

func main() {
	x := flag.Int("x", 0, "x position")
	y := flag.Int("y", 0, "y position")
	flag.Parse()

	if *x == 0 || *y == 0 {
		flag.Usage()
		return
	}

	moveMouse(*x, *y)
}
