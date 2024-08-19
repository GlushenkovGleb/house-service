package main

import (
	"github.com/GlushenkovGleb/house-service/internal/action"
	"github.com/GlushenkovGleb/house-service/internal/handler"
	"github.com/GlushenkovGleb/house-service/internal/server"
	"github.com/GlushenkovGleb/house-service/internal/store"
)

func main() {
	// 1*. init logger
	// 2. init db
	db, err := store.New()
	if err != nil {
		panic(err)
	}
	// 3. init actions
	act := action.New(db)

	// 4. init handlers
	hand := handler.New(act)

	// 4. init and start server
	serv := server.New(hand)
	if servErr := serv.ListenAndServe(); servErr != nil {
		panic(servErr)
	}
}
