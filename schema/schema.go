package schema

import (
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/repo"
)

type BookAuthor struct {
	bookRepo   *repo.BookRepository
	authorRepo *repo.Author
	log        logger.StructLogger
}

func NewBookAuthor(bookRepo *repo.BookRepository, authorRepo *repo.Author, lgr logger.StructLogger) *BookAuthor {
	return &BookAuthor{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
		log:        lgr,
	}
}
