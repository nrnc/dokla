package ingest

type (
	PlayStore struct {
		Name       string `json:"name,omitempty"`
		GmailId    string `json:"mail,omitempty"`
		Review     string `json:"review,omitempty"`
		Title      string `json:"title,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		Avatar     string `json:"avatar,omitempty"`
		PostId     string `json:"post_id,omitempty"`
		PathParams `json:"params,omitempty"`
	}
)
