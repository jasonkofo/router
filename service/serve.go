package service

import (
	"net/http"
)

// ListenAndServe calls go's ListenAndServe method after ensuring that the
// service has been adequately setup
func (s *Service) ListenAndServe() {
	s.setup()
	if err := http.ListenAndServe(getPortString(), s.Router); err != nil {
		panic(err)
	}
}
