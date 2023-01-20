package department

import (
	model "restApi/internal/app/model/departments"
	//."restApi/internal/app/store"
)

type DepartmentRepositor struct {
	store *Store
}

func (r *DepartmentRepositor) NrDepartmentAddNewPosition(d *model.NrDepartment) error {

	return r.store.db.QueryRow(
		"INSERT INTO nr_department (date, time, user_name, seccond_name) VALUES ($1, $2,$3,$4) RETURNING id",
		d.Date,
		d.Time,
		d.Day0,
		d.TypeRequest1,
		d.ProjectManager,
		d.TypeRequest2,
		d.TypeMessage,
		d.AlertData,
		d.ConsumerInformation,
		d.ReactionDescription,
		d.LPname,
		d.DRUlp,
		d.DescriptionRandMovment,
		d.CaseStatus,
		d.FirstMessageData,
		d.FinalAnswer,
		d.NumberIzv,
		d.Valid,
	).Scan(&d.ID)

}
