package koa

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
	OPTION
)

type controller interface{}

type node struct {
	segment    string
	methods    []Method
	controller controller
	children   []*node
}

type Router struct {
	tree *node
}

func NewRouter() *Router {
	dumbNode := node{
		children: make([]*node, 0),
	}
	return &Router{&dumbNode}
}

func validatePath(path string) []string {

}

// expect to define a new path in router tree
func (r Router) Get(path string, controller controller) *Router {
	segs := validatePath(path)
	// pathToRegexp
}
