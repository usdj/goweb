package main

import (
		"io"
		"net/http"
		"os/exec"
		"log"
	   )

func reLaunch(){
    cmd := exec.Command("sh", "./deploy.sh")
    err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>U cant't see me!I'm a deployer</h1>")
	reLaunch()
}


func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
