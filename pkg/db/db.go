package db

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfingDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(cnf ConfingDB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
	)
	// TODO implement the jwt and then delete this
	fmt.Println(jwt.ErrEd25519Verification)
	db, err := sqlx.Connect("postgres", dsn)
	return db, err
}
