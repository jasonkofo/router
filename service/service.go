package service

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	config "github.com/jasonkofo/configservice/configservice"
	"github.com/jasonkofo/servicestarter/commonviews"
	"github.com/julienschmidt/httprouter"
)

var service *Service
var services []commonviews.ServiceConfig
var serv config.Service
var once sync.Once

// Service definition
type Service struct {
	Name              string             `json:"name"`
	InitialConfigPort int                `json:"initialConfigPort"`
	HTTPPort          int                `json:"httpPort"`
	EnableHTTPS       bool               `json:"enableHTTPS"`
	Router            *httprouter.Router `json:"router"`
	Services          []commonviews.ServiceConfig
}

func initService() {
	once.Do(func() {
		f, err := ioutil.ReadFile("config/router.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(f, &service)
		if err != nil {
			panic(err)
		}

		service.Router = httprouter.New()
	})
}
