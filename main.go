package main

import (
	"fmt"
	"net/http"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/k8sretriever"
	"k8s.io/client-go/rest"
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
		user := ffuser.NewUser("user-A")
		ftKey1000, _ := ffclient.StringVariation("key-1000", user, "")
		println(ftKey1000)
		fmt.Fprintf(w, "Executing toggle: "+ftKey1000+"!")
		fmt.Println("Endpoint Hit")
	})

	http.ListenAndServe(":10000", nil)
	print("Ending route register")
}

func main() {
	println("Inicio main")

	var err error
	// if len(os.Args) > 1 {
	// 	switch os.Args[1] {
	// 	case "file":
	// err := ffclient.Init(ffclient.Config{
	// 	PollingInterval: 3 * time.Second,
	// 	Retriever: &fileretriever.Retriever{
	// 		Path: "test/toggles/keys.yaml",
	// 	},
	// })
	// println("Carregando File Retriever")
	// defer ffclient.Close()
	// case "k8s":
	config, err := rest.InClusterConfig()
	if err != nil {
		print("Cluster k8s não encontrado")
		panic("Cluster k8s não encontrado")
	}

	err = ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second,
		Retriever: &k8sretriever.Retriever{
			Namespace:     "default",
			ConfigMapName: "ft-data",
			Key:           "key-1000",
			ClientConfig:  *config,
		},
	})
	println("Carregando K8s Retriever")
	defer ffclient.Close()
	// default:
	// 	println("Escolha uma das opções de retriever: [File, k8s]")
	// 	return
	// }

	if err == nil {
		user := ffuser.NewUser("user-A")
		ftKey1000, _ := ffclient.StringVariation("key-1000", user, "")
		println(ftKey1000)
		handleRequests()
		println("Inicio main")
	}
	// } else {
	// 	println("Escolha uma das opções de retriever: [File, k8s]")
	// 	return
	// }

}
