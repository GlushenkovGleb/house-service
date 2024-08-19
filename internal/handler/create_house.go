package handler

import (
	"fmt"
	"github.com/GlushenkovGleb/house-service/internal/handler/dto"
	"github.com/go-chi/render"
	"net/http"
)

func (c *Client) CreateHouse(w http.ResponseWriter, r *http.Request) {
	req := &dto.CreateHouseRequest{}
	if err := render.Bind(r, req); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	house := req.ToHouse()
	err := c.action.CreateHouse(house)
	fmt.Println(house)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	render.Render(w, r, dto.ToCreateHouseResponse(house))
	return
}
