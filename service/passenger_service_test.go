package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/repository"
)


var psgservice = NewPassengerService(repository.NewPassengerRepository(connections.GetConnection()))

func TestService(t *testing.T) {
	res, _ := psgservice.GetPassengerData(context.Background(), "3937094243023851424", "8fe24407-1df6-4b99-96c9-4dee6b72fdfb")
	jsonData, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonData))
	
}