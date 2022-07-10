package domain

type JobSuggestionsGraph interface {
	GetAll() ([]*Job, error)
	DeleteAll()
	Register(job *Job) (int64, error)
}
