package ingest

type (
	Response struct {
		Success bool        `json:"success,omitempty"`
		Id      interface{} `json:"id,omitempty"`
		Error   *Error      `json:"error,omitempty"`
	}
	Error struct {
		Msg  string `json:"msg,omitempty"`
		Code int32  `json:"code,omitempty"`
	}
)
