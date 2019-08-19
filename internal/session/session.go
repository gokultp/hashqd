package session

const (
	tubeDefault = "default"
)

type Session struct {
	Tube string
}

func NewSession() *Session {
	return &Session{
		Tube: tubeDefault,
	}
}

func (s *Session) SetTube(tube string) {
	s.Tube = tube
}
