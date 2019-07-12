package main

import (
	// #cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
	// #include <CoreGraphics/CoreGraphics.h>
	// static void releaseCGEvent(CGEventRef o) {
	// 	CFRelease(o);
	// }
	"C"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// routes:
// /
// /ws
// /wasm_exec.js
// /main.wasm

var upgrader = websocket.Upgrader{}

type move struct {
	X int
	Y int
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var data = move{}
		if err := json.Unmarshal(message, &data); err != nil {
			log.Println(err)
			break
		}

		println(data.X, data.Y)
		moveMouse(&data)

		// log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, []byte("ok"))
		if err != nil {
			// log.Println("write:", err)
			break
		}
	}
}

func moveMouse(pos *move) {
	point := C.CGPointMake(C.CGFloat(pos.X), C.CGFloat(pos.Y))

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
	addr := flag.String("addr", ":8080", "http server address")
	cert := flag.String("cert", "", "https cert")
	key := flag.String("key", "", "https key")
	flag.Parse()

	http.HandleFunc("/ws", ws)
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	println("Server started at", *addr)
	log.Fatal(http.ListenAndServeTLS(*addr, *cert, *key, nil))
}
