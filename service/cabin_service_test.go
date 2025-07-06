package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/repository"
)

var cabinService = NewCabinService(repository.NewAircraftManagementRepositoryImpl(connections.GetConnection()))
func TestXxx(t *testing.T) {
	res, _ := cabinService.GetSeatingMapAndPriceByCabinId(context.Background(), "fd7f5c51-0bc2-417c-a228-9b0ffd97e43c", "ECONOMY", "MYR")
	jsonData, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonData))	
}