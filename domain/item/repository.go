package item

import (
	"errors"
)

var (
	ErrFailedToAddItem = errors.New("failed to add the item to the repository")
)

type ItemRepository interface {
	Add(Item) error
}
