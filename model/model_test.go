package model

import (
	"errors"
	"testing"
	"time"
)

func TestReducedMessage_Reduce(t *testing.T) {
	reducedMessage := ReducedMessage{
		[]string{"1", "2"},
		[]string{"3", "4"},
		[]string{"5", "6"},
	}

	if result, _ := reducedMessage.Reduce(func(current, next []string) (res []string, err error) {
		res = append(current, next...)
		return res, err
	}); len(result) != 6 {
		t.Fatalf("Array only has %d elements: %s", len(result), result)
	} else {
		t.Logf("Result array: %s", result)
	}
}

func TestReducedMessage_ReduceError(t *testing.T) {
	reducedMessage := ReducedMessage{
		[]string{"1", "2"},
		[]string{"1", "2"},
	}

	if _, err := reducedMessage.Reduce(func(current, next []string) (res []string, err error) {
		err = errors.New("this is an error")
		return res, err
	}); err == nil {
		t.Fatal("Error shouldn't be nil")
	} else if err.Error() != "this is an error" {
		t.Logf("Unexpected error: %s", err)
	} else {
		t.Logf("Error message: %s", err)
	}
}

func TestReducedMessage_ReduceLengthError(t *testing.T) {
	reducedMessage := ReducedMessage{}

	if _, err := reducedMessage.Reduce(func(current, next []string) (res []string, err error) {
		return res, err
	}); err == nil {
		t.Fatal("Error shouldn't be nil")
	} else if err.Error() != "length too short" {
		t.Logf("Unexpected error: %s", err)
	} else {
		t.Logf("Error message: %s", err)
	}
}

func TestSatellitesReqBody_Prepare(t *testing.T) {
	tempSatellites := []TempSatellite{
		{ID: "kenobi", X: 1, Y: 1, CreatedAt: time.Now()},
		{ID: "sato", X: 1, Y: 1, CreatedAt: time.Now()},
	}

	satellitesReqBody := SatellitesReqBody{
		Satellites: []Satellite{
			{Name: "kenobi", Distance: 1, Message: []string{"1"}},
			{Name: "sato", Distance: 2, Message: []string{"2"}},
		},
	}

	satellitesData := satellitesReqBody.Prepare(tempSatellites)

	if satellitesData[0].TempSatelliteID != "kenobi" &&
		satellitesData[1].TempSatelliteID != "sato" {
		t.Fatalf("Incompatible data")
	} else {
		t.Log(satellitesData)
	}
}
