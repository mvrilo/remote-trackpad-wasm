//+build js,wasm
package main

import (
	"fmt"
	"syscall/js"
)

func domListener() {
	doc := js.Global().Get("document")
	div := doc.Call("getElementById", "main")

	style := div.Get("style")
	style.Set("position", "absolute")
	style.Set("backgroundColor", "#fefefe")
	style.Set("top", "0")
	style.Set("left", "0")
	style.Set("width", "100%")
	style.Set("height", "100%")

	var onTouchEnd js.Func
	onTouchEnd = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		touches := event.Get("changedTouches")
		js.Global().Call("alert", touches.Length())
		if touches.Length() > 0 {
			// message := touches.Index(0).Get("pageX")
			// js.Global().Call("alert", message)
		}
		fmt.Printf("#%v\n", event)
		return nil
	})

	var onTouchMove js.Func
	onTouchMove = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		event.Call("preventDefault")

		touches := event.Get("touches")
		if touches.Length() > 0 {
			fmt.Printf("%#v\n", touches.Index(0).Get("pageX"))
		}
		fmt.Printf("#%v\n", event)

		return nil
	})

	div.Set("ontouchmove", onTouchMove)
	div.Set("ontouchend", onTouchEnd)
}

func wsListener() {
	websocket := js.Global().Get("WebSocket")
	ws := websocket.New("ws://localhost:8080/ws")

	var onOpen js.Func
	onOpen = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("open!")
		// onOpen.Release()
		return nil
	})

	var onClose js.Func
	onClose = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// onClose.Release()
		return nil
	})

	ws.Set("onopen", onOpen)
	// ws.Set("onmessage", onMessage)
	ws.Set("onclose", onClose)
	// ws.Set("onerror", onError)
}

func main() {
	println("start")
	done := make(chan struct{})

	domListener()
	wsListener()

	<-done
}
