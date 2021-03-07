package utils

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const TimeLayout = "2006-01-02"

//func ParseQueryTimeToFrom(r *http.Request) (fromTime time.Time, toTime time.Time, err error) {
//	queryUrl := r.URL.Query()
//
//	from := strings.Join(queryUrl["from"], "")
//	to := strings.Join(queryUrl["to"], "")
//	fromTime, err = ParseTime(from)
//	if err != nil {
//		return fromTime, toTime, errors.New("can not parse from date")
//	}
//	toTime, err = ParseTime(to)
//	if err != nil {
//		return fromTime, toTime, errors.New("can not parse to date")
//	}
//	return fromTime, toTime, err
//}
//
//func contains(s []string, str string) bool {
//	for _, v := range s {
//		if v == str {
//			return true
//		}
//	}
//
//	return false
//}

func ParseUrlQueryPage(r *http.Request) (ClientPageInt int, err error) {
	queryUrl := r.URL.Query()
	ClientPageStr := strings.Join(queryUrl["page"], "")
	if ClientPageStr == "" {
		return 0, errors.New("can't find query page, please check number of page")
	}
	ClientPageInt, err = strconv.Atoi(ClientPageStr)
	if err != nil {
		return 0, err
	}
	return ClientPageInt, nil
}

func ParseUrlQueryID(r *http.Request) (ClientIdInt int, err error) {
	queryUrl := r.URL.Query()
	ClientIdStr := strings.Join(queryUrl["id"], "")
	if ClientIdStr == "" {
		return 0, errors.New("can't find query id, please check number of id")
	}
	ClientIdInt, err = strconv.Atoi(ClientIdStr)
	if err != nil {
		return 0, err
	}
	return ClientIdInt, nil
}

//func ParseTime(Time string) (ParsedTime time.Time, err error) {
//	ParsedTime, err = time.Parse(TimeLayout, Time)
//	return ParsedTime, err
//}
