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

func NewStorage() (*Storage, error) {
	db, err := memdb.NewMemDB(schema())
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Insert(p *domain.Pet) error {
	txn := s.db.Txn(true)

	last, err := txn.Last("pet", "id")
	if err != nil {
		txn.Abort()
		return err
	}

	id := 0
	if last != nil {
		id = last.(*domain.Pet).Id
	}

	p.Id = id + 1
	if err := txn.Insert("pet", p); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (s *Storage) Get(id int) (*domain.Pet, error) {
	txn := s.db.Txn(false) // read-only transactipm

	raw, err := txn.First("pet", "id", id)
	if err != nil {
		txn.Abort()
		return nil, err
	}

	if raw == nil {
		txn.Abort()
		return nil, fmt.Errorf("no content")
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
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}
}
