package model

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lib/pq"
	"net/http"
)

func (a *Advertisement) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Length(1, 200),validation.Required),
		validation.Field(&a.Description, validation.Length(0, 1000)),
		validation.Field(&a.Cost, validation.By(IsPositive),validation.Required),
		validation.Field(&a.ImgUrl, validation.By(ValidateImg)),
	)
}

func IsPositive(value interface{}) error {
	s, _ := value.(float64)
	if s < 0 {
		return errors.New("cost cant be negative")
	}
	return nil
}

func ValidateImg(value interface{}) error {

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	defer client.CloseIdleConnections()
	s, _ := value.(pq.StringArray)
	for _, UrlImg := range s {
		if UrlImg != "" {
			resp, err := client.Get(UrlImg)

			if err != nil {
				return err
			}
			if resp.Status != "200 OK" {
				errMessage := fmt.Sprintf("incorrect url for img, url: %s", UrlImg)
				resp.Body.Close()
				return errors.New(errMessage)
			}
			resp.Body.Close()

		}
	}

	return nil
}
