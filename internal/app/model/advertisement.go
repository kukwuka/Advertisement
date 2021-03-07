package model

import (
	"github.com/lib/pq"
	"time"
)

type Advertisement struct {
	Date        time.Time      `json:"date"`
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Cost        float64        `json:"cost"`
	ImgUrl      pq.StringArray `json:"img_url"`
}
