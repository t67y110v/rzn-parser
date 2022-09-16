package model

type Department struct {
	EducationDepartment         bool `json:"educationDepartment"`
	SourceTrackingDepartment    bool `json:"sourceTrackingDepartment"`
	PeriodicReportingDepartment bool `json:"periodicReportingDepartment"`
	InternationalDepartment     bool `json:"internationalDepartment"`
	DocumentationDepartment     bool `json:"documentationDepartment"`
	NrDepartment                bool `json:"nrDepartment"`
	DbDepartment                bool `json:"dbDepartment"`
}
