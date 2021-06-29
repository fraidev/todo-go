package entities

import "github.com/google/uuid"

type Todo struct {
	ID        uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	Name      string
	Completed bool
}
