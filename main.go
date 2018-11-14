package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"strings"
)




// ============================================================
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/here-isochrone/{lng}/{lat}/{time}/{appid}/{appcode}", v1HereIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8003", router))

}
// ============================================================


// ============================================================
func v1HereIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var jsonResult map[string]string

	if isochrone, msg := v1DoHereIsochrone(params["lng"], params["lat"], params["time"], params["appid"], params["appcode"]); msg == "" {
		jsonResult = map[string]string{"here": isochrone}
	} else {
		jsonResult = map[string]string{"intersects": ""}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonResult)
}
// ============================================================


// ============================================================
func v1DoHereIsochrone(sxLng string, syLat string, sTime string, sAppID string, sAppCode string) (geojson string, msg string) {

	here_url := "https://isoline.route.api.here.com/routing/7.2/calculateisoline.json?app_id=" + sAppID + "&app_code=" + sAppCode + "&mode=fastest;car;traffic:disabled&start=geo!" + syLat + "," + sxLng + "&range=" + sTime + "&rangetype=time"

	startSearchText := "[{id:0,shape:"
	endSearchText   := "}]}],start:"

	geojson = ""
	msg     = ""

	response, err := http.Get(here_url)
	if err == nil {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			geojson = ""
			msg     = err.Error()
		} 

		jsonText := strings.Replace(string(body), "\"", "", -1)

		nStart   := strings.Index(jsonText, startSearchText) + len(startSearchText)
		nEnd     := strings.Index(jsonText, endSearchText)

		x := strings.Split(jsonText[nStart:nEnd], ",")

		var s []string
		for n := 0; n < len(x); n+=2 {
			
			switch num := n; num {
				case 0:
					s = append(s, (x[n] + "," + x[n+1] +"],"))
				case len(x) - 2:
					s = append(s, ("[" + x[n] + "," + x[n+1]))
				default:
					s = append(s, "[" + (x[n] + "," + x[n+1]) + "],")
			}
		}

		geojson = "[" + strings.Join(s, "") + "]"
	} 

	return
}
// ============================================================
