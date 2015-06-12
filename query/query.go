package query

/*
	Regarding Fullstory Queries

	There are two different scopes for queries, the
	User Scope and the Event Scope.


	USER SCOPED QUERIES
	====

	User Scoped queries match a single attribute for a user.
	Many of these are straightforward (e.g. user email equals
	so-and-so), but queries in other time ranges can be
	much more complex and custom (e.g. users created within
	the past day).

	Here is an example. To write the query for email@email.com,
	we need to send this query to Fullstory:

	{
	  "TimeRange": null,
	  "UserTipCriteria": [
	    {
	      "Field": "UserEmail",
	      "Op": "==",
	      "Value": [
	        "email@email.com"
	      ]
	    }
	  ],
	  "EventCriteria": []
	}

	To do this, make a new query and add a User Criterion:

	q := query.NewFsQuery()
	q.AddUserCriterion(query.FsCriterion{
		Field: query.UserEmail,
		Op:    query.OpEquals,
		Value: []string{"email@email.com"},
	})

	Event Scoped queries are even more complex. They require the
	specification of
		1. An event type
		2. A search criterion for that event
		3. An optional third criterion

	For example, here is a query for all users who
		1. Click
		2. On the CSS Selector .test
		3. 3 seconds after page load

		{
		  "TimeRange": null,
		  "UserTipCriteria": [],
		  "EventCriteria": [
		    [
		      {
		        "Field": "EventType",
		        "Op": "==",
		        "Value": "click"
		      },
		      {
		        "Field": "$EventTargetSelector",
		        "Op": "==",
		        "Value": [
		          ".test"
		        ]
		      },
		      {
		        "Field": "EventPageOffset",
		        "Op": "==",
		        "Value": 3000
		      }
		    ]
		  ]
		}

*/

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

func (q *FsQuery) AddEventCriteria(criteria ...FsCriterion) {
	q.EventCriteria = append(q.EventCriteria, criteria)
}

func (q *FsQuery) AddUserCriterion(c FsCriterion) {
	q.UserTipCriteria = append(q.UserTipCriteria, c)
}
