package main

import (
	"encoding/json"
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

func executeByToggle(userKey string) []string {
	user := ffuser.NewUser(userKey)
	ftKey1000, _ := ffclient.StringVariation("key-1000", user, "")
	ftKey2000, _ := ffclient.StringVariation("key-2000", user, "")
	ftKey3000, _ := ffclient.StringVariation("key-3000", user, "")

	return []string{ftKey1000, ftKey2000, ftKey3000}
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func writeJson(w *http.ResponseWriter, keys []string) {
	keysJson, _ := json.Marshal(keys)
	fmt.Fprintf(*w, "{ %s }", string(keysJson))
}

func handleRequests() {

	print("Starting route register")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {}) // Avoid duplicate calls on browser
	http.HandleFunc("/health", health)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
    	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		keys := executeByToggle("user-A")
		keys = deleteEmpty(keys)
		json.NewEncoder(w).Encode(keys)

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
			Path: "/etc/config/release-toggles",
			//Path: "test/toggles/keys.yaml",
		},
	})
	println("Carregando File Retriever")
	defer ffclient.Close()

	if err == nil {
		handleRequests()
		println("Inicio main")
	}

}
