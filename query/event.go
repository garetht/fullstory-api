package query

const eventType = "EventType"

/*
  When we search for users who have clicked on the
  CSS selector .test, this is in fact two criteria:
  (event type, equals, click) and (css selector, equals, .test).

  The event type criteria do not make sense on their own,
  so these structs are provided for convenience.
*/

var ClickEventCriterion = FsCriterion{
	Field: eventType,
	Op:    OpEquals,
	Value: "click",
}
var VisitedEventCriterion = FsCriterion{
	Field: eventType,
	Op:    OpEquals,
	Value: "navigate",
}
var ChangedEventCriterion = FsCriterion{
	Field: eventType,
	Op:    OpEquals,
	Value: "change",
}

const (
	ClickChangeCssSelector = "$EventTargetSelector"
	ClickChangeText        = "EventTargetText"
)

const (
	VisitedHasHost       = "PageUrlHost"
	VisitedHasPath       = "PageUrlPath"
	VisitedHasQueryParam = "PageUrlQuery"
)

const (
	EventSecondsFromPageLoad     = "EventPageOffset"
	EventSecondsFromSessionStart = "EventSessionOffset"
	EventPageDuration            = "PageDuration"
	EventPageActiveDuration      = "PageActiveDuration"
)
