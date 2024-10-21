package domain

import (
	"fmt"
	"time"
)

type Pet struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (p *Pet) IsValidDate() error {
	d, err := time.Parse(time.DateOnly, p.DateOfBirth)
	if err != nil {
		return err
	}

	if d.After(time.Now()) {
		return fmt.Errorf("dateOfBirth must be before now")
	}
	return nil
}
