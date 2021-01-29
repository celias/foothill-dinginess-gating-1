package sim

import (
	"strings"
)

// CSV input indicies
const (
	csvIdxTime = iota
	csvIdxRiderID
	csvIdxDriverID
	csvIdxEvent
)

// generateSimFromCSV returns a new sim from the given input
func generateSimFromCSV(input string) (sim, error) {
	s := newSim()

	rawLines := strings.Split(strings.Trim(input, " \n"), "\n")
	for idx, row := range rawLines {
		e := simEvent{}
		values := strings.Split(row, ",")

		// HACK: Determine if we have a header by assuming a failed time
		// parse in the first row means it is a header.
		parsedTime, timeParseErr := parseTime(values[csvIdxTime])
		if idx == 0 && timeParseErr != nil {
			continue
		} else if timeParseErr != nil {
			return s, timeParseErr
		}

		e.When = parsedTime
		e.Driver = DriverID(values[csvIdxDriverID])
		e.Rider = RiderID(values[csvIdxRiderID])
		e.Event = Event(values[csvIdxEvent])

		switch e.Event {
		case EventRequest:
			s.riderRequests[e.Rider] = e.When
		case EventPickUp:
			s.riderPickUps[e.Rider] = e.When
		case EventDropOff:
			s.riderDropOffs[e.Rider] = e.When
		}

		s.events = append(s.events, e)
	}

	return s, nil
}
