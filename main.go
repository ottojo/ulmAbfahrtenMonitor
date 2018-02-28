package main

import (
	"log"
	"github.com/ottojo/ulmAbfahrtenMonitor/swu"
)

const stopId = "1255" //Test Value, Stop "Heilmeyersteige"

func main() {

	stops := swu.GetStopsNear(48.412970, 9.950862) //Test Value, Coordinates near Heilmeyersteige
	log.Println("Stops: ", stops)
	log.Println(printDepartures(swu.NewSession(stopId).GetDepartures(), 5))
}

func printDepartures(departures []swu.ItdDeparture, count int) string {
	s := "\n"
	for i, d := range departures {
		if i >= count {
			return s
		}
		s += d.ItdServingLine.Number + " Richtung " + d.ItdServingLine.Direction + " in " + d.Countdown + " min.\n"
	}
	return s
}
