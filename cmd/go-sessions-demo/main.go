package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var sessionStore = sessions.NewCookieStore(
	[]byte(os.Getenv("SESSION_AUTHENTICATION_KEY")),
	[]byte(os.Getenv("SESSION_ENCRYPTION_KEY")))

func home(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Application Error", http.StatusInternalServerError)
		return
	}

	if username := session.Values["username"]; username == "" {
		http.Redirect(w, r, "/login.html", 302)
	} else {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, "Hello %s <a href='/logout'>Logout</a>", username)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "foo" && password == "secret" {
		session, _ := sessionStore.Get(r, "session-name")
		session.Values["username"] = username
		session.Save(r, w)

		log.WithField("username", username).Info("session.Save")
	}

	http.Redirect(w, r, "/", 302)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "session-name")
	session.Values["username"] = ""
	session.Save(r, w)

	log.WithField("username", "").Info("session.Save")
	http.Redirect(w, r, "/", 302)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/", home)

	n := negroni.Classic()
	n.UseHandler(context.ClearHandler(mux))

	port := os.Getenv("PORT")
	if port == "" {
		log.WithField("PORT", port).Fatal("$PORT must be set")
	}
	n.Run(":" + port)
}