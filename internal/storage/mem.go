package storage

import (
	// "fmt"

	"fmt"
	"iskaypet-challenge/internal/domain"

	"github.com/hashicorp/go-memdb"
)

type Storage struct {
	db *memdb.MemDB
}

type Repository interface {
	Upsert(p *domain.Pet) error
	Get(n string) (*domain.Pet, error)
}

func NewStorage() (*Storage, error) {
	db, err := memdb.NewMemDB(schema())
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Upsert(p *domain.Pet) error {
	txn := s.db.Txn(true)

	if err := txn.Insert("pet", p); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (s *Storage) Get(id string) (*domain.Pet, error) {
	txn := s.db.Txn(false) // read-only transactipm

	raw, err := txn.First("pet", "id", id)
	if err != nil {
		txn.Abort()
		return nil, err
	}

	pet, ok := raw.(*domain.Pet)
	if !ok {
		txn.Abort()
		return nil, fmt.Errorf("row is not a pet")
	}

	txn.Commit()
	return pet, nil
}

func (s *Storage) List() ([]*domain.Pet, error) {
	txn := s.db.Txn(false) // read-only transactipm

	it, err := txn.Get("pet", "id")
	if err != nil {
		txn.Abort()
		return nil, err
	}

	var res []*domain.Pet
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*domain.Pet)
		res = append(res, p)
	}

	txn.Commit()
	return res, nil
}

func schema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"pet": {
				Name: "pet",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}
}
