package sim

import (
	"sort"
	"time"
)

// Sim runs the simulation
func Sim(input string) (results Results, err error) {
	emptyResult := Results{
		Results: []Result{},
	}
	// Parse the input into a sim
	s, err := generateSimFromCSV(input)
	if err != nil {
		return emptyResult, err
	}

	return s.Run()
}

// Results is a set of results
type Results struct {
	Results []Result `json:"results"`
}

// Result encapulates a result
type Result struct {
	Rider            RiderID       `json:"rider"`
	RequestToPickUp  time.Duration `json:"requestToPickUp"`
	RequestToDropOff time.Duration `json:"requestToDropOff"`
	PickUpToDropOff  time.Duration `json:"pickUpToDropOff"`
}

// RiderID represents a rider
type RiderID string

// DriverID represents a driver
type DriverID string

// Event represents an event
// Request, Pickup, Drop-Off
type Event string

// Event types
const (
	EventRequest Event = "Request"
	EventPickUp  Event = "Pickup"
	EventDropOff Event = "Drop-Off"
)

type simEvent struct {
	When   time.Time
	Rider  RiderID
	Driver DriverID
	Event  Event
}

type sim struct {
	events        []simEvent
	riderRequests map[RiderID]time.Time
	riderPickUps  map[RiderID]time.Time
	riderDropOffs map[RiderID]time.Time
}

// newSim returns a new sim with an allocated slice of events.
func newSim() sim {
	return sim{
		events:        []simEvent{},
		riderRequests: map[RiderID]time.Time{},
		riderPickUps:  map[RiderID]time.Time{},
		riderDropOffs: map[RiderID]time.Time{},
	}
}

// Run simulation
func (s *sim) Run() (Results, error) {
	r := Results{
		Results: []Result{},
	}

	// Sort the events before processing them
	s.SortEvents()

	resultsByRider := map[RiderID]Result{}

	for _, e := range s.events {
		switch e.Event {
		case EventRequest:
			resultsByRider[e.Rider] = Result{Rider: e.Rider}
		case EventPickUp:
			result := resultsByRider[e.Rider]
			duration := s.riderPickUps[e.Rider].Sub(s.riderRequests[e.Rider])
			result.RequestToPickUp = duration
			resultsByRider[e.Rider] = result
		case EventDropOff:
			result := resultsByRider[e.Rider]
			duration := s.riderDropOffs[e.Rider].Sub(s.riderPickUps[e.Rider])
			result.PickUpToDropOff = duration
			result.RequestToDropOff = result.RequestToPickUp + result.PickUpToDropOff
			resultsByRider[e.Rider] = result
		}
	}

	for _, rs := range resultsByRider {
		r.Results = append(r.Results, rs)
	}

	return r, nil
}

func (s *sim) SortEvents() {
	sort.SliceStable(s.events, func(i, j int) bool {
		return s.events[i].When.Before(s.events[i].When)
	})
}
