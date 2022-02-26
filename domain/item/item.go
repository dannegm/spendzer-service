package item

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrInvalidItem is returned when the Item is not valid
	ErrInvalidItem = errors.New("Is not a valid Item")
)

type Item struct {
	ID   uuid.UUID
	Name string
}

func NewItem(name string) (Item, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Item{}, ErrInvalidItem
	}

	// Create an item and initialize all the values to avoid nil pointer exceptions
	return Item{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

func (i *Item) GetID() uuid.UUID {
	return i.ID
}

func (i *Item) SetID(id uuid.UUID) {
	if i == nil {
		i = &Item{}
	}
	i.ID = id
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) SetName(name string) {
	if i == nil {
		i = &Item{}
	}
	i.Name = name
}
