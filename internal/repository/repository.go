package repository

import (
	"awesomeMassage/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetTherapists(therapist ...string) ([]*models.Therapist, error)
	OneTherapist(id int) (*models.Therapist, error)
	GetUsers(user ...string) ([]*models.User, error)
	OneUser(id int) (*models.User, error)
}
