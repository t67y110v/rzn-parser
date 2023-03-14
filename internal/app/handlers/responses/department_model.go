package responses

type DepartmentRes struct {
	Departments           Department
	MonitoringSpecialist  bool `json:"monitoring_specialist"`
	MonitoringResponsible int  `json:"monitoring_responsible"`
}
