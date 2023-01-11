package main

import (
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	MessageList []SendMessageLog`json:"send_message_list"`
}

/* GET API for get send messages log
   @param --> http.ResponseWriter, *http.Request
   @param value --> w http.ResponseWriter, r *http.Request
   description --> GET API for get send messages log
   @return --> null
*/
func (app *Config) getMessageLog(w http.ResponseWriter, r *http.Request) {

	logList := getSendMessageJsonList()


	payload := jsonResponse{
		Error:   false,
		MessageList: logList,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}