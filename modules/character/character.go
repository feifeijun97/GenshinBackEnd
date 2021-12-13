package character

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/feifeijun97/GenshinBackEnd/repository"
)

const (
	Male        string = "male"
	MaleCode    int    = 0
	Female      string = "female"
	FemaleCode  int    = 1
	UnknownCode int    = 2
)

const (
	POTRAIT = 1
)

type Character struct {
	ID           int
	Name         string
	Title        string `gorm:"column:title"`
	Description  string
	WeaponTypeId int `gorm:"column:weapon_type_id"`
	ElementId    int `gorm:"column:element_id"`
	Gender       int8
	BirthMonth   int8   `gorm:"column:birth_month"`
	BirthDay     int8   `gorm:"column:birth_day"`
	Affilation   string `gorm:"column:affiliation"`
	RegionId     int    `gorm:"column:region_id"`
	Substat      string
	Rarity       int
}

type CharacterJson struct {
	Name        string `json:"name"`
	Rarity      int    `json:"rarity"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Substat     string `json:"substat"`
	WeaponType  string `json:"weapon_type"`
	Region      string `json:"region"`
	Affiilation string `json:"affiliation"`
	Element     string `json:"element"`
	Gender      string `json:"gender"`
	Birthday    []int  `json:"birthday"`
}

type ImageStorage struct {
	Table      string
	TableValue int
	Type       int8
	Location   string
	FileName   string
}

type Element struct {
	Id   int
	Name string
}

type WeaponType struct {
	Id   int
	Name string
}

func (c *Character) String() string {
	return fmt.Sprintf("Character<%d %s>", c.ID, c.Name)
}

func (c *Character) GetCharacterById(id int) {
	tx := repository.Conn.First(&c, id)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func GenerateCharactersFromJson(folderPath string) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(files)
	}

	characters := []Character{}
	for _, file := range files {
		jsonFile, err := os.Open(folderPath + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		byteCharacter, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic(err)
		}

		var character CharacterJson

		//load the character byte into struct
		json.Unmarshal(byteCharacter, &character)

		//convert from json struct to db struct
		c := convertToDbStruct(character)

		//append into slice
		characters = append(characters, c)
	}

	//Batch insert character
	tx := repository.Conn.Create(&characters)
	if tx.Error != nil {
		panic(tx)
	}
}

// Convert the character json struct to struct that compatible to the database
// Eg: convert region to region_id instead
func convertToDbStruct(cj CharacterJson) Character {
	c := Character{}

	c.Name = cj.Name
	c.Title = cj.Title
	c.Description = cj.Description
	c.BirthDay = int8(cj.Birthday[0])
	c.BirthMonth = int8(cj.Birthday[1])
	c.Substat = cj.Substat
	c.Rarity = cj.Rarity
	c.Affilation = cj.Affiilation

	switch cj.Gender {
	case Male:
		c.Gender = int8(MaleCode)
	case Female:
		c.Gender = int8(FemaleCode)
	default:
		c.Gender = int8(UnknownCode)
	}

	switch cj.Region {
	case "Mondstadt":
		c.RegionId = 1
	case "Liyue":
		c.RegionId = 2
	case "Inazuma":
		c.RegionId = 3
	case "Snezhnaya", "Fatui":
		c.RegionId = 4
	default:
		c.RegionId = 0
	}

	element := Element{}
	tx := repository.Conn.Select("id").Where("name LIKE ?", "%"+cj.Element+"%").First(&element)
	if tx.Error != nil {
		panic(tx.Error)
	}
	c.ElementId = element.Id

	weaponType := WeaponType{}
	tx = repository.Conn.Select("id").Where("name LIKE ?", "%"+cj.WeaponType+"%").First(&weaponType)
	if tx.Error != nil {
		panic(tx.Error)
	}
	c.WeaponTypeId = weaponType.Id

	return c

}

func CreateCharacterPotraitImages(assetFolderPath string) {
	characters := []Character{}
	tx := repository.Conn.Find(&characters)
	if tx.Error != nil {
		panic(tx.Error)
	}

	for _, c := range characters {
		imageUrl := assetFolderPath + "\\assets\\images\\characters\\" + c.Name + "\\portrait"
		_, err := os.Stat(imageUrl)
		if err != nil {
			continue
		}
		source, err := os.Open(imageUrl)
		if err != nil {
			panic(err)
		}

		//create folder
		destDir := "src/images/characters/" + c.Name
		err = os.Mkdir(destDir, 0777)
		if err != nil {
			panic(err)
		}

		fileName := "potrait.jpg"
		destDir += "/" + fileName
		dest, err := os.Create(destDir)
		if err != nil {
			panic(err)
		}
		io.Copy(dest, source)
		imgStorage := ImageStorage{Table: "character", TableValue: c.ID, Location: destDir, FileName: fileName, Type: POTRAIT}

		tx = repository.Conn.Create(imgStorage)
		if tx.Error != nil {
			panic(tx.Error)
		}
		source.Close()
		dest.Close()
	}
}

// func tableName() string {
// 	return "character"
// }
