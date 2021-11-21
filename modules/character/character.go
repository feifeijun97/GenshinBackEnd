package character

import (
	"fmt"
)

type Character struct {
	// bun.BaseModel `bun:"character,alias:c"`
	ID   int64
	Name string
}

func (c Character) String() string {
	return fmt.Sprintf("Character<%d %s>", c.ID, c.Name)
}

// func tableName() string {
// 	return "character"
// }
