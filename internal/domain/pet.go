package domain

import (
	"fmt"
	"time"
)

 
type Pet struct {
	Name     string `json:"name"`
	Kind      string `json:"kind"`
	Gender   string `json:"gender"`
	BirthDay string `json:"dateOfBirth"`
}


func (p *Pet) IsValidDate() error {
	d, err := time.Parse(time.DateOnly, p.BirthDay)
	if err != nil {
		return err
	}

	if d.After(time.Now()) {
		return fmt.Errorf("dateOfBirth must be before now")
	}
	return nil
}

