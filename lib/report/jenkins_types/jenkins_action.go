package jenkins_types

// Action should match up exactly with the Jenkins "action" schema
type Action struct {
	Class           string   `json:"_class"`
	TriggeredBuilds []Build  `json:"triggeredBuilds"`
	Causes          []Causes `json:"causes"`
}

// Causes should match up exactly with the Jenkins "causes" schema
type Causes struct {
	UpstreamBuild   int    `json:"upstreamBuild"`
	UpstreamProject string `json:"upstreamProject"`
	UpstreamURL     string `json:"upstreamURL"`
}
