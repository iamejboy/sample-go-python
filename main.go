package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"os/exec"
)

const (
	EnvPort = "PORT"
)

func main() {
	
	cmd := exec.Command("python",  "-c", "import pythonfile; print pythonfile.cat_strings('Hello ', 'RPaaS-v2')")
	out, err := cmd.CombinedOutput()
    	if err != nil { fmt.Println(err); }
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("%s\nYou are a rock star!", out))
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			io.WriteString(w, fmt.Sprintf("%s: %s\n", pair[0], pair[1]))
		}
	})

	port := os.Getenv(EnvPort)
	log.Println("[INFO] Listening on", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
