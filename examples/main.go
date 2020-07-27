package main

import (
	"fmt"

	jkdcovid "github.com/junkd0g/client_covid"
)

/*
	go get "github.com/junkd0g/client_covid"
	go run main.go
*/

func main() {

	//get request to api/csse/{country}
	csseData, cssDataError := jkdcovid.GetCSSEData("USA")
	if cssDataError != nil {
		fmt.Println(cssDataError.Error())
	}
	fmt.Println(csseData)

	// get request to  api/hotspot/{days}
	hotspotData, hotspotDataError := jkdcovid.GetHotspotData(15)
	if hotspotDataError != nil {
		fmt.Println(hotspotDataError.Error())
	}
	fmt.Println(hotspotData)

	//get request to api/world
	worldHistoryData, worldHistoryDataError := jkdcovid.GetWorldHistoryData()
	if worldHistoryDataError != nil {
		fmt.Println(worldHistoryDataError.Error())
	}
	fmt.Println(worldHistoryData)

	//get request to api/continent
	continentData, continentDataError := jkdcovid.GetContinentData()
	if continentDataError != nil {
		fmt.Println(continentDataError.Error())
	}
	fmt.Println(continentData)

	//get request to api/total
	worldTotalData, worldTotalDataError := jkdcovid.GetWorldTotalData()
	if worldTotalDataError != nil {
		fmt.Println(worldTotalDataError.Error())
	}
	fmt.Println(worldTotalData)

	//post request to api/country
	countryData, countryDataError := jkdcovid.CountryData("Greece")
	if countryDataError != nil {
		fmt.Println(countryDataError.Error())
	}
	fmt.Println(countryData)

	//post request to api/sort
	sortedData, sortedDataError := jkdcovid.SortedData("cases")
	if sortedDataError != nil {
		fmt.Println(sortedDataError.Error())
	}
	fmt.Println(sortedData)

	//post request to api/compare/all
	compareCountriesData, compareCountriesDataError := jkdcovid.CompareCountriesData("Greece", "S. Korea")
	if compareCountriesDataError != nil {
		fmt.Println(compareCountriesDataError.Error())
	}
	fmt.Println(compareCountriesData)

}
