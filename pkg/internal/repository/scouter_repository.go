package repository

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Scouter struct {
	Client *resty.Client
}

func NewScouter() (*Scouter, error) {
	//resp, err := client.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"username":"test","password":"test"}`).
	//	Post("https://jsonplaceholder.typicode.com/posts")
	//
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}

	//fmt.Println("Response Status Code:", resp.StatusCode())
	//fmt.Println("Response Body:", resp.String())
	return &Scouter{
		Client: resty.New(),
	}, nil
}

func (receiver Scouter) GetProfile() error {
	_, err := receiver.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"test","password":"test"}`).
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
