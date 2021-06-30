package model

import (
	"github.com/jihanlugas/goproject.git/config"
	"log"
)

type ProjectTask struct {
	ID       	int		`json:"id"`
	ProjectId	int		`json:"projectId"`
	Name     	string	`json:"name"`
	IsComplete 	int		`json:"isComplete"`
}

func GetProjectTasks(projectId int) ([]ProjectTask, error) {
	log.Println("Hit GetProjectTasks")
	db := config.DbConn()
	defer db.Close()
	rows, err := db.Query("SELECT id, project_id, name, is_complete FROM projecttasks WHERE project_id = ? ", projectId)

	if err != nil {
		log.Println("Hit err1")
		log.Println(err)
		return nil, err
	}

	projecttask := []ProjectTask{}

	for rows.Next() {
		var t ProjectTask
		if err := rows.Scan(&t.ID, &t.ProjectId, &t.Name, &t.IsComplete); err != nil {
			log.Println("Hit err2")
			return nil, err
		}

		projecttask = append(projecttask, t)
	}

	return projecttask, nil
}

func (t *ProjectTask) CreateTask() error {
	log.Println("Hit http://localhost:8010/task")
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("INSERT INTO projecttasks(project_id, name, is_complete) VALUES(?, ?, ?)", t.ProjectId, t.Name, t.IsComplete)

	if err != nil {
		return err
	}

	return nil
}

func (t *ProjectTask) DeleteTask() error {
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("DELETE FROM projecttasks WHERE id = ? ", t.ID)

	return err
}

func DeleteProjectTask(projectId int) error {
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("DELETE FROM projecttasks WHERE project_id = ? ", projectId)

	return err
}
