package request

type Request interface {
	JsonParams() (string, error)
	ResponseName() string
}

type Response interface {
}
