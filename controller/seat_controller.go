package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SeatingController interface {
	GetPassangerSeating(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}