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
		where = fmt.Sprintf(`select id, name, description, created_at, updated_at from therapists where name = %s`, therapist[0])
	}
	
	query := fmt.Sprintf(`select id, name, description, created_at, updated_at from therapists %s`, where)
	
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
