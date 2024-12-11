package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Fatal: http.ListenAndServe: ", err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// no need to print the payload to stdout

	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Print("Error: webhookHandler: io.ReadAll: ", err)
	// }
	// defer r.Body.Close()

	// log.Print("Info: webhookHandler: new request, payload: ")
	// fmt.Println(string(body))

	log.Print("Info: /webhook: receive a new request")

	if r.Method != "POST" {
		log.Print("Warn: /webhook: http method isn't 'POST': drop")
		return
	}

	var httpHeaderXGitHubEvent string = r.Header.Get("X-GitHub-Event")

	if httpHeaderXGitHubEvent == "" {
		log.Print("Warn: /webhook: http header 'X-GitHub-Event' is empty: drop")
		return
	}

	if httpHeaderXGitHubEvent == "ping" {
		log.Print("Info: /webhook: http header 'X-GitHub-Event' is 'ping': ignore")
		return
	}

	if httpHeaderXGitHubEvent != "push" {
		log.Printf("Info: /webhook: http 'X-GitHub-Event' is a undefinded event '%s': ignore", httpHeaderXGitHubEvent)
		return
	}

	log.Print("Info: /webhook: detect 'push' event, calling action")

	err := action()
	if err != nil {
		log.Print("Error: action error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print("Info: action success")
	w.WriteHeader(http.StatusOK)
}

func action() error {
	cmd := exec.Command("/usr/bin/ls", "-l")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
