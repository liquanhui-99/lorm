package lorm

type Router struct {
	trees map[string]*node
}

type node struct {
	// 路由地址
	path string
	// 子路径
	children map[string]*node
	// 业务逻辑的处理程序
	handle HandleFunc
}

func newRouter() *Router {
	return &Router{
		trees: map[string]*node{},
	}
}
