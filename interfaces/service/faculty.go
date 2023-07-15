package interfaces

import (
	"university-api/domain/model"
	"university-api/domain/repository"
)

type ApiFacultyServiceInterface interface {
	GetFaculty(tx repository.TransactionInterface, faculty_cd string) (*model.Faculty, error)
	DatabaseRepository() repository.DatabaseInterface
}
