package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func ConnectToPostgreDb() {
	// context := context.Background()

	// create a dsn for postgre DB
	// the format of dsn can refer to : https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// pgxConn, err := pgx.Connect(context, dsn)
	// if err != nil {
	// 	panic(err)
	// }
	// Conn = pgxConn
	// fmt.Println("Successfully establish database connection")

	// // defer Conn.Close(context)

	// establish new database connection using GORM
	// configure the table name as singular, default is plural
	pgxConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	Conn = pgxConn
	fmt.Println("Successfully establish database connection using gorm")

}
