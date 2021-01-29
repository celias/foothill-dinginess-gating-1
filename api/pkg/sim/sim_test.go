package sim_test

import (
	"encoding/json"
	"fmt"
	"simplesim/api/pkg/sim"
	"testing"
)

func TestSim(t *testing.T) {
	result, err := sim.Sim(input)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Results) < 1 {
		t.Fatal("No results")
	}

	output, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", output)
}

const input = `

  Time,Rider ID,Driver ID,Event
2017-07-24 07:53:41,R35,,Request
2017-07-24 07:55:26,R35,D39,Pickup
2017-07-24 07:56:59,R23,,Request
2017-07-24 07:58:58,R23,D54,Pickup
2017-07-24 08:11:01,R66,,Request
2017-07-24 08:21:49,R66,D39,Pickup
2017-07-24 08:32:58,R23,D54,Drop-Off
2017-07-24 08:36:36,R35,D39,Drop-Off
2017-07-24 08:37:12,R71,,Request
2017-07-24 08:49:50,R71,D54,Pickup
2017-07-24 08:51:20,R38,,Request
2017-07-24 08:51:43,R66,D39,Drop-Off
2017-07-24 09:04:55,R38,D39,Pickup
2017-07-24 09:11:20,R38,D39,Drop-Off
2017-07-24 09:24:34,R45,,Request
2017-07-24 09:31:50,R45,D54,Pickup
2017-07-24 09:36:20,R71,D54,Drop-Off
2017-07-24 09:39:40,R45,D54,Drop-Off
  `
