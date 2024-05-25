package community

import (
	"github.com/google/uuid"
)

type Community struct {
	Name     string
	Icon     string
	Id       uuid.UUID
	Users    []uuid.UUID
	Channels []uuid.UUID
}
