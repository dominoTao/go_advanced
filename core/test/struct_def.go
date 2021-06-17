package testing

type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
type JobInfo struct {
	Skills []string `json:"skills"`
}
type EEEdfe struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo	`json:"job_info"`
}
