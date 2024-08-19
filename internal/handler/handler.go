package handler

import (
	"github.com/GlushenkovGleb/house-service/internal/action"
	"net/http"
)

type Handler interface {
	CreateHouse(w http.ResponseWriter, r *http.Request)
	CreateFlat(w http.ResponseWriter, r *http.Request)
	UpdateFlat(w http.ResponseWriter, r *http.Request)
	GetFlats(w http.ResponseWriter, r *http.Request)
}

type Client struct {
	action action.Action
}

func (c *Client) GetFlats(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ Handler = (*Client)(nil)

func New(action action.Action) *Client {
	return &Client{
		action: action,
	}
}
