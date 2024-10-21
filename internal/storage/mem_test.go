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

	id := 1
	p := domain.Pet{
		// Id:          id,
		Name:        "tango",
		Kind:        "labrador",
		Gender:      "male",
		DateOfBirth: "25/03/2012",
	}

	if err := db.Insert(&p); err != nil {
		t.Error(err)
	}

	pet, err := db.Get(id)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(&p, pet) {
		t.Errorf("got bad entity")
	}
}

func Test_List(t *testing.T) {
	db, err := storage.NewStorage()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		p := domain.Pet{
			Id:          i,
			Name:        fmt.Sprintf("tango_%d", i),
			Kind:        "labrador",
			Gender:      "male",
			DateOfBirth: "25/03/2012",
		}

		if err := db.Insert(&p); err != nil {
			t.Error(err)
		}
	}

	pets, err := db.List()
	if err != nil {
		t.Error(err)
	}

	if len(pets) != 5 {
		t.Errorf("wrong list entities retrieved, size 5 != %d", len(pets))
	}

}
