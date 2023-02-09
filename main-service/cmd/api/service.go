package main

import (
	"service/data"
	"time"
)

/* function for initialize and save login user/message objects to corresponding tables
   @param --> null
   @param value --> null
   description --> save login users and messages to db
   @return --> null
*/
func initializeService() {
	data.MakeLogger()
	data.SaveLoginUser()
	data.SaveMessage()
}

/* function for run auto login/auto logout/auto message functions as go routines
   @param --> null
   @param value --> null
   description --> run auto login/auto logout/auto message functions as go routines
   @return --> null
*/
func mainService(){
	login_users := data.GetUsers("login")
	logout_users := data.GetUsers("logout")
	//messages := data.GetMessages()

	data.Wg.Add(3)
	go data.AutoLogin(login_users)
	time.Sleep(time.Duration(10000 * time.Millisecond))
	//go data.AutoMessage(messages)
	time.Sleep(time.Duration(10000 * time.Millisecond))
	go data.AutoLogout(logout_users)
	data.Wg.Wait()
}

/* function for run generate csv log for send messages
   @param --> null
   @param value --> null
   description --> function for generate csv log for send messages
   @return --> null
*/
func logService(){
	msgList := getSendMessageJsonList()
	saveSendMessageCSV(msgList)
}
