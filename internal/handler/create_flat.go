package handler

import (
	"github.com/GlushenkovGleb/house-service/internal/handler/dto"
	"github.com/go-chi/render"
	"net/http"
)

func (c *Client) CreateFlat(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateFlatRequest{}
	render.Bind(r, &req)
	flat := req.ToFlat()

	if err := c.action.CreateFlat(flat); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	render.Render(w, r, dto.ToCreateFlatResponse(flat))
}
