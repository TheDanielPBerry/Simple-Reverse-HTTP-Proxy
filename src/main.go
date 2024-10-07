package main

import (
	"fmt"
	ini "gopkg.in/ini.v1"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var portMap map[string]int

func MapPort(w http.ResponseWriter, r *http.Request, sitePrefix string) {
	if sitePrefix == "" {
		sitePrefix = "default"
	}

	port, ok := portMap[sitePrefix]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	path := strings.Replace(r.URL.Path, sitePrefix, "", 1)
	url := fmt.Sprintf("http://localhost:%d/%s?%s", port, path, r.URL.Query().Encode())

	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		return
	}
	r.URL.Query()

	req.Header = r.Header.Clone()
	req.Form = r.Form

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	for name, value := range resp.Header {
		w.Header().Add(name, value[0])
	}
	io.Copy(w, resp.Body)
}

func InitializePortMap(cfg *ini.File) (map[string]int, bool) {
	section, err := cfg.GetSection("PORTS")
	if err != nil {
		fmt.Println("Invalid PORTS section", err)
		return nil, false
	}

	newPorts := make(map[string]int)
	for route, port := range section.KeysHash() {
		portValue, err := strconv.Atoi(port)
		if err == nil {
			newPorts[route] = portValue
			continue
		}
		fmt.Println("Invalid Route=Port entry: ", err)
	}
	return newPorts, true
}

func main() {
	//Initialize port config
	cfg, err := ini.Load("../config.ini")
	if err != nil {
		fmt.Println("Fail to read file:", err)
		return
	}
	var ok bool
	portMap, ok = InitializePortMap(cfg)
	if !ok {
		return
	}

	//Refresh the port mappings for a route without restarting application
	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		cfg, err := ini.Load("../config.ini")
		if err != nil {
			fmt.Println("Fail to read file:", err)
			return
		}

		passCode := cfg.Section("").Key("REFRESH_PASSCODE").String()

		userIn := r.FormValue("passcode")
		if userIn != passCode {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
		}

		newPorts, ok := InitializePortMap(cfg)
		if !ok {
			fmt.Println("Invalid PORTS section", err)
			fmt.Fprint(w, "Invalid configuration")
			return
		}

		portMap = newPorts
		fmt.Println("Refresh Successful: \n", portMap)
		fmt.Fprint(w, "Refresh Successful")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sitePrefix := strings.Split(r.URL.Path, "/")
		MapPort(w, r, sitePrefix[1])
	})

	http.ListenAndServe(":80", nil)
}
