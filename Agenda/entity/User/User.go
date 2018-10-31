package User

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Username           string
	Password           string
	Email              string
	SponsorMeeting     []string
	ParticipantMeeting []string
}

//register an  user with name, password, email
func RegisterAnUser(user *User) {
	AllUserInfo := GetAllUserInfo()
	flog, err := os.OpenFile("data/input_output.log", os.O_APPEND|os.O_WRONLY, 0600)
	defer flog.Close()
	check(err)
	logger := log.New(flog, "", log.LstdFlags)
	logger.Printf("agenda register -u %s -p %s -e %s", user.Username, user.Password, user.Email)

	if _, ok := AllUserInfo[user.Username]; !ok {
		AllUserInfo[user.Username] = *user
		os.Stdout.WriteString("register succeed!\n")
		logger.Print("register succeed!\n")
	} else {
		os.Stdout.WriteString("The userName have been registered\n")
		logger.Print("The userName have been registered\n")
	}
}

//Get all user infomation
func GetAllUserInfo() map[string]User {

	byteIn, err := ioutil.ReadFile("data/User.json")
	check(err)
	var allUserInfo map[string]User
	json.Unmarshal(byteIn, &allUserInfo)
	return allUserInfo
}

func check(r error) {
	if r != nil {
		log.Fatal(r)
	}
}


func LogIn(user *User){
	AllUserInformation := GetAllUserInfo()

	flog, err := os.OpenFile("data/input_output.log", os.O_APPEND|os.O_WRONLY, 0600)
	defer flog.Close()

	if(err != nil){
		log.Fatal(err)
	}

	logger := log.New(flog, "", log.LstdFlags)
	logger.Printf("agenda login -u %s -p %s", user.Username, user.Password)

	//get current username
	fin, err1 := os.Open("data/current.txt")
	if(err1 != nil){
		log.Fatal(err1)
	}
	defer fin.Close()
	reader := bufio.NewReader(fin)
	name, _ := reader.ReadString('\n')

	if name != "" {
		os.Stdout.WriteString("You have log in already.\n")
		logger.Print("You have log in already.\n")
		return
	} 

	_, ok := AllUserInformation[user.Username]
	if  !ok {
		os.Stdout.WriteString("Username is not correct.\n")
		logger.Print("Username is not correct.\n")
	} else {
		correctPassword := AllUserInformation[user.Username].Password
		if correctPassword == user.Password {
			fout, _ := os.Create("data/current.txt")
			defer fout.Close()
			fout.WriteString(user.Username)
			os.Stdout.WriteString("You have logged in successfully.\n")
			logger.Print("You have logged in successfully.\n")
		} else {
			os.Stdout.WriteString("Password is incorrect!\n")
			logger.Print("Password is incorrect!\n")
		}
	}
}



func LogOut() {
    flog, err := os.OpenFile("data/input_output.log", os.O_APPEND|os.O_WRONLY, 0600)
    defer flog.Close()
    if(err != nil){
		log.Fatal(err)
	}

    logger := log.New(flog, "", log.LstdFlags)
    logger.Printf("Agenda logout")

    fin, err1 := os.Open("data/current.txt")
    if(err1 != nil){
		log.Fatal(err1)
	}
    defer fin.Close()
    reader := bufio.NewReader(fin)
    name, _ := reader.ReadString('\n')

    if name == "" {
        os.Stdout.WriteString("You have not logged in.\n")
        logger.Print("You have not logged in.\n")
    } else {
        fout, _ := os.Create("data/current.txt")
        defer fout.Close()
        fout.WriteString("")
        os.Stdout.WriteString("You have logged out successfully.\n")
        logger.Print("You have logged out successfully.\n")
    }
}
