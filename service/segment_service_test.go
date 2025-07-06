package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/repository"
)

var sgmt = NewSegmentService(repository.NewSegmentManagementRepository(connections.GetConnection()))
func TestSegment(t *testing.T) {
	res, _ := sgmt.GetSegmentData(context.Background(), "3937094243023851424", "ECONOMY")
	jsonData, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonData))

}