package airlineseatmap

import (
	"net/http"

	"github.com/rushbuilding/airlineseatmap/app"
	"github.com/rushbuilding/airlineseatmap/controller"
	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/repository"
	"github.com/rushbuilding/airlineseatmap/service"
)

func main() {

	db := connections.GetConnection()
	cabinRepository := repository.NewAircraftManagementRepositoryImpl(db)
	passengerRepository := repository.NewPassengerRepository(db)
	segmentRepository := repository.NewSegmentManagementRepository(db)

	cabinService := service.NewCabinService(cabinRepository, segmentRepository)
	passengerService := service.NewPassengerService(passengerRepository)
	segmentSercvice := service.NewSegmentService(segmentRepository)

	seatingController := controller.NewSeatController(cabinService, passengerService, segmentSercvice)
	router := app.NewRouter(seatingController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}