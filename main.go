//+build darwin

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

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type event struct {
	Type string
	Data *position
}

type position struct {
	X int
	Y int
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		defer conn.Close()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Println(err)
				break
			}

			println("-> receive data:", string(msg))

			var evt = event{}
			if err := json.Unmarshal(msg, &evt); err != nil {
				log.Println(err)
				break
			}

			if evt.Type == "tap" {
				tap(evt.Data)
			} else {
				move(evt.Data)
			}
		}
	}()
}

func makeEvent(mouseEvent C.CGEventType, pos *position) {
	point := C.CGPointMake(C.CGFloat(pos.X), C.CGFloat(pos.Y))

	evt := C.CGEventCreateMouseEvent(
		0,
		mouseEvent,
		point,
		C.kCGMouseButtonLeft,
	)

	defer C.releaseCGEvent(evt)
	C.CGEventPost(C.kCGHIDEventTap, evt)
}

func tap(pos *position) {
	makeEvent(C.kCGEventLeftMouseDown, pos)
}

func move(pos *position) {
	makeEvent(C.kCGEventMouseMoved, pos)
}

func main() {
	addr := flag.String("addr", ":8080", "http server address")
	cert := flag.String("cert", "", "https cert")
	key := flag.String("key", "", "https key")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/ws", wsHandler)
	router.Handle("/", http.FileServer(http.Dir("./assets")))

	server := &http.Server{
		Handler: router,
		Addr:    *addr,
	}

	println("Server started at", *addr)
	log.Fatal(server.ListenAndServeTLS(*cert, *key))
}
