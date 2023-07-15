package repository

import "university-api/domain/model"

type ApiFacultyInterface interface {
	GetFaculty(tx TransactionInterface, faculty_cd string) (*model.Faculty, error)
}
