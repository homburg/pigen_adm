package main

import (
	"bytes"
	"fmt"
	"homburg/pigen_adm/res"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const listenAddr = "127.0.0.1:8086"

var logins map[string]string
var baseLogins []string

func accessControl(w http.ResponseWriter, r *http.Request) bool {
	return true
}

// Convert newlines to <br>
func newlineToHtmlBreak(s string) string {
	return strings.Replace(s, "\n", "<br>", -1)
}

// Run command and return html
func commandToHtml(firstCmd string, cmdSegments ...string) (string, error) {
	cmd := exec.Command(firstCmd, cmdSegments...)
	out, err := cmd.Output()
	if nil != err {
		return "", err
	}

	outStr := strings.TrimRight(string(out), "\n")
	return outStr, nil
}

type templateData struct {
	Hostname  string
	GoVersion string
}

func main() {
	// html, err := ioutil.ReadFile(filepath.Join(exePath, "server.template.html"))
	tmpl := template.New("server")
	template.Must(tmpl.Parse(pigen_adm.ServerTemplate))

	var buf bytes.Buffer
	hostname, _ := os.Hostname()
	err := tmpl.Execute(&buf, templateData{hostname, runtime.Version()})
	html := buf.String()

	if nil != err {
		log.Fatal(err)
	}

	log.Println("Started")

	// landscape-sysinfo
	http.HandleFunc("/landscape/sysinfo", func(w http.ResponseWriter, r *http.Request) {
		if !accessControl(w, r) {
			return
		}

		outStr, _ := commandToHtml("landscape-sysinfo")
		fmt.Fprintln(w, outStr)
	})

	// post actions
	http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		if !accessControl(w, r) {
			return
		}

		if r.Method == "POST" {
			action := r.FormValue("action")

			if action == "make_thumbnails" {
				cmd := exec.Command("sudo", "-u", "thomas", "/home/thomas/bin/make_thumbnails")
				out, err := cmd.Output()
				if nil != err {
					log.Fatal(err)
				}
				fmt.Fprint(w, string(out))
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if "/" != r.URL.String() {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if !accessControl(w, r) {
			return
		}

		fmt.Fprint(w, string(html))
	})

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
