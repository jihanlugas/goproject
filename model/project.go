package model

import (
	"github.com/jihanlugas/goproject.git/config"
	"time"
)

type Project struct {
	ID       	int    `json:"id"`
	Name     	string `json:"name"`
	Location	string `json:"location"`
	Description	string `json:"description"`
	StartAt		time.Time `json:"startAt"`
	EndAt		time.Time `json:"endAt"`
	ProjectTask []ProjectTask `json:"projectTask"`
}

func GetProjects(start, count int) ([]Project, error) {
	db := config.DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT id, name, location, description, start_at, end_at FROM projects LIMIT ? OFFSET ?", count, start)

	if err != nil {
		return nil, err
	}

	projects := []Project{}

	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Location, &p.Description, &p.StartAt, &p.EndAt); err != nil {
			return nil, err
		}

		projecttasks, err := GetProjectTasks(p.ID)

		if err != nil {
			return nil, err
		}

		p.ProjectTask = projecttasks

		projects = append(projects, p)
	}

	return projects, nil
}

func (p *Project) GetProject() error {
	db := config.DbConn()
	defer db.Close()

	err := db.QueryRow("SELECT name, location, description, start_at, end_at FROM projects where id = ?",
		p.ID).Scan(&p.Name, &p.Location, &p.Description, &p.StartAt, &p.EndAt)

	if err != nil {
		return err
	}

	projecttasks, err := GetProjectTasks(p.ID)

	if err != nil {
		return err
	}

	p.ProjectTask = projecttasks

	return nil
}

func (p *Project) CreateProject() error {
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("INSERT INTO projects(name, location, description, start_at, end_at) VALUES(?, ?, ?, ?, ?)", p.Name, p.Location, p.Description, p.StartAt, p.EndAt)

	if err != nil {
		return err
	}

	return nil
}

func (p *Project) UpdateProject() error {
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("UPDATE projects SET name=?, location=?, description=?, start_at=?, end_at=? WHERE id=?",
		p.Name, p.Location, p.Description, p.StartAt, p.EndAt, p.ID)

	return err
}

func (p *Project) DeleteProject() error {
	db := config.DbConn()
	defer db.Close()

	if err := DeleteProjectTask(p.ID); err != nil {
		return err
	}

	_, err := db.Exec("DELETE FROM projects WHERE id = ? ", p.ID)

	return err
}



