package goutils

import (
	"fmt"
	"log"
	"strings"
)

//return true if err is not nil and false if err is nil
func IsError(err error) bool {
	if err != nil {
		log.Println("CheckErr log")
		log.Println(err)
		return true
	}
	return false
}

//panic if error
func PanicIfError(err error) {
	if err != nil {
		log.Println("PanicIfError log")
		log.Println(err)
		panic(err)
	}
}

//convert a string to string map to a query string
func StringsMapToQueryString(queryParamsMap map[string]string) string {
	var queryParamsString = ""
	if queryParamsMap != nil {
		queryParamsString = "?"
		for param, value := range queryParamsMap {
			queryParamsString = queryParamsString + fmt.Sprintf("%s=%s&", param, value)
		}
		queryParamsString = strings.TrimSuffix(queryParamsString, "&")
	}
	return queryParamsString
}
