package router

import (
	response "GoCrudORM/controllers/response"
	templateHandler "GoCrudORM/helpers"
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Route   string
	Headers map[string]string
	Hosts   string
	Body    []string
}

type Router struct {
	routes map[string]func(*response.Response, *Request)
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func(*response.Response, *Request)),
	}
}

func (r *Router) AddRoute(route string, handler func(*response.Response, *Request)) {
	r.routes[route] = handler
}

/**
 * Handle the request
 * @param conn
 * @param router
 */
func HandleRequest(conn net.Conn, router *Router) {
	scanner := bufio.NewScanner(conn)
	var header, body []string
	f := true
	contentLength := 0

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			f = false
			if strings.HasPrefix(header[0], "POST") {
				for _, h := range header {
					parts := strings.SplitN(h, ": ", 2)
					if len(parts) == 2 && parts[0] == "Content-Length" {
						contentLengthStr := parts[1]
						contentLength, _ = strconv.Atoi(contentLengthStr)
						break
					}
				}
			}
		}
		if f {
			header = append(header, ln)
		}
		if !f && ln == "" {
			break
		}
	}

	req := parseRequest(header)
	response := response.Response{Connect: conn}
	handler, found := router.routes[req.Route]
	if !found {
		handler = NotFoundHandler
	}

	if req.Method == "POST" {
		count := 1
		bodyString := ""
		closeBodyLabel := ""
		for scanner.Scan() && contentLength > 0 {
			ln := scanner.Text()
			if count == 1 {
				if strings.Contains(string(ln[0]), "{") {
					closeBodyLabel = "}"
				} else if strings.Contains(string(ln[0]), "[") {
					closeBodyLabel = "]"
				}
			}
			count++
			contentLength -= (len(ln) + count)
			if contentLength <= 0 {
				if !strings.Contains(ln, closeBodyLabel) {
					ln += closeBodyLabel
				}
				bodyString += strings.TrimSpace(ln)
				break
			}
			bodyString += strings.TrimSpace(ln)
		}
		body = append(body, strings.TrimSpace(bodyString))
	}

	req.Body = body

	fmt.Println("RESULT BODY ROUTER: ", req.Body)

	handler(&response, &req)
}

func parseRequest(header []string) Request {
	req := Request{}
	req.Headers = make(map[string]string)
	for i, h := range header {
		if i == 0 {
			spl := strings.Split(h, " ")
			req.Method = spl[0]
			req.Route = spl[1]
		} else {
			spl := strings.Split(h, ": ")
			if spl[0] == "Host" {
				req.Hosts = spl[1]
			} else {
				req.Headers[spl[0]] = spl[1]
			}
		}
	}
	return req
}

func NotFoundHandler(response *response.Response, request *Request) {
	err := errors.New("The resource requested was not found")

	response.SendData(404, templateHandler.GetErrorViewTemplate(err))
}
