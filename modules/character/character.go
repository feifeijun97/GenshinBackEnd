package character

import (
	"fmt"

	"github.com/feifeijun97/GenshinBackEnd/repository"
)

type Character struct {
	ID   int64
	Name string
}

func (c Character) String() string {
	return fmt.Sprintf("Character<%d %s>", c.ID, c.Name)
}

func (c *Character) GetCharacterById(id int) {
	tx := repository.Conn.First(&c, id)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

// func tableName() string {
// 	return "character"
// }
