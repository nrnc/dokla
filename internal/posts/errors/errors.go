package errors

import "fmt"

// Error represents any error in the application
type Error interface {
	Code() int
	Status() string
	Message() string
}

type DoklaError struct {
	ErrCode   int    `json:"code"`
	ErrStatus string `json:"status"`
	ErrMsg    string `json:"message"`
}

type FetchError struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"message"`
}

func NewDoklaError(c int, s, m string) *DoklaError {
	return &DoklaError{
		ErrStatus: s,
		ErrCode:   c,
		ErrMsg:    m,
	}
}

func (e *FetchError) Error() string {
	return fmt.Sprintf(
		"{\"error\":{\"code\":%d, \"status\": %s}}",
		e.ErrCode,
		e.ErrMsg,
	)
}

func (e *DoklaError) Code() int {
	return e.ErrCode
}

func (e *DoklaError) Status() string {
	return e.ErrStatus
}

func (e *DoklaError) Message() string {
	return e.ErrMsg
}

func (e *DoklaError) Error() string {
	return fmt.Sprintf(
		"{\"error\":{\"code\":%d,\"message\": %s, \"status\": %s}}",
		e.Code(),
		e.Message(), e.Status())
}

//used by http.StatusCoder
func (e *DoklaError) StatusCode() int {
	return e.Code()
}

// Service Error Constants
var (
	ErrFetchFailed  = DoklaError{500, "", "fetch failed"}
	ErrIngestFailed = DoklaError{500, "", "ingest failed"}
	DEFAULT_ERROR   = "Internal Server Error"
)
