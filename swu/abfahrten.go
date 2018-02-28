package swu

import (
	"bytes"
	"io/ioutil"
	"encoding/xml"
	"regexp"
	"net/url"
	"log"
	"net/http"
	"strconv"
)

func GetStopsNear(lat, long float64) []Stop {
	request, state := getRequestState()
	return getStopsList(lat, long, request, state)
}

func getRequestState() (request string, state string) {
	resp, err := http.Get("https://echtzeit.swu.de/")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var regexRequest = regexp.MustCompile(`(?ms)var request = (\d+);.*var state = (\d+);`)
	match := regexRequest.FindStringSubmatch(string(body))
	if len(match) != 3 {
		log.Panic("No valid request and state found.")
	}
	request = match[1]
	state = match[2]
	return
}

func getStopsList(lat, long float64, request, state string) []Stop {

	data := url.Values{}
	data.Set("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	data.Add("lng", strconv.FormatFloat(long, 'f', -1, 64))
	data.Add("request", request)
	data.Add("state", state)

	apiUrl := "https://echtzeit.swu.de/php/phpsqlajax_genxml.php?src=stopsn"

	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var r xmlStops
	xml.Unmarshal(body, &r)

	return r.Stops

}

func getDepartures(sessionID string) []ItdDeparture {
	resp, err := http.Get("https://www.ding.eu/ding3/XSLT_DM_REQUEST?useRealtime=1&sessionID=" + sessionID + "&requestID=1&dmLineSelectionAll=1&limit=1000&outputFormat=XML")
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	var r itdRequest
	xml.Unmarshal(body, &r)
	return r.ItdDepartureMonitorRequest.ItdDepartureList.ItdDepartures
}

func getSessionID(stopID string) string {
	resp, err := http.Get("https://www.ding.eu/ding3/XSLT_DM_REQUEST?sessionID=0&type_dm=stopID&name_dm=900" +
		stopID + "&outputFormat=XML&useRealtime=1")
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	var r itdRequest
	xml.Unmarshal(body, &r)
	return r.SessionID
}
