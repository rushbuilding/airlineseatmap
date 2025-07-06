package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rushbuilding/airlineseatmap/controller"
)

func NewRouter(seatController controller.SeatingController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/seatmap", seatController.GetPassangerSeating)

	return router
}