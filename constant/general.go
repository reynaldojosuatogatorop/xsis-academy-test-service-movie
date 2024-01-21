package constant

const (
	TimeLocation     = "Asia/Jakarta"
	MethodNotAllowed = "Method Not Allowed"
)

type ResultError struct {
	Code InternalError
	Err  error
}
