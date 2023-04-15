package requests

type CreateUser struct {
	Email                 string     `json:"email"`
	Role                  string     `json:"user_role"`
	Password              string     `json:"password"`
	Name                  string     `json:"name"`
	SeccondName           string     `json:"seccond_name"`
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
	CMKDepartment               bool `json:"cmk_department"`
	SalesDepartment             bool `json:"sales_department"`
}

type EmailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Email                 string     `json:"email"`
	Name                  string     `json:"name"`
	SeccondName           string     `json:"seccond_name"`
	Role                  string     `json:"user_role"`
	Departments           Department `json:"departments"`
	MonitoringSpecialist  bool       `json:"monitoring_specialist"`
	MonitoringResponsible int        `json:"monitoring_responsible"`
}

type Email struct {
	Email string `json:"email"`
}
