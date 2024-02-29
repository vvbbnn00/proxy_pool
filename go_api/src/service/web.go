package service

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// ProxyNode is a struct to store the proxy information
type ProxyNode struct {
	Anonymous  string `json:"anonymous"`
	CheckCount int    `json:"check_count"`
	FailCount  int    `json:"fail_count"`
	Https      bool   `json:"https"`
	LastStatus bool   `json:"last_status"`
	LastTime   string `json:"last_time"`
	Proxy      string `json:"proxy"`
	Region     string `json:"region"`
	Source     string `json:"source"`
}

// routeGetProxy is a handler to get a random proxy
func routeGetProxy(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Get a random proxy from the Proxies map
	mutex.RLock()
	var selectedProxy string
	keyLen := len(Keys)
	if keyLen > 0 {
		selectedProxy = Proxies[Keys[rand.Intn(keyLen)]] // Randomly select a proxy
	}
	mutex.RUnlock()

	// Load the proxy information
	proxyNode := ProxyNode{}
	err := json.Unmarshal([]byte(selectedProxy), &proxyNode)
	if err != nil {
		log.Println("Error unmarshalling the proxy:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response
	err = json.NewEncoder(w).Encode(proxyNode)
	if err != nil {
		log.Println("Error encoding the response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// RegisterRoutes is a function to register the routes
func RegisterRoutes() {
	// Register the /get route
	http.HandleFunc("/get", routeGetProxy)
	http.HandleFunc("/get/", routeGetProxy)
}
