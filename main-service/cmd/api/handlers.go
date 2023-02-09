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

func (app *Config) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	// ParseForm parses the raw query from the URL and updates r.Form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return 
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//data.Log.Info(username)
	//data.Log.Info(password)

	// check if user exists 
	storedPassword, exists := users[username]
	if exists {
		// Get registers and returns a session for the given name and session store.
		// session.id is the name of the cookie that will be stored in the client's browser
		session, _ := store.Get(r, "session.id")
		if storedPassword == password {
			session.Values["authenticated"] = true 
			// saves all sessions used during the current request
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
		w.Write([]byte("Login successfully!"))
	}
}

func (app *Config) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get registers and returns a session for the given name and session store.
	session, _ := store.Get(r, "session.id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("Logout Successful"))
}

func (app *Config) healthcheck(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		w.Write([]byte("Welcome!"))
		return
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
    return
	}
}
