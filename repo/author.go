package repo

import (
	"github.com/suaas21/graphql-dummy/model"
)

type AuthorRepository interface {
	CreateAuthor(author *model.Author) error
	GetAuthorByID(key string) (*model.Author, error)
	ListAuthorByIDs(keys []string) ([]model.Author, error)
	UpdateAuthor(author model.Author) error
	DeleteAuthor(id uint) error
}
