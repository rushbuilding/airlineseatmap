package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/repository"
)

var con = connections.GetConnection()

var cabinService = NewCabinService(
	repository.NewAircraftManagementRepositoryImpl(con),
	repository.NewSegmentManagementRepository(con))
func TestXxx(t *testing.T) {
	res, _ := cabinService.GetSeatingMapAndPriceByCabinId(context.Background(), "3937094243023851424", "ECONOMY", "MYR")
	jsonData, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonData))	
}