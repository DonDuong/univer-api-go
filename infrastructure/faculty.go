package infrastructure

import (
	"database/sql"
	_ "embed"
	"university-api/domain/model"
	"university-api/domain/repository"
)

type ApiFaculty struct {
	database repository.DatabaseInterface
}

func NewFacultyRepository(
	database repository.DatabaseInterface,
) repository.ApiFacultyInterface {
	return &ApiFaculty{
		database: database,
	}
}

//go:embed sql/faculty/get.sql
var sqlGet string

func (api *ApiFaculty) GetFaculty(tx repository.TransactionInterface, faculty_cd string) (*model.Faculty, error) {
	var faculty = &model.Faculty{}

	err := tx.Get().(*sql.Tx).QueryRow(sqlGet, faculty_cd).Scan(
		&faculty.Facuty_cd,
		&faculty.Faculty_name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return faculty, nil
}
