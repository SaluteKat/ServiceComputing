package User

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"bufio"
)

type User struct {
	Username           string `json:"Username"`
	Password           string `json:"Password"`
	Email              string `json:"Email"`
	SponsorMeeting     []string
	ParticipantMeeting []string
}

func (u User) GetName() string {
	return u.Username;
}

func (u *User) SetName(name string) {
	u.Username = name;
}

func (u User) GetPassword() string {
	return u.Password;
}

func (u *User) SetPassword(password string) {
	u.Password = password;
}

func (u User) GetEmail() string {
	return u.Email;
}

func (u *User) SetEmail(email string) {
	u.Email = email;
}

//register an  user with name, password, email
func RegisterAnUser(user *User) {
	flog, err := os.OpenFile("data/input_output.log", os.O_APPEND|os.O_WRONLY, 0666)
	defer flog.Close()
	check(err)
	logger := log.New(flog, "", log.LstdFlags)
	logger.Printf("agenda register -u %s -p %s -e %s", user.Username, user.Password, user.Email)
	
	var userinfo User
	userinfo.Username = user.Username
	userinfo.Password = user.Password
	userinfo.Email = user.Email
	userinfo.SponsorMeeting = make([]string, 0, 5)
	userinfo.ParticipantMeeting = make([]string, 0, 5)
	//userinfo := &User.User{user.Username, user.Password, user.Email, make([]string, 0, 5), make([]string, 0, 5)}
	file, err := os.OpenFile("data/User.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	num := detectUser(user.Username)
	AllUserInfo := GetAllUserInfo()
	if num == 1 {
		AllUserInfo = append(AllUserInfo, userinfo)
		encoder := json.NewEncoder(file)
		encoder.Encode(AllUserInfo)
		os.Stdout.WriteString("register succeed!\n")
		logger.Print("register succeed!\n")
	} else {
		os.Stdout.WriteString("The userName have been registered\n")
		logger.Print("The userName have been registered\n")
	}
	file.Close();
}

//whether the user exits
func detectUser(name string) int{
	file, err := os.OpenFile("data/User.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More(){
		var users User
		decoder.Decode(&users)
		if users.Username == name{
			file.Close()
			return 0
		}
	}
	return 1

}

//Get all user infomation
func GetAllUserInfo() []User {
	var users []User
	byteIn, err := ioutil.ReadFile("data/User.json")
	check(err)
	jsonStr := string(byteIn)
	json.Unmarshal([]byte(jsonStr), &users)
	return users
}

func check(r error) {
	if r != nil {
		log.Fatal(r)
	}
}


func LogIn(user *User){
	users := GetAllUserInfo();
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
	var length int = len(users)
	var flag = false
	for ii := 0; ii < length; ii++ {
		if users[ii].Username == user.Username {
			if users[ii].Password == user.Password{
				fout, _ := os.Create("data/current.txt")
				defer fout.Close()
				fout.WriteString(user.Username)
				os.Stdout.WriteString("You have logged in successfully.\n")
				logger.Print("You have logged in successfully.\n")
			} else {
				os.Stdout.WriteString("Password is incorrect!\n")
				logger.Print("Password is incorrect!\n")
			}
			flag = true
		}
	}
	if  flag == false {
		os.Stdout.WriteString("Username is not correct.\n")
		logger.Print("Username is not correct.\n")
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


