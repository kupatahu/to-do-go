package todo

import (
	"github.com/kupatahu/to-do-go/entity"
	"github.com/kupatahu/to-do-go/httperr"
)

type Repo interface {
	Get() ([]*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(todo *entity.Todo) (*entity.Todo, httperr.HttpErr)
}

type usecase struct {
	repo Repo
}

func newUsecase(repo Repo) *usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Get() ([]*entity.Todo, error) {
	return u.repo.Get()
}

func (u *usecase) Create(todo *entity.Todo) (*entity.Todo, error) {
	return u.repo.Create(todo)
}

func (u *usecase) Update(todo *entity.Todo) (*entity.Todo, httperr.HttpErr) {
	return u.repo.Update(todo)
}
