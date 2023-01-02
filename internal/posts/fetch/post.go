package fetch

type Post struct {
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Username   string `json:"username,omitempty" bson:"username,omitempty"`
	Content    string `json:"content,omitempty" bson:"content,omitempty"`
	Title      string `json:"title,omitempty" bson:"title,omitempty"`
	CreatedAt  string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Avatar     string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	PostId     string `json:"post_id,omitempty" bson:"post_id,omitempty"`
	PathParams `json:"params,omitempty" bson:"inline"`
	Meta       map[string]interface{} `json:"meta" bson:"meta,omitempty"`
}

type PathParams struct {
	App    string `json:"app,omitempty" bson:"app,omitempty"`
	Tenant string `json:"tenant,omitempty" bson:"tenant,omitempty"`
	Source string `json:"source,omitempty" bson:"source,omitempty"`
}
