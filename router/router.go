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
	fmt.Println(router)
	scanner := bufio.NewScanner(conn)
	var header, body []string
	f := true
	contentLength := 0

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			f = false
			if strings.HasPrefix(header[0], "POST") || strings.HasPrefix(header[0], "PUT") || strings.HasPrefix(header[0], "DELETE") {
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

	// // Parse the request to extract the query string
	// requestParts := strings.Split(req.Route, " ")
	// if len(requestParts) >= 2 {
	// 	requestLine := requestParts[1]
	// 	pathAndQuery := strings.SplitN(requestLine, "?", 2)
	// 	path := pathAndQuery[0]
	// 	query := ""
	// 	if len(pathAndQuery) == 2 {
	// 		query = pathAndQuery[1]
	// 	}
	// 	queryParams := parseQuery(query)

	// 	fmt.Println("PARAM ", path)
	// 	fmt.Println(queryParams)
	// }

	fmt.Println("METHOD ", req.Method)
	if req.Method == "POST" || req.Method == "PUT" || req.Method == "DELETE" {
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
				fmt.Println("LAST READ :", contentLength)
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

func parseQuery(query string) map[string]string {
	params := make(map[string]string)
	keyValuePairs := strings.Split(query, "&")

	for _, pair := range keyValuePairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			paramName := parts[0]
			paramValue := parts[1]
			params[paramName] = paramValue
		}
	}

	return params
}

func NotFoundHandler(response *response.Response, request *Request) {
	err := errors.New("The resource requested was not found")

	response.SendData(404, templateHandler.GetErrorViewTemplate(err))
}
