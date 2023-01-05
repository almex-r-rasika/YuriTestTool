package data

import (
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB
var isConnect = false

type LoginUser struct {
  UserId   string `gorm:"unique"`
  Password   string
  LoginTime string
  LogoutTime string
}

type Messages struct {
  MessageId string
  SendTime   string
  SendUserId   string
  Address string
  Subject string
  Line1 string
  Line2 string
  Line3 string
  Line4 string
  Line5 string
  Line6 string
  Line7 string
  Line8 string
  Line9 string
  Line10 string
}

type SendMessage struct {
  MessageId string
  SendTime   string
  SendUserId   string
  Address string
  Subject string
  Line1 string
  Line2 string
  Line3 string
  Line4 string
  Line5 string
  Line6 string
  Line7 string
  Line8 string
  Line9 string
  Line10 string
  PostTime time.Time
  Result string
}

var loginUser []LoginUser
var messages []Messages

/* function for make a connection to the database
    @param --> null
    @param value --> null
    description --> if connection has established before, ignore the making new connection
	@return --> null
*/
func connectToDatabase(){
	if !isConnect{
		time.Sleep(time.Duration(10000 * time.Millisecond))
		db = makeDbConnection()
	}
	isConnect = true
}

/* Save message list to the database
    @param --> null
    @param value --> null
    description --> read message.sql file and execute it
    @return --> null
*/
func SaveMessage() {

	connectToDatabase()
	absPath, _ := filepath.Abs("../main-service/data/message.sql")
	executeSqlFile(db, absPath)
	Log.Info("saved messages to the database")
}

/* Save login users to the database
    @param --> null
    @param value --> null
    description --> read loginuser.sql file and execute it
    @return --> null
*/
func SaveLoginUser() {

	connectToDatabase()
	absPath, _ := filepath.Abs("../main-service/data/loginuser.sql")
	executeSqlFile(db, absPath)
	Log.Info("saved users to the database")
}

/* function for get user object from login_users table
    @param --> getUserFor string
    @param value --> login/logout
    description --> if getUserFor parameter = login, get login user objects for auto login function
                   if getUserFor parameter = logout, get login user objects for auto logout function
	@return --> login user objects array
*/
func GetUsers(getUserFor string) ([]LoginUser){

	if(strings.Compare(getUserFor, "login") == 1){
		db.Order("login_time").Find(&loginUser)
		Log.Info("get login user objects from the database")	
	}else if(strings.Compare(getUserFor, "logout") == 1){
		db.Order("logout_time").Find(&loginUser)
		Log.Info("get logout user objects from the database")
	} else{
		db.Order("login_time").Find(&loginUser)
		Log.Info("get login user objects from the database")
	}
	return loginUser
}

/* function for get message object from messages table
    @param --> null
    @param value --> null
    description --> get all message objects from message table
    @return --> message objects array
*/
func GetMessages() ([]Messages){

	db.Order("send_time").Find(&messages)	
	Log.Info("get message objects from the database")
	return messages
}









