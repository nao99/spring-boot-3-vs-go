package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDbSession(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)

	session, openingSessionError := gorm.Open(dialector, &gorm.Config{})
	if openingSessionError != nil {
		return nil, openingSessionError
	}

	return session, nil
}
