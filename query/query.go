package query

const (
	OP_IS          = "=="
	OP_IS_NOT      = "!="
	OP_STARTS_WITH = "><"
)

type FsCriterion struct {
	Field string
	Op    string
	Value interface{}
}

type FsQuery struct {
	EventCriteria   [][]FsCriterion
	TimeRange       []string
	UserTipCriteria []FsCriterion
}

func NewFsQuery() (fsq *FsQuery) {
	fsq = &FsQuery{}
	fsq.EventCriteria = [][]FsCriterion{}
	fsq.UserTipCriteria = []FsCriterion{}
	return
}

func (q *FsQuery) AddEventCriterion(eventType FsCriterion, value FsCriterion) {
	query := []FsCriterion{eventType, value}
	q.EventCriteria = append(q.EventCriteria, query)
}

func (q *FsQuery) AddUserCriterion(c FsCriterion) {
	q.UserTipCriteria = append(q.UserTipCriteria, c)
}
