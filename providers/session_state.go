package providers

import (
	"fmt"
	"time"
)

type SessionState struct {
	AccessToken  string
	ExpiresOn    time.Time
	RefreshToken string
	Email        string
	User         string
	IDToken      string
}

func (s *SessionState) IsExpired() bool {
	if !s.ExpiresOn.IsZero() && s.ExpiresOn.Before(time.Now()) {
		return true
	}
	return false
}

func (s *SessionState) String() string {
	o := fmt.Sprintf("Session{%s", s.accountInfo())
	if s.AccessToken != "" {
		o += " token:true"
	}
	if !s.ExpiresOn.IsZero() {
		o += fmt.Sprintf(" expires:%s", s.ExpiresOn)
	}
	if s.RefreshToken != "" {
		o += " refresh_token:true"
	}
	if s.IDToken != "" {
		o += " id_token:true"
	}
	return o + "}"
}

func (s *SessionState) accountInfo() string {
	return fmt.Sprintf("email:%s user:%s", s.Email, s.User)
}
