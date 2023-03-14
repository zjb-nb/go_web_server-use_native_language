package server

import (
	"afterclass/server/router"
	"net/http"
	"strings"
)

//利用树形结构实现路由
// 根节点是""/"
type HandleBasedOnTree struct {
	root *Node
}

func NewHandleBasedOnTree() BaseHandle {
	return &HandleBasedOnTree{
		root: &Node{path: "/"},
	}
}

type Node struct {
	path     string
	children []*Node
	handler  MyHandleFunc
}

//实现路由的查找
func (h *HandleBasedOnTree) ServeHTTP(ctx *router.MyContext) {
	paths := strings.Split(strings.Trim(ctx.R.URL.Path, "/"), "/")
	cur := h.root
	for _, path := range paths {
		matchChild, found := h.findMatchChild(cur, path)
		// [home,sign] 当出现home找不到时,sign也没必有找
		if !found {
			ctx.W.WriteHeader(http.StatusNotFound)
			_, _ = ctx.W.Write([]byte("404 Page not found"))
		}
		cur = matchChild
	}

	if cur.handler == nil {
		ctx.W.WriteHeader(http.StatusNotFound)
		_, _ = ctx.W.Write([]byte("404 Page not found"))
	}
	cur.handler(ctx)
}

//实现路由的新增
/*
url : "/home/sign"
*/
func (h *HandleBasedOnTree) Router(method string, url string, handleFunc MyHandleFunc) {
	paths := strings.Split(strings.Trim(url, "/"), "/")
	cur := h.root
	// [home,sign]找到一个，下一个应该是去他儿子里面找
	for index, path := range paths {
		matchChild, found := h.findMatchChild(cur, path)
		if found {
			cur = matchChild
		} else {
			//没找到就在原来树上添加 注意如果home没有，后面的sign也要一起添加
			h.buildTree(cur, paths[index:], handleFunc)
			return
		}
	}
	cur.handler = handleFunc
}

func (h *HandleBasedOnTree) findMatchChild(root *Node, path string) (*Node, bool) {
	path = strings.Trim(path, "/")
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func (h *HandleBasedOnTree) buildTree(root *Node, paths []string, handleFunc MyHandleFunc) {
	cur := root
	for _, path := range paths {
		node := &Node{path: path}
		cur.children = append(cur.children, node)
		cur = node
	}
	cur.handler = handleFunc
}
