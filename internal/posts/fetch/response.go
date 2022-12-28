package fetch

type (
	Response struct {
		Posts []Post `json:"posts,omitempty"`
		Error *Error `json:"error,omitempty"`
	}

	Error struct {
		Msg  string `json:"msg"`
		Code int32  `json:"code"`
	}
)
