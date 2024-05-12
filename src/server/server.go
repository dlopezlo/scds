package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/dlopezlo/scds/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Element struct {
	Host         []string `json:"host"`
	Hostname     string   `json:"hostname"`
	Port         int      `json:"port,omitempty"`
	ProxyCommand string   `json:"proxyCommand,omitempty"`
	ProxyJump    string   `json:"proxyjump,omitempty"`
	User         string   `json:"user,omitempty"`
}

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	// Middlewares
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Timeout(time.Second * 3))
	// handlers
	s.Router.Get("/", WelcomeHandler)
	s.Router.Get("/sshconfig/{username}", getSSHConfigHandler)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to my TFG"))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func getSSHConfigHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Hosts    []Element `json:"hosts"`
		Username string
	}
	// grab username and make some sanity checks
	username, err := utils.GetUsername(chi.URLParam(r, "username"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data.Username = username
	// read the file with the config in json format
	sshData, err := os.ReadFile("ssh-data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error reading data file:", err)
		return
	}

	// Read the byts into an object
	err = json.Unmarshal(sshData, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error parsing json:", err)
		return
	}

	// load the template
	tpl, err := template.New("template.tpl").ParseFiles("template.tpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error with templating system:", err)
	}

	// Execute renders the template with the data and writes to
	// responseWrite back to the user.
	err = tpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error with templating system:", err)
	}
}
