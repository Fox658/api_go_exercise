package main

import (
	"fmt"
	"net/http"

	"github.com/Fox658/api_go_exercise/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service")

	fmt.Println(`
 ________  ________          ________  ________  ___          _______      ___    ___ _______   ________     
|\   ____\|\   __  \        |\   __  \|\   __  \|\  \        |\  ___ \    |\  \  /  /|\  ___ \ |\   __  \    
\ \  \___|\ \  \|\  \       \ \  \|\  \ \  \|\  \ \  \       \ \   __/|   \ \  \/  / | \   __/|\ \  \|\  \   
 \ \  \  __\ \  \\\  \       \ \   __  \ \   ____\ \  \       \ \  \_|/__  \ \    / / \ \  \_|/_\ \   _  _\  
  \ \  \|\  \ \  \\\  \       \ \  \ \  \ \  \___|\ \  \       \ \  \_|\ \  /     \/   \ \  \_|\ \ \  \\  \| 
   \ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\       \ \_______\/  /\   \    \ \_______\ \__\\ _\ 
    \|_______|\|_______|        \|__|\|__|\|__|     \|__|        \|_______/__/ /\ __\    \|_______|\|__|\|__|
                                                                          |__|/ \|__|                        
	`)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
