package handler

import (
	"github.com/GlushenkovGleb/house-service/internal/handler/dto"
	"github.com/go-chi/render"
	"net/http"
)

func (c *Client) UpdateFlat(w http.ResponseWriter, r *http.Request) {
	req := &dto.UpdateFlatRequest{}
	if err := render.Bind(r, req); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	flat, err := c.action.UpdateFlat(req.ID, req.Status)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	render.Render(w, r, dto.ToUpdateFlatResponse(flat))
}
