package repository

import (
	"context"
	"gol/domain"

	"github.com/jackc/pgx/v5"
)

type ContactRepository struct { 
  DB *pgx.Conn
}

func NewContactRepository(db *pgx.Conn) *ContactRepository {
	return &ContactRepository{
		DB: db,
	}
}

func (r *ContactRepository) Create (data domain.Contact) error { 
  _, err := r.DB.Query(
    context.Background(),`
    INSERT INTO contact (name, email, phone, address) VALUES 
    ($1, $2, $3, $4)
    `, data.Name, data.Email, data.Phone, data.Address,
  )
  return err 
}

func (r* ContactRepository) Update (data domain.Contact) error{
  _, err := r.DB.Query(
    context.Background(),`
      UPDATE contact SET name = $2, email = $3, phone = $4, address = $5 WHERE id = $1
    `, data.Id, data.Name, data.Email, data.Phone, data.Address,
  )
  return err 
}

func (r *ContactRepository) Delete (data domain.Contact) error { 
  _, err := r.DB.Exec(
    context.Background(),`
    DELETE FROM contact WHERE id = $1
    `, data.Id,
  )
  return err 
}

func (r *ContactRepository) FindAll() ([]domain.Contact, error) { 
  rows, err := r.DB.Query(
    context.Background(),`
      SELECT id, name, email, phone, address FROM contact
    `,
  )
  var result []domain.Contact
  for rows.Next() { 
    var p domain.Contact 
    err = rows.Scan(&p.Id, &p.Name, &p.Email, &p.Phone, &p.Address)
    result = append(result, p)
  }
  return result, err
}