package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type MessageRequestBody struct {
    PatientNumList []string `json:"patient_num_list"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Url string `json:"url"`
}

type Message struct {
	Error string      `json:"error"`
	ErrorCode string      `json:"errorCode"`
	ErrorString string      `json:"errorString"`
}

var count int

/* function for send message to the given destination 
   @param --> subject, line1, line2, line3, line4, line5, line6, line7, line8, line9, line10, loginId, patientList[], sendTime, token, count
   @param value --> subject, line1, line2, line3, line4, line5, line6, line7, line8, line9, line10, loginId, patientList[], sendTime, token, count
   description --> for each message when the send time comes message will send to the destination
   @return --> Message Object
*/
func DoMessage(subject string,line1 string,line2 string,line3 string,line4 string,line5 string,line6 string,line7 string,line8 string,line9 string, line10 string,loginId string, patientList []string,sendTime string, token string, count int) Message{

	var patients = patientList
	subject = "[yuri-test] "+time.Now().Format("hh:mm:ss ")+strconv.Itoa(count)+" "+subject

	jsonbody := getMessageBody(line1, line2, line3, line4, line5, line6, line7, line8, line9, line10)

	requestBody := &MessageRequestBody{
        PatientNumList: patients,
		Subject: subject,
		Body: jsonbody,
		Url: "",
    }

    jsonString, err := json.Marshal(requestBody)
    if err != nil {
        Log.Fatal(err.Error())
    }

	Log.Debug("message request body "+string(jsonString))

	endpoint := "https://msgc.smapa-checkout.jp/v1/hospital/messages/send"
    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
    if err != nil {
        Log.Fatal(err.Error())
    }

    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// add CA Certificate
	client := addCaCertificate()

	// our context, we use context.Background()
	ctx := context.Background()

	// when we want to wait till
	until, _ := time.Parse(time.RFC3339, sendTime)

	// and now we are waiting till the send_time
	waitUntil(ctx, until)

    response, err := client.Do(req)
    if err != nil {
        Log.Fatal(err.Error())
    }

	defer response.Body.Close()
	
   // Read the response body on the line below.
   body, err := ioutil.ReadAll(response.Body)
   if err != nil {
      Log.Fatal(err.Error())
   }

   Log.Debug("message response body "+string(body))
   
   // Convert the body to type message object
   var message Message
   json.Unmarshal([]byte(body), &message)

   // validate the response
   if(message.Error == "0"){
	Log.Info("user "+loginId+" sent message at "+sendTime)
   }else{
    Log.Error("error "+message.Error+" errorcode "+message.ErrorCode+" errorstring "+message.ErrorString)
   }
   return message
}

/* function for when the send_time comes send messages to destination
   @param --> messages
   @param value --> Messages objects array
   description --> for each message when the send time comes send messages to destination 
   @return --> null
*/
func AutoMessage(messages []Messages){

	for _, message := range messages {
		count++
		logged_users := GetLoggedUsers()
		token := logged_users[message.SendUserId]
		patient := GetPatientList(message.Address, token)
	    DoMessage(message.Subject,message.Line1,message.Line2,message.Line3,message.Line4,message.Line5,message.Line6,message.Line7,message.Line8,message.Line9,message.Line10,message.SendUserId,patient,message.SendTime,token,count)
	}

	Log.Info("all messages have sent")
}

/* function for create message body for a message
   @param --> null
   @param value --> null
   description --> create message body for DoMessage function
   @return --> message body
*/
func getMessageBody(line1 string, line2 string, line3 string, line4 string, line5 string, line6 string, line7 string, line8 string, line9 string, line10 string) string{

	jsonbody := fmt.Sprintf("[{\"index\":0,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":1,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":2,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":3,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":4,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":5,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":6,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":7,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":8,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":9,\"message\":\"%s\",\"color\":\"#000000\"},{\"index\":10,\"message\":\"\",\"color\":\"#000000\"},{\"index\":11,\"message\":\"\",\"color\":\"#000000\"},{\"index\":12,\"message\":\"\",\"color\":\"#000000\"},{\"index\":13,\"message\":\"\",\"color\":\"#000000\"},{\"index\":14,\"message\":\"\",\"color\":\"#000000\"},{\"index\":15,\"message\":\"\",\"color\":\"#000000\"},{\"index\":16,\"message\":\"\",\"color\":\"#000000\"},{\"index\":17,\"message\":\"\",\"color\":\"#000000\"},{\"index\":18,\"message\":\"\",\"color\":\"#000000\"},{\"index\":19,\"message\":\"\",\"color\":\"#000000\"},{\"index\":20,\"message\":\"\",\"color\":\"#000000\"},{\"index\":21,\"message\":\"\",\"color\":\"#000000\"},{\"index\":22,\"message\":\"\",\"color\":\"#000000\"},{\"index\":23,\"message\":\"\",\"color\":\"#000000\"},{\"index\":24,\"message\":\"\",\"color\":\"#000000\"},{\"index\":25,\"message\":\"\",\"color\":\"#000000\"},{\"index\":26,\"message\":\"\",\"color\":\"#000000\"},{\"index\":27,\"message\":\"\",\"color\":\"#000000\"},{\"index\":28,\"message\":\"\",\"color\":\"#000000\"},{\"index\":29,\"message\":\"\",\"color\":\"#000000\"}]",line1, line2, line3, line4, line5, line6, line7, line8, line9, line10)

	return jsonbody
}

