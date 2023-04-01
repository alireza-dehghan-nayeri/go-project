package service

import (
	"github.com/alireza-dehghan-nayeri/go-project/api/repository"
	"github.com/alireza-dehghan-nayeri/go-project/models"
)

// QuoteService QuoteService struct
type QuoteService struct {
	repository repository.QuoteRepository
}

// NewQuoteService : returns the QuoteService struct instance
func NewQuoteService(r repository.QuoteRepository) QuoteService {
	return QuoteService{
		repository: r,
	}
}

// Save -> calls quote repository save method
func (repo QuoteService) Save(quote models.Quote) error {
	return repo.repository.Save(quote)
}

// FindAll -> calls quote repo find all method
func (repo QuoteService) FindAll(quote models.Quote, keyword string) (*[]models.Quote, int64, error) {
	return repo.repository.FindAll(quote, keyword)
}

// Update -> calls quoterepo update method
func (repo QuoteService) Update(quote models.Quote) error {
	return repo.repository.Update(quote)
}

// Delete -> calls quote repo delete method
func (repo QuoteService) Delete(id int64) error {
	var quote models.Quote
	quote.ID = id
	return repo.repository.Delete(quote)
}

// Find -> calls quote repo find method
func (repo QuoteService) Find(quote models.Quote) (models.Quote, error) {
	return repo.repository.Find(quote)
}
