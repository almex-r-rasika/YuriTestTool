package data

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type LogoutRequestBody struct {
    ID string `json:"id"`
}

type Logout struct {
	Error string      `json:"error"`
	ErrorCode string      `json:"errorCode"`
	ErrorString string      `json:"errorString"`
}

/* function for each user to logout
   @param --> id, logout_time
   @param value --> loginuser id, logout time
   description --> function for each user to logout
   @return --> Logout object
*/
func DoLogout(id string, logout_time string) Logout{

	requestBody := &LogoutRequestBody{
        ID: id,
    }

    jsonString, err := json.Marshal(requestBody)
    if err != nil {
        Log.Fatal(err.Error())
    }

	Log.Debug("logout request body "+string(jsonString))

	endpoint := "https://msgc.smapa-checkout.jp/v1/ctrl/logout"
    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
    if err != nil {
        Log.Fatal(err.Error())
    }

    req.Header.Set("Content-Type", "application/json")

	// add CA Certificate
	client := addCaCertificate()

	// our context, we use context.Background()
	ctx := context.Background()

	// when we want to wait till
	until, _ := time.Parse(time.RFC3339, logout_time)

	// waiting till the logout_time
	waitUntil(ctx, until)

    response, err := client.Do(req)
    if err != nil {
        Log.Fatal(err.Error())
    }

	defer response.Body.Close()
	
   // Read the response body on the line below
   body, err := ioutil.ReadAll(response.Body)
   if err != nil {
      Log.Fatal(err.Error())
   }

   Log.Debug("logout response body "+string(body))

   // Convert the body to type logout object
   var logout Logout
   json.Unmarshal([]byte(body), &logout)

   // validate the response
   if(logout.Error == "0"){
	Log.Info("user "+id+" logout")
   }else{
    Log.Error("error "+logout.Error+" errorcode "+logout.ErrorCode+" errorstring "+logout.ErrorString)
   }
   return logout
}

/* function for when the logout_time comes users to autologout
   @param --> logout_users
   @param value --> LoginUser objects array
   description --> for each user when the logout time comes autologout happens
   @return --> null
*/
func AutoLogout(logout_users []LoginUser){

	for _, user := range logout_users {
		DoLogout(user.UserId,user.LogoutTime)
	}

	Log.Info("all users have finished logout")
}


