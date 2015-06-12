package session

const Endpoint = "/searchSessions"

type FsSessions struct {
	Clips    []FsClip
	Sessions []FsSession
}

type FsClip struct {
	Markers []FsMarker
}

type FsMarker struct {
	EventPageOffsetMs    int
	EventSessionOffsetMs int
	IndivId              string
	PageId               string
	PageStart            int
	SessionId            string
	UserId               string
}

type FsSession struct {
	Created           int
	DurationInSeconds int
	NumEvents         int
	NumPages          int
	UserSessionId     string
}
