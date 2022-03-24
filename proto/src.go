package server

type User interface {
	Createuser(request Request) Response
}

type Request struct {
	// @inject_tag: required:"true" validate:"required"
	Name string
	Id   int
}
type Response struct {
	Result string
}
