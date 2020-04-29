package service

import (
	"fmt"

	"github.com/jasonkofo/router/repository"
	"github.com/jasonkofo/servicestarter/commonviews"
	ss "github.com/jasonkofo/servicestarter/servicestarter"
)

// getPortString returns the string for the port of the HTTP server
func getPortString() string {
	if service != nil {
		return fmt.Sprintf(":%v", service.HTTPPort)
	}
	return ""
}

// New returns a new instance of the Service, or an error
func New() (*Service, error) {
	services := []commonviews.ServiceConfig{}
	initService()

	// Get service registration from the configservice.
	url := fmt.Sprintf("http://localhost:%v/services", service.InitialConfigPort)
	err := ss.GetJSON(&services, url)
	if err != nil {
		return nil, err
	}
	service.Services = services
	return service, nil
}

// setup sets the API handlers up for use
func (s *Service) setup() {
	ss.Handle(s.Router, "GET", "/ping", repository.Ping)
}
