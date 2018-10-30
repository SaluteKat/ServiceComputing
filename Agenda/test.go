package main

import (
	//"bufio"
	//	"encoding/json"
	"fmt"
	//"log"
	//	"io/ioutil"
	//"os"
	"time"
)

type User struct {
	Username string
	Password string
	Email    string
}


func DateToString(date time.Time) string {
	return date.Format("2016-03-21 13:00")
}
func StringToDate(date string) (time.Time, error) {
	the_time, err := time.Parse("2016-03-21 13:00", date)
	return the_time, err
}

func main() {
	var str string
	str = "2017-10-30 10:00"
	t, _ := StringToDate(str)
	newStr := DateToString(t)
	fmt.Println(newStr)

}
