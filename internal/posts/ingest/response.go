package ingest

type (
	Response struct {
		Id    interface{} `json:"id,omitempty"`
		Error *Error      `error:"error,omitempty"`
	}
	Error struct {
		Msg  string `json:"msg,omitempty"`
		Code int32  `json:"code,omitempty"`
	}
)
