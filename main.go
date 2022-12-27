package main

import (
	"fmt"
	"net/http"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/notifier"
	"github.com/thomaspoignant/go-feature-flag/notifier/slacknotifier"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
	fmt.Println("Endpoint Hit: Health check ok!")
}

func executeByToggle(userKey string, w *http.ResponseWriter) {
	user := ffuser.NewUser(userKey)
	ftKey1000, _ := ffclient.StringVariation("key-1000", user, "")
	fmt.Fprintf(*w, "Executing toggle: %s for %s \n", ftKey1000, userKey)
	println(ftKey1000)
}

func handleRequests() {

	print("Starting route register")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {}) // Avoid duplicate calls on browser
	http.HandleFunc("/health", health)                                               // Avoid duplicate calls on browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		executeByToggle("user-A", &w)
		executeByToggle("user-B", &w)
		executeByToggle("user-C", &w)
		executeByToggle("user-D", &w)
		executeByToggle("user-E", &w)

		fmt.Println("Endpoint Hit")
	})

	http.ListenAndServe(":10000", nil)
	print("Ending route register")
}

func main() {
	println("Inicio main")

	err := ffclient.Init(ffclient.Config{
		Notifiers: []notifier.Notifier{
			&slacknotifier.Notifier{
				SlackWebhookURL: "https://hooks.slack.com/services/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",
			},
		},
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			//Path: "/etc/config/release-toggles",
			Path: "test/toggles/keys.yaml",
		},
	})
	println("Carregando File Retriever")
	defer ffclient.Close()

	if err == nil {
		user := ffuser.NewUser("user-A")
		ftKey1000, _ := ffclient.StringVariation("key-1000", user, "")
		println(ftKey1000)
		handleRequests()
		println("Inicio main")
	}

}
