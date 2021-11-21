package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bun"
)

//Connect To PostgreSQL db using dotenv db credentials
// func ConnectToPostgreSQL() {
// 	//declare the psql info for connecting db
// 	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		os.Getenv("DB_HOST"), dbPort, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
// }

func ConnectToPostgreDb() {
	context := context.Background()

	// create a dsn for postgre DB
	// the format of dsn can refer to : https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	pq_db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	//close the db at at last using defer
	defer pq_db.Close()

	db := bun.NewDB(pq_db, pgdialect.New())

	//set the table name to singular (disable automatically add 's' after the table name)
	character := new(Character)
	if err := db.NewSelect().Model(character).
		Where("id = ?", 1).Scan(context); err != nil {
		panic(err)
	}

	fmt.Print(character.String())
}

type Character struct {
	bun.BaseModel `bun:"character,alias:c"`
	ID            int64
	Name          string
}

func (c Character) String() string {
	return fmt.Sprintf("Character<%d %s>", c.ID, c.Name)
}
