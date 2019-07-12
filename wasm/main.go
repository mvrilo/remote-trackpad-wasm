//+build js,wasm
package main

import (
	"fmt"
	"syscall/js"
)

func domListener(ws js.Value) {
	doc := js.Global().Get("document")
	div := doc.Call("getElementById", "main")

	style := div.Get("style")
	style.Set("position", "absolute")
	style.Set("backgroundColor", "#fefefe")
	style.Set("top", "0")
	style.Set("left", "0")
	style.Set("width", "100%")
	style.Set("height", "100%")

	onTouchMove := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")

		touch := event.Get("changedTouches")
		if touch.Length() > 0 {
			point := touch.Index(0)
			x := point.Get("pageX").Int()
			y := point.Get("pageY").Int()
			data := fmt.Sprintf(`{ "x": %d, "y": %d }`, x, y)
			ws.Call("send", data)
		}
		return nil
	})

	div.Set("ontouchmove", onTouchMove)
}

func wsListener(ws js.Value) {
	var onOpen js.Func
	onOpen = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("open!")
		// onOpen.Release()
		return nil
	})
	ws.Set("onopen", onOpen)
}

func main() {
	println("start")
	done := make(chan struct{})

	websocket := js.Global().Get("WebSocket")
	ws := websocket.New("wss://cthulhu.local:8080/ws")

	domListener(ws)
	wsListener(ws)

	<-done
}
