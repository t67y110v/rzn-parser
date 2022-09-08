package model

type NrDepartment struct {
	ID                     int    `json:"id"`
	Date                   string `json:"date"`
	Time                   string `json:"time"`
	Day0                   string `json:"day_0"`
	TypeRequest1           string `json:"typerequest1"`
	ProjectManager         string `json:"ProjectManager"`
	TypeRequest2           string `json:"typerequest2"`
	TypeMessage            string `json:"typemessage"`
	AlertData              string `json:"alertdata"`
	ConsumerInformation    string `json:"consumerinformation"`
	ReactionDescription    string `json:"reactiondescription"`
	LPname                 string `json:"lpname"`
	DRUlp                  string `json:"drulp"`
	DescriptionRandMovment string `json:"movmentdescription"`
	CaseStatus             string `json:"casestatus"`
	FirstMessageData       string `json:"firstdatamessage"`
	FinalAnswer            string `json:"finalanswer"`
	NumberIzv              string `json:"numberizv"`
	Valid                  string `json:"valid"`
}
