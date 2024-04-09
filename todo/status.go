package todo

type status string

const (
	statusPending   status = "pending"
	statusCompleted status = "completed"
	statusCanceled  status = "canceled"
	statusDeleted   status = "deleted"
)

func (s *status) isValid() bool {
	switch *s {
	case statusPending, statusCompleted, statusCanceled, statusDeleted:
		return true
	}

	return false
}
