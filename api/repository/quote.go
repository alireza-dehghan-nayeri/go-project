package repository

import (
	"github.com/alireza-dehghan-nayeri/go-project/infrastructure"
	"github.com/alireza-dehghan-nayeri/go-project/models"
)

// QuoteRepository -> QuoteRepository
type QuoteRepository struct {
	db infrastructure.Database
}

// NewQuoteRepository : fetching database
func NewQuoteRepository(db infrastructure.Database) QuoteRepository {
	return QuoteRepository{
		db: db,
	}
}

// Save -> Method for saving quote to database
func (p QuoteRepository) Save(quote models.Quote) error {
	return p.db.DB.Create(&quote).Error
}

// FindAll -> Method for fetching all quotes from database
func (repo QuoteRepository) FindAll(quote models.Quote, keyword string) (*[]models.Quote, int64, error) {
	var quotes []models.Quote
	var totalRows int64 = 0

	queryBuider := repo.db.DB.Order("created_at desc").Model(&models.Quote{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			repo.db.DB.Where("quote.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(quote).
		Find(&quotes).
		Count(&totalRows).Error
	return &quotes, totalRows, err
}

// Update -> Method for updating Quote
func (repo QuoteRepository) Update(quote models.Quote) error {
	return repo.db.DB.Save(&quote).Error
}

// Find -> Method for fetching quote by id
func (repo QuoteRepository) Find(quote models.Quote) (models.Quote, error) {
	var quotes models.Quote
	err := repo.db.DB.
		Debug().
		Model(&models.Quote{}).
		Where(&quote).
		Take(&quotes).Error
	return quotes, err
}

// Delete Deletes quote
func (repo QuoteRepository) Delete(quote models.Quote) error {
	return repo.db.DB.Delete(&quote).Error
}
