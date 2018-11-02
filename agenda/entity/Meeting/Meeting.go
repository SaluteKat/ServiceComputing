package Meeting

import (
	"agenda/entity/User"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Meeting struct {
	Title        string
	Sponsor      string
	Participants []string
	StartTime    time.Time
	EndTime      time.Time
	Id           string
}

//time manage function
func DateToString(date time.Time) string {
	return date.Format("2006-01-02/15:04")
}
func StringToDate(date string) (time.Time, error) {
	the_time, err := time.Parse("2006-01-02/15:04", date)
	return the_time, err
}
func SmallDate(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}
func LargeDate(date1, date2 time.Time) bool {
	return date1.After(date2) || date1.Equal(date2)
}
func TimeContact(newDateS, newDateE, oriDateS, oriDateE time.Time, userName string, meetingId string) bool {
	if LargeDate(newDateS, newDateE) {
		fmt.Println("Invaild time: start time can't be greater than end time")
		return true
	}
	if SmallDate(newDateE, oriDateS) || LargeDate(newDateS, oriDateE) {   
		return false
	}
	fmt.Println("time contract with " + userName + "' meeting with id:" + meetingId)
	return true
}

func CreateAMeeting(meeting *Meeting) {
	bytes,_ := ioutil.ReadFile("data/current.txt")
	currentName := string(bytes)
	if currentName == "" {
		fmt.Println("You haven't logged in")
		return
	}
	var allMeetings map[string]Meeting = GetAllMeetingInfo()
	meeting.Id = strconv.Itoa(len(allMeetings)) //initial id is 0
	meeting.Sponsor = currentName
	meeting.Participants = append(meeting.Participants, currentName)
	allMeetings[meeting.Id] = *meeting

	var allUser []User.User = GetAllUserInfo()
	//check all participanter exist, and time contract
	for _, pName := range meeting.Participants {
		//check if the user exist
		var userind int = GetNameID(pName)
		if userind < 0 {
			fmt.Println("No such user:" + pName + "!")
			return
		}
		//check if the user's old meeting is contract with the new one
		for _, meetingId := range allUser[userind].ParticipantMeeting {
			if TimeContact(meeting.StartTime, meeting.EndTime, allMeetings[meetingId].StartTime, allMeetings[meetingId].EndTime, pName, meetingId) {
				return
			}
		}
		allUser[userind].ParticipantMeeting = append(allUser[userind].ParticipantMeeting, meeting.Id)
	}
	var nowindex int = GetNameID(currentName)
	allUser[nowindex].SponsorMeeting = append(allUser[nowindex].SponsorMeeting, meeting.Id)

	fout, _ := os.Create("data/Meeting.json")
	defer fout.Close()
	b, _ := json.Marshal(allMeetings)
	fout.Write(b)
	foutuser, _ := os.Create("data/User.json")
	defer foutuser.Close()
	buser, _ := json.Marshal(allUser)
	foutuser.Write(buser)
}

func GetNameID(name string) int{
	var allUser []User.User = GetAllUserInfo()
	num := len(allUser)
	for a := 0; a < num; a++ {
		if allUser[a].Username == name{
			return a
		}
	}
	return -1
}

//load all meeting infomation
func GetAllMeetingInfo() map[string]Meeting {
	byteIn, err := ioutil.ReadFile("data/Meeting.json")
	if err != nil {
		log.Fatal(err)
	}
	var allMeetingInfo map[string]Meeting
	json.Unmarshal(byteIn, &allMeetingInfo)
	if allMeetingInfo == nil {
		allMeetingInfo = make(map[string]Meeting)
	}
	return allMeetingInfo
}

func check(r error) {
	if r != nil {
		log.Fatal(r)
	}
}

//load all user infomation to User.AllUserInfo
func GetAllUserInfo() []User.User {
	var users []User.User
	byteIn, err := ioutil.ReadFile("data/User.json")
	check(err)
	jsonStr := string(byteIn)
	json.Unmarshal([]byte(jsonStr), &users)
	return users
}

//delete a meeting by title
func DeleteMeeting(title string){
	bytes,_ := ioutil.ReadFile("data/current.txt")
	curuser := string(bytes)
	var allMeetings map[string]Meeting = GetAllMeetingInfo()
	var allUser []User.User = GetAllUserInfo()

	var meetingleng int = len(allMeetings)
	var flag bool = false
	var value Meeting
	var meetingindex string
	for jj := 0; jj < meetingleng; jj++ {
		strindex := strconv.Itoa(jj)
		if allMeetings[strindex].Title == title {
			meetingindex = strindex
			flag = true
			value = allMeetings[strindex]
		}
	}
	if !flag {
		fmt.Println("Meeting does not exist!")
		return
	}
	if (value.Sponsor != curuser){
		fmt.Println("Sorry, you are not the sponser of this meeting")
		return
	}
	for _, pName := range value.Participants {
		//check if the user exist
		var userind int = GetNameID(pName)
		if userind < 0 {
			fmt.Println("No such user:" + pName + "!")
			return
		}
		//check if the user's old meeting is contract with the new one
		var length int = len(allUser[userind].ParticipantMeeting)
		for ii := 0; ii < length; ii++ {
			if allUser[userind].ParticipantMeeting[ii] == value.Id {
				allUser[userind].ParticipantMeeting = append(allUser[userind].ParticipantMeeting[:ii], allUser[userind].ParticipantMeeting[ii+1:]...)
			}
		}
	}
	var nowindex int = GetNameID(value.Sponsor)
	var leng int = len(allUser[nowindex].SponsorMeeting)
	for aa := 0; aa < leng; aa++ {
		if allUser[nowindex].SponsorMeeting[aa] == value.Id {
			allUser[nowindex].SponsorMeeting = append(allUser[nowindex].SponsorMeeting[:aa], allUser[nowindex].SponsorMeeting[aa+1:]...)
		}
	}

	delete(allMeetings, meetingindex)
	fmt.Println("Delete meeting successfully")

	fout, _ := os.Create("data/Meeting.json")
	defer fout.Close()
	b, _ := json.Marshal(allMeetings)
	fout.Write(b)
	foutuser, _ := os.Create("data/User.json")
	defer foutuser.Close()
	buser, _ := json.Marshal(allUser)
	foutuser.Write(buser)
}

