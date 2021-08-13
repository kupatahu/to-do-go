package entity

type Todo struct {
	Id    string `json:"id"`
	Value string `json:"value"`
	Done  bool   `json:"done"`
}

func New(id string, value string) *Todo {
	return &Todo{Id: id, Value: value}
}
