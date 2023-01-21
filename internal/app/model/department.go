package model

type Department struct {
	ClientDepartment            bool `json:"client_department"`
	EducationDepartment         bool `json:"education_department"`
	SourceTrackingDepartment    bool `json:"source_tracking_department"`
	PeriodicReportingDepartment bool `json:"periodic_reporting_department"`
	InternationalDepartment     bool `json:"international_department"`
	DocumentationDepartment     bool `json:"documentation_department"`
	NrDepartment                bool `json:"nr_department"`
	DbDepartment                bool `json:"db_department"`
}
