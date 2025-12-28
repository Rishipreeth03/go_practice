package storage

import "github.com/Rishipreeth03/Api_Handler/Api_Handler/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudentList() ([]types.Student, error)
}
