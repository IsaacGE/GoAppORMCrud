package controllers

import (
	"fmt"
	"net"
)

type Response struct {
	Connect net.Conn
}

func (r Response) SendData(statusCode int, content string) {
	fmt.Fprintf(r.Connect, "HTTP/1.1 %d\r\n", statusCode)
	fmt.Fprintf(r.Connect, "Content-Type: text/html\r\n")
	fmt.Fprint(r.Connect, "Access-Control-Allow-Origin: *\r\n")
	fmt.Fprint(r.Connect, "Connection: Keep-Alive\r\n")
	fmt.Fprint(r.Connect, "Keep-Alive: timeout=5, max=997\r\n")
	fmt.Fprintf(r.Connect, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(r.Connect, "\r\n")
	fmt.Fprint(r.Connect, content)
}

func (r Response) SendJSON(statusCode int, content string) {
	fmt.Fprintf(r.Connect, "HTTP/1.1 %d\r\n", statusCode)
	fmt.Fprintf(r.Connect, "Content-Type: application/json\r\n")
	fmt.Fprint(r.Connect, "Access-Control-Allow-Origin: *\r\n")
	fmt.Fprint(r.Connect, "Connection: Keep-Alive\r\n")
	fmt.Fprint(r.Connect, "Keep-Alive: timeout=5, max=997\r\n")
	fmt.Fprintf(r.Connect, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(r.Connect, "\r\n")
	fmt.Fprint(r.Connect, content)
}
