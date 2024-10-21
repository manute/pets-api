package storage_test

import (
	"fmt"
	"iskaypet-challenge/internal/domain"
	"iskaypet-challenge/internal/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Upsert(t *testing.T) {
	db, err := storage.NewStorage()
	if err != nil {
		t.Fatal(err)
	}

	id := "tango"
	p := domain.Pet{
		Name:     id,
		Kind:     "labrador",
		Gender:   "male",
		BirthDay: "25/03/2012",
	}

	if err := db.Upsert(&p); err != nil {
		t.Error(err)
	}

	pet, err := db.Get(id)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(&p, pet) {
		t.Errorf("not equal entities")
	}
}

func Test_List(t *testing.T) {
	db, err := storage.NewStorage()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		id := fmt.Sprintf("tango_%d", i)
		p := domain.Pet{
			Name:     id,
			Kind:     "labrador",
			Gender:   "male",
			BirthDay: "25/03/2012",
		}

		if err := db.Upsert(&p); err != nil {
			t.Error(err)
		}
	}

	pets, err := db.List()
	if err != nil {
		t.Error(err)
	}

	if len(pets) != 5 {
		t.Errorf("not list entities retrieved, size 5 != %d", len(pets))
	}

}
