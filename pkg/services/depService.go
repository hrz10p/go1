package services

import (
	"database/sql"
	"main/pkg/models"
)

type DepService struct {
	db *sql.DB
}

func NewDepService(db *sql.DB) *DepService {
	return &DepService{db: db}
}

func (s *DepService) GetAllDeps() ([]models.Department, error) {
	rows, err := s.db.Query("SELECT * FROM deps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deps []models.Department

	for rows.Next() {
		dep := models.Department{}

		err := rows.Scan(
			&dep.ID,
			&dep.Dep_name,
			&dep.Staff_quantity,
		)
		if err != nil {
			return nil, err
		}

		deps = append(deps, dep)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deps, nil
}

func (s *DepService) CreateDep(d models.Department) error {
	_, err := s.db.Exec("INSERT INTO deps (dep_name, staff_quantity) VALUES ($1, $2)", d.Dep_name, d.Staff_quantity)
	if err != nil {
		return err
	}

	return nil
}
