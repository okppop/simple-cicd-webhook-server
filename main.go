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
		log.Fatal("exit: server init: ", err)
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

	log.Print("Info: webhookHandler: new request, calling action")

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
	// change this to any action you want
	cmd := exec.Command("/usr/bin/ls", "-l")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
