package main

import (
	"fmt"
	"net/http"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
	fmt.Println("Endpoint Hit: Health check ok!")
}

func handleRequests() {

	print("Starting route register")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {}) // Avoid duplicate calls on browser
	http.HandleFunc("/health", health)                                               // Avoid duplicate calls on browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userA := ffuser.NewUser("user-A")
		ftKey1000A, _ := ffclient.StringVariation("key-1000", userA, "")
		fmt.Fprintln(w, "Executing toggle: "+ftKey1000A+" for user-A!")
		println(ftKey1000A)

		userB := ffuser.NewUser("user-B")
		ftKey1000B, _ := ffclient.StringVariation("key-1000", userB, "")
		fmt.Fprintln(w, "Executing toggle: "+ftKey1000B+" for user-B!")
		println(ftKey1000B)

		userC := ffuser.NewUser("user-C")
		ftKey1000C, _ := ffclient.StringVariation("key-1000", userC, "")
		fmt.Fprintln(w, "Executing toggle: "+ftKey1000C+" for user-C!")
		println(ftKey1000C)

		userD := ffuser.NewUser("user-D")
		ftKey1000D, _ := ffclient.StringVariation("key-1000", userD, "")
		fmt.Fprintln(w, "Executing toggle: "+ftKey1000D+" for user-D!")
		println(ftKey1000D)

		userE := ffuser.NewUser("user-E")
		ftKey1000E, _ := ffclient.StringVariation("key-1000", userE, "")
		fmt.Fprintln(w, "Executing toggle: "+ftKey1000E+" for user-E!")
		println(ftKey1000E)

		fmt.Println("Endpoint Hit")
	})

	http.ListenAndServe(":10000", nil)
	print("Ending route register")
}

func main() {
	println("Inicio main")

	err := ffclient.Init(ffclient.Config{
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
