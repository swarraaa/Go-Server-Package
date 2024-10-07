package goServer

import "net/http"

type node struct {
    children map[string]*node
    handler  http.HandlerFunc
}

type Router struct {
    root *node
}

func NewRouter() *Router {
    return &Router{root: &node{children: make(map[string]*node)}}
}

func (r *Router) AddRoute(path string, handler http.HandlerFunc) {
    currentNode := r.root
    for _, part := range splitPath(path) {
        if _, ok := currentNode.children[part]; !ok {
            currentNode.children[part] = &node{children: make(map[string]*node)}
        }
        currentNode = currentNode.children[part]
    }
    currentNode.handler = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    currentNode := r.root
    for _, part := range splitPath(req.URL.Path) {
        if _, ok := currentNode.children[part]; ok {
            currentNode = currentNode.children[part]
        } else {
            http.NotFound(w, req)
            return
        }
    }
    if currentNode.handler != nil {
        currentNode.handler(w, req)
    } else {
        http.NotFound(w, req)
    }
}

func splitPath(path string) []string {
    return []string{}  
}
