package common

type Resp struct {
	Data interface{}
	Err  error
}

const (
	Success      = 200
	NotFound     = 400
	Unauthorized = 401
	ServerError  = 500
)
