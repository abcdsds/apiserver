package repository

import (
	"github.com/go-resty/resty/v2"
)

const (
	xLogPath          = "/xlog"
	activeServicePath = "/activeService"
	profilePath       = "/profile-data"
)

type Scouter struct {
	Client *resty.Client
}

func NewScouter() (*Scouter, error) {
	return &Scouter{
		Client: resty.New(),
	}, nil
}

func (receiver Scouter) GetProfile() (*Profile, error) {
	_, err := receiver.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"test","password":"test"}`).
		Post(profilePath)
	if err != nil {
		return nil, err
	}

	// Logic

	return &Profile{}, nil
}

func (receiver Scouter) GetActiveService() (*ActiveService, error) {
	_, err := receiver.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"test","password":"test"}`).
		Post(activeServicePath)
	if err != nil {
		return nil, err
	}

	return &ActiveService{}, nil
}
