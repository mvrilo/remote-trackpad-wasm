//+build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

var ws js.Value = js.Global().Get("WebSocket").New("wss://10.0.0.3:4430/ws")
var div js.Value = js.Global().Get("document").Call("getElementById", "main")

func fireEvent(eventName string, args []js.Value) {
	event := args[0]
	event.Call("preventDefault")

	touch := event.Get("changedTouches")
	if touch.Length() == 0 {
		return
	}

	point := touch.Index(0)
	x := point.Get("pageX").Int()
	y := point.Get("pageY").Int()

	ws.Call("send", fmt.Sprintf(`{ "type": "%s", "data": { "x": %d, "y": %d } }`, eventName, x, y))
}

func main() {
	done := make(chan struct{})

	style := div.Get("style")
	style.Set("position", "absolute")
	style.Set("backgroundColor", "#fefefe")
	style.Set("top", "0")
	style.Set("left", "0")
	style.Set("width", "100%")
	style.Set("height", "100%")

	var isMove bool

	onTouchStart := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		isMove = false
		return nil
	})

	onTouchMove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		isMove = true
		fireEvent("move", args)

		return nil
	})

	onTouchEnd := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !isMove {
			fireEvent("tap", args)
		}

		return nil
	})

	div.Set("ontouchstart", onTouchStart)
	div.Set("ontouchmove", onTouchMove)
	div.Set("ontouchend", onTouchEnd)

	<-done
}
