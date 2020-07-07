package request

type Request interface {
	JsonParams() (string, error)
}
