package responses

type CreateUser struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	SeccondName string `json:"seccond_name"`
}

type Error struct {
	Message string `json:"message"`
}

type UserUpdate struct {
	Role                  string     `json:"user_role"`
	Departments           Department `json:"departments"`
	MonitoringSpecialist  bool       `json:"monitoring_specialist"`
	MonitoringResponsible int        `json:"monitoring_responsible"`
}

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

type Result struct {
	Result bool `json:"result"`
}
