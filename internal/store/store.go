package store

import (
	"fmt"
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store struct {
	db *sqlx.DB
}

const initTableHousesSt = `
CREATE TABLE IF NOT EXISTS houses (
	id serial PRIMARY KEY,
	address TEXT NOT NULL,
	year INTEGER NOT NULL,
	developer VARCHAR(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_address UNIQUE (address)
)
`

const createEnumTypeSt = `
CREATE TYPE FLAT_STATUS AS ENUM ('created', 'approved', 'declined', 'on moderation')
`

const initTableFlatsSt = `
CREATE TABLE IF NOT EXISTS flats (
    id serial PRIMARY KEY,
    number INTEGER NOT NULL,
    house_id INTEGER NOT NULL REFERENCES houses (id) ON DELETE CASCADE,
    price INTEGER NOT NULL,
    rooms INTEGER NOT NULL,
    status FLAT_STATUS NOT NULL DEFAULT 'created',
    CONSTRAINT unique_house_id_id UNIQUE (number, house_id)
)
`

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func New() (*Store, error) {
	db, err := sqlx.Connect("pgx", "postgres://user:password@localhost/house?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// init houses
	db.MustExec(initTableHousesSt)
	log.Println("init houses")

	//// init enum
	//db.MustExec(createEnumTypeSt)
	//log.Println("init enums")

	// init flats
	db.MustExec(initTableFlatsSt)
	log.Println("init flats")

	return &Store{db: db}, nil
}

func Insert(table string, tuple domain.Insertable) squirrel.InsertBuilder {
	cols, vals := tuple.Args()
	fmt.Println(vals)
	return psql.Insert(table).
		Columns(cols...).
		Values(vals...)
}
