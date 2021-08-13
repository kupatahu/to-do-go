package todo

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
	"github.com/kupatahu/to-do-go/entity"
	"github.com/kupatahu/to-do-go/httperr"
	"github.com/sdomino/scribble"
)

type repo struct {
	db *scribble.Driver
}

func newRepo(db *scribble.Driver) *repo {
	return &repo{db: db}
}

func (r *repo) Get() ([]*entity.Todo, error) {
	todos := []*entity.Todo{}
	records, err := r.db.ReadAll("todo")

	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil
		}

		return nil, err
	}

	for _, record := range records {
		todo := &entity.Todo{}
		err := json.Unmarshal(record, todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *repo) Create(todo *entity.Todo) (*entity.Todo, error) {
	todo.Id = uuid.NewString()
	err := r.db.Write("todo", todo.Id, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repo) Update(todo *entity.Todo) (*entity.Todo, httperr.HttpErr) {
	err := r.db.Write("todo", todo.Id, todo)

	if err != nil {
		if err.Error() == "missing resource - unable to save record" {
			httperr := httperr.NewNotFound("todo "+todo.Id+"not exist", err)
			return nil, httperr
		}

		return nil, httperr.NewInternalServer("", err)
	}

	return todo, nil
}
