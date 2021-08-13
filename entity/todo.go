package entity

type Todo struct {
	Id    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Done  bool   `json:"done,omitempty"`
}

func New(id string, value string) *Todo {
	return &Todo{Id: id, Value: value}
}
