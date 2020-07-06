package database

import (
	"testing"
)

func TestAddValidDatabase(t *testing.T) {
	ids := make(map[int]int64, 0)
	ids[1] = 10
	ids[2] = 2

	Initialize(ids, 1, 10)

	influxDb := InfluxDb{"localhost", 8086, "statusok", "", ""}

	databaseTypes := DatabaseTypes{InfluxDb: influxDb}

	AddNew(databaseTypes)

	if len(dbList) != 1 {
		t.Error("Not able to add databse to list")
	}
}
