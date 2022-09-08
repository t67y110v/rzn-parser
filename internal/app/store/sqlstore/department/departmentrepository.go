package department

import (
	"restApi/internal/app/model"
	//."restApi/internal/app/store"
)

type DepartmentRepositor struct {
	store *Store
}

func (r *DepartmentRepositor) NrDepartmentAddNewPosition(d *model.NrDepartment) error {

	return r.store.db.QueryRow(
		"INSERT INTO nrdepartment (date, time, userName, seccondName) VALUES ($1, $2,$3,$4) RETURNING id",
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