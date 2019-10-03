package jenkins_types

import (
	"encoding/json"
)

/*
 * _class: "jenkins.metrics.impl.TimeInQueueAction"
 */

type TimeInQueue struct {
	BlockedDurationMillis   int
	BlockedTimeMillis       int
	BuildableDurationMillis int
	BuildableTimeMillis     int
	BuildingDurationMillis  int
	ExecutingTimeMillis     int
	ExecutorUtilization     int
	SubTaskCount            int
	WaitingDurationMillis   int
	WaitingTimeMillis       int
}

func (b *Build) GetTimeInQueue() {
	// var project string

	for _, action := range b.Actions {
		if action.Class == "jenkins.metrics.impl.TimeInQueueAction" {
			actionBytes, _ := json.Marshal(action)
			json.Unmarshal(actionBytes, &b.TimeInQueue)
		}
	}

}
