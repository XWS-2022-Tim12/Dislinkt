package persistence

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Register(session neo4j.Session, job *domain.Job) (int64, error) {
	var jobId int64
	session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"CREATE (job:JOB {userId: $userId, position: $position, description: $description, requirements: $requirements})"+
				"RETURN ID(job), job.userId",
			map[string]interface{}{
				"userId":       job.UserId,
				"position":     job.Position,
				"description":  job.Description,
				"requirements": job.Requirements})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			jobId = result.Record().Values[0].(int64)

		}
		return jobId, nil
	})
	return jobId, nil
}

func DeleteAll(session neo4j.Session) {
	session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var records, err = tx.Run(
			"MATCH (n) DETACH DELETE n", map[string]interface{}{})
		return records, err
	})
	return
}

func GetAll(session neo4j.Session) (jobs []*domain.Job, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (job:JOB) RETURN ID(job),job.userId, job.position, job.description, job.requirements", map[string]interface{}{})

		for records.Next() {
			job := domain.Job{
				Id:           records.Record().Values[0].(int64),
				UserId:       records.Record().Values[1].(string),
				Position:     records.Record().Values[2].(string),
				Description:  records.Record().Values[3].(string),
				Requirements: records.Record().Values[4].(string),
			}
			jobs = append(jobs, &job)
		}

		if err != nil {
			return nil, err
		}

		return jobs, nil
	})
	if err != nil {
		return nil, err
	}
	return jobs, nil

}
