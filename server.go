package main

import (
	"bytes"
	"encoding/base64"
	"github.com/vova616/screenshot"
	_ "image"
	"image/jpeg"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println(err)
			continue
		}
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, img, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
		_, err = conn.Write([]byte(encoded))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Millisecond)
	}
}
