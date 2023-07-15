package services

import (
	"fmt"
	"university-api/domain/model"
	"university-api/domain/repository"
	interfaces "university-api/interfaces/service"
)

type ApiFacultyService struct {
	DatabaseInterface repository.DatabaseInterface
}

func NewApiFacultyService(
	DatabaseInterface repository.DatabaseInterface,
) interfaces.ApiFacultyServiceInterface {
	return &ApiFacultyService{
		DatabaseInterface: DatabaseInterface,
	}
}

func (service *ApiFacultyService) GetFaculty(tx repository.TransactionInterface, faculty_cd string) (*model.Faculty, error) {
	fmt.Println(faculty_cd)
	faculty, err := service.DatabaseInterface.ApiFaculty().GetFaculty(tx, faculty_cd)
	if err != nil {
		fmt.Println(faculty)
		return nil, err
	}
	fmt.Println("faculty service: ", faculty)
	return faculty, err
}

func (service *ApiFacultyService) DatabaseRepository() repository.DatabaseInterface {
	return service.DatabaseInterface
}
