package models

import (
	"time"
)

// Quote -> Quote struct to save Quote on database
type Quote struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:3000" json:"body" `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method that returns tablename of Quote model
func (quote *Quote) TableName() string {
	return "quote"
}

// ResponseMap -> response map of Quote
func (quote *Quote) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = quote.ID
	resp["title"] = quote.Title
	resp["body"] = quote.Body
	resp["created_at"] = quote.CreatedAt
	resp["updated_at"] = quote.UpdatedAt
	return resp

}
