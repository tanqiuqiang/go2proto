package server

type User interface {
	Createuser(request Request) Response
}

type Request struct {
	Name string
}
type Response struct {
	Result string
}
