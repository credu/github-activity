package models

type Payload struct {
	Action  string
	Number  int
	Size    int
	RefType string `json:"ref_type"`
}

type Repo struct {
	Id   int
	Name string
	Url  string
}

type Activity struct {
	Type    string
	Payload Payload
	Repo    Repo
}
