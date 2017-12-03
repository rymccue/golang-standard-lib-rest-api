package repositories

import (
	"database/sql"
	"github.com/rymccue/golang-standard-lib-rest-api/models"
)

func CreateJob(db *sql.DB, title, description string, userID int) (int, error) {
	const query = `
		insert into jobs (
			title,
			description,
			user_id
		) values (
			$1,
			$2,
			$3
		) returning id
	`
	var id int
	err := db.QueryRow(query, title, description, userID).Scan(&id)
	return id, err
}

func UpdateJob(db *sql.DB, jobID int, title, description string) error {
	const query = `
		update jobs set
			title = $1,
			description = $2
		where id = $3
	`
	_, err := db.Exec(query, title, description, jobID)
	return err
}

func DeleteJob(db *sql.DB, id int) error {
	const query = `delete from jobs where id = $1`
	_, err := db.Exec(query, id)
	return err
}

func GetJobByID(db *sql.DB, id int) (*models.Job, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			jobs
		where
			id = $1
	`
	var job models.Job
	err := db.QueryRow(query, id).Scan(&job.ID, &job.Title, &job.Description, &job.UserID)
	return &job, err
}

func GetJobs(db *sql.DB, page, resultsPerPage int) ([]*models.Job, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			jobs
		limit $1 offset $2
	`
	jobs := make([]*models.Job, 0)
	offset := (page - 1) * resultsPerPage

	rows, err := db.Query(query, resultsPerPage, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var job models.Job
		err = rows.Scan(&job.ID, &job.Title, &job.Description, &job.UserID)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, err
}
