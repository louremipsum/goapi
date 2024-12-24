// Package tools provides database interfaces and implementations
package tools

import (
	log "github.com/sirupsen/logrus"
)

// LoginDetails stores user authentication information
type LoginDetails struct {
	AuthToken string
	Username  string
}

// CoinDetails stores user coin balance information
type CoinDetails struct {
	Coins    int64
	Username string
}

// DatabaseInterface defines the required methods for database operations
type DatabaseInterface interface {
	// GetUserLoginDetails retrieves user authentication details
	GetUserLoginDetails(username string) *LoginDetails
	// GetUserCoins retrieves user coin balance
	GetUserCoins(username string) *CoinDetails
	// SetupDatabase initializes the database connection
	SetupDatabase() error
}

// NewDatabase creates a new database instance
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
