package gie

type Router struct {
	Url2Handler map[string]HandlerFunc
}
