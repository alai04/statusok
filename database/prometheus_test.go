package database

import (
	"testing"
)

func TestAddValidPrometheus(t *testing.T) {
	ids := make(map[int]int64, 0)
	ids[1] = 10
	ids[2] = 2

	Initialize(ids, 1, 10)

	err := AddNew(DatabaseTypes{Prometheus: &Prometheus{Port: "18080"}})

	if len(dbList) != 1 {
		t.Error("Not able to add databse to list")
	}
	if err != nil {
		t.Error("Init valid prometheus failed")
	}
}

func TestAddInValidPrometheus(t *testing.T) {
	ids := make(map[int]int64, 0)
	ids[1] = 10
	ids[2] = 2

	Initialize(ids, 1, 10)

	err := AddNew(DatabaseTypes{Prometheus: &Prometheus{Port: "a"}})
	if err == nil {
		t.Error("Init invalid prometheus should not succeed")
	}

	err = AddNew(DatabaseTypes{Prometheus: &Prometheus{Port: "100"}})
	if err == nil {
		t.Error("Init invalid prometheus should not succeed")
	}

	err = AddNew(DatabaseTypes{Prometheus: &Prometheus{Port: "100000"}})
	if err == nil {
		t.Error("Init invalid prometheus should not succeed")
	}
}
