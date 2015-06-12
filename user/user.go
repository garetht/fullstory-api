package main

const fsUserEndpoint = "/searchUsers"

type FsUsers struct {
	MatchingUsers int
	TotalUsers    int
	Users         []FsUser
}

type FsUser struct {
	AppKey               string
	AvgSessionSec        int
	Created              int
	DisplayName          string
	Email                string
	IndivId              string
	LastBrowser          string
	LastCity             string
	LastDevice           string
	LastOperatingSystem  string
	LastPage             string
	LastSessionNumEvents int
	LastSessionNumPages  int
	LastSessionSec       int
	LastSessionTime      int
	Meta                 string
	NumEvents            int
	NumSessions          int
	PageLatLong          string
	TotalSec             int
}
