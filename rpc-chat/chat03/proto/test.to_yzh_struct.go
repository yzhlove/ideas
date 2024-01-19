package proto

type Req struct { //

	Name string `json:"Name"`

	Token int64 `json:"Token"`
}
type Resp struct { //

	Content string `json:"Content"`

	Timestamp int64 `json:"Timestamp"`
}
