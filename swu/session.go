package swu

type Session struct {
	sessionId string
}

func NewSession(stopId string) *Session {
	return &Session{sessionId: getSessionID(stopId)}
}

func (s *Session) GetDepartures() []ItdDeparture {
	return getDepartures(s.sessionId)
}
