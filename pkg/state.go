package pkg

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// StateSeparator is the character that saperates the parts of the state in the string from Tello
	StateSeparator = ";"

	// KeyValSeparator is the character that saperates the key/value pair in a state
	KeyValSeparator = ":"
)

// State represents the state of the Tello, it's expected to be in the following format e.g.
// "pitch:0;roll:0;yaw:0;vgx:0;vgy:0;vgz:0;templ:89;temph:91;tof:10;h:0;bat:3;baro:48.26;time:0;agx:0.00;agy:-6.00;agz:-1003.00;"
//
// Source: file:///Users/danish/Downloads/Tello%20SDK%20Documentation%20EN_1.3_1122.pdf
type State struct {
	Pitch              int
	Roll               int
	Yaw                int
	VGX                int
	VGY                int
	VGZ                int
	LowestTemperature  int
	HighestTemperature int
	TOF                int
	Height             int
	Battery            int
	Barometer          float64
	MotorsOnTime       int
	AGX                float64
	AGY                float64
	AGZ                float64
}

// NewState takes the string returned from the Tello API and returns an instance of State
func NewState(s string) (State, error) {

	state := State{}

	if len(s) == 0 {
		return state, errors.New("Empty state string")
	}

	states := strings.Split(s, StateSeparator)
	for _, keyVal := range states {
		parts := strings.Split(keyVal, KeyValSeparator)
		if len(parts) == 2 {
			key := parts[0]
			val := parts[1]

			switch key {
			case "pitch":
				state.Pitch, _ = strconv.Atoi(val)
			case "roll":
				state.Roll, _ = strconv.Atoi(val)
			case "yaw":
				state.Yaw, _ = strconv.Atoi(val)
			case "vgx":
				state.VGX, _ = strconv.Atoi(val)
			case "vgy":
				state.VGY, _ = strconv.Atoi(val)
			case "vgz":
				state.VGZ, _ = strconv.Atoi(val)
			case "templ":
				state.LowestTemperature, _ = strconv.Atoi(val)
			case "temph":
				state.HighestTemperature, _ = strconv.Atoi(val)
			case "tof":
				state.TOF, _ = strconv.Atoi(val)
			case "h":
				state.Height, _ = strconv.Atoi(val)
			case "bat":
				state.Battery, _ = strconv.Atoi(val)
			case "baro":
				state.Barometer, _ = strconv.ParseFloat(val, 64)
			case "time":
				state.MotorsOnTime, _ = strconv.Atoi(val)
			case "agx":
				state.AGX, _ = strconv.ParseFloat(val, 64)
			case "agy":
				state.AGY, _ = strconv.ParseFloat(val, 64)
			case "agz":
				state.AGZ, _ = strconv.ParseFloat(val, 64)
			default:
				log.Warn("Unknown status key ", keyVal)
			}
		}
	}

	return state, nil
}
