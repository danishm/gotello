package pkg

import (
	"testing"
)

func TestNewState(t *testing.T) {
	testStateStr := "pitch:0;roll:0;yaw:0;vgx:0;vgy:0;vgz:0;templ:89;temph:91;tof:10;h:0;bat:3;baro:48.26;time:0;agx:0.00;agy:-6.00;agz:-1003.00;"
	expectedState := State{
		Pitch:              0,
		Roll:               0,
		Yaw:                0,
		VGY:                0,
		VGX:                0,
		VGZ:                0,
		HighestTemperature: 91,
		LowestTemperature:  89,
		TOF:                10,
		Battery:            3,
		Barometer:          48.26,
		MotorsOnTime:       0,
		AGX:                0,
		AGY:                -6.00,
		AGZ:                -1003.00,
	}

	state, err := NewState(testStateStr)
	if err != nil {
		t.Fatal("Error parsing state string:", testStateStr)
	}

	if state != expectedState {
		t.Fatalf("Expected: %+v\nActual: %+v", expectedState, state)
	}
}
