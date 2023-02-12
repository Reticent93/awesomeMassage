package dbrepo

import (
	"awesomeMassage/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) GetTherapists(therapist ...string) ([]*models.Therapist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	where := ""
	if len(therapist) > 0 {
		where = fmt.Sprintf(`select id, name, created_at, updated_at from therapists where name=%s`, therapist[0])
	}

	query := fmt.Sprintf(`select id, name,  created_at, updated_at from therapists %s`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var therapists []*models.Therapist

	for rows.Next() {
		var therapist models.Therapist
		err := rows.Scan(&therapist.ID, &therapist.TherapistName, &therapist.CreatedAt, &therapist.UpdatedAt)
		if err != nil {
			return nil, err
		}
		therapists = append(therapists, &therapist)
	}
	return therapists, err
}

func (m *PostgresDBRepo) OneTherapist(id int) (*models.Therapist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, created_at, updated_at from therapists where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var therapist models.Therapist

	err := row.Scan(&therapist.ID, &therapist.TherapistName, &therapist.CreatedAt, &therapist.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &therapist, nil
}

func (m *PostgresDBRepo) GetUsers(user ...string) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	where := ""
	if len(user) > 0 {
		where = fmt.Sprintf(`select id, first_name,last_name, email, password, role, created_at, updated_at from users where email=%s`, user[0])
	}

	query := fmt.Sprintf(`select id, first_name,last_name, email, password, role, created_at, updated_at from users %s`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil

}

func (m *PostgresDBRepo) OneUser(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, first_name,last_name, email, password, role, created_at, updated_at from users where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var user models.User

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
