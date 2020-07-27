package jkdcovid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	url    = "http://stats-covid.com/api/"
	client = &http.Client{}
)

// GetCSSEData get request for api/csse/{country} endpoint
// Response example:
/*
	{
    	"country": "US",
    	"data": [
        	{
            	"county": "<nil>",
            	"province": "Diamond Princess",
            	"cases": 49,
            	"deaths": 0,
            	"recovered": 0
        	},
        	{
            	"county": "<nil>",
            	"province": "Grand Princess",
            	"cases": 103,
            	"deaths": 3,
            	"recovered": 0
			}
		]
	}
*/
func GetCSSEData(country string) (CSSEResponse, error) {

	body, err := getRequestWithParameter("csse/", country)
	if err != nil {
		return CSSEResponse{}, err
	}

	var responseData CSSEResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return CSSEResponse{}, errUnmarshal
	}
	return responseData, nil
}

// GetHotspotData get request for api/hotspot/{days} endpoint
// Response example:
/*
	{
    "mostCases": {
        "country": "Brazil",
        "data": [
            20599,
            26417,
            26928,
            33274,
            16409,
            11598,
            28936,
            28633,
            30925,
            30830,
            27075
        ]
    },
    "secondCases": {
        "country": "USA",
        "data": [
            18263,
            22577,
            24266,
            24146,
            20007,
            20848,
            20801,
            19699,
            21140,
            24720,
            22681
        ]
    },
    "thirdCases": {
        "country": "Russia",
        "data": [
            8338,
            8371,
            8572,
            8952,
            9268,
            8485,
            8858,
            8529,
            8823,
            8718,
            8846
        ]
    },
    "mostDeaths": {
        "country": "USA",
        "data": [
            1505,
            1199,
            1193,
            967,
            605,
            768,
            1031,
            995,
            1036,
            921,
            670
        ]
    },
    "secondDeaths": {
        "country": "Brazil",
        "data": [
            1086,
            1156,
            1124,
            956,
            480,
            623,
            1262,
            1349,
            1473,
            1005,
            904
        ]
    },
    "thirdDeaths": {
        "country": "Mexico",
        "data": [
            463,
            447,
            371,
            364,
            151,
            237,
            470,
            1092,
            816,
            625,
            341
        ]
    }
}
*/
func GetHotspotData(nDays int) (HotSpotResponse, error) {
	body, err := getRequestWithParameter("hotspot/", strconv.Itoa(nDays))
	if err != nil {
		return HotSpotResponse{}, err
	}

	var responseData HotSpotResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return HotSpotResponse{}, errUnmarshal
	}
	return responseData, nil
}

// GetWorldHistoryData get request for api/world endpoint
// Response example:
/*
	{
    "cases": [
        555,
        654,
        941,
        1434,

    ],
    "deaths": [
        17,
        18,
        26,
        42,
        56,
        82,
    ],
    "recovered": [
        28,
        30,
        36,
        39,
        52,
    ],
    "casesDaily": [
        99,
        287,
        493,
        684,
        809,
    ],
    "deathsDaily": [
        1,
        8,
        16,
        14,
        26,
        49,
    ],
    "recoveredDaily": [
        2,
        6,
        3,
        13,
        9,
        46,
        19,
        17,
    ]
}
*/
func GetWorldHistoryData() (WorldHistoryResponse, error) {
	body, err := getRequestNoParameter("world")
	if err != nil {
		return WorldHistoryResponse{}, err
	}

	var responseData WorldHistoryResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return WorldHistoryResponse{}, errUnmarshal
	}
	return responseData, nil
}

// GetContinentData get request for api/continent endpoint
// Response example:
/*
	{[
    {
        "updated": 1591969378141,
        "cases": 1313866,
        "todayCases": 884,
        "deaths": 56394,
        "todayDeaths": 21,
        "recovered": 678380,
        "todayRecovered": 111,
        "active": 579092,
        "critical": 11988,
        "casesPerOneMillion": 3051.48,
        "deathsPerOneMillion": 130.98,
        "tests": 5424046,
        "testsPerOneMillion": 12597.45,
        "population": 430566996,
        "continent": "South America",
        "activePerOneMillion": 1344.95,
        "recoveredPerOneMillion": 1575.55,
        "criticalPerOneMillion": 27.84,
        "countries": [
            "Argentina",
            "Bolivia",
            "Brazil",
            "Chile",
            "Colombia",
            "Ecuador",
            "Falkland Islands (Malvinas)",
            "French Guiana",
            "Guyana",
            "Paraguay",
            "Peru",
            "Suriname",
            "Uruguay",
            "Venezuela"
        ]
    },
    {
        "updated": 1591969378149,
        "cases": 8901,
        "todayCases": 5,
        "deaths": 124,
        "todayDeaths": 0,
        "recovered": 8371,
        "todayRecovered": 22,
        "active": 406,
        "critical": 2,
        "casesPerOneMillion": 217.71,
        "deathsPerOneMillion": 3.03,
        "tests": 2070918,
        "testsPerOneMillion": 50652.24,
        "population": 40885025,
        "continent": "Australia/Oceania",
        "activePerOneMillion": 9.93,
        "recoveredPerOneMillion": 204.74,
        "criticalPerOneMillion": 0.05,
        "countries": [
            "Australia",
            "Fiji",
            "French Polynesia",
            "New Caledonia",
            "New Zealand",
            "Papua New Guinea"
        ]
    }
]

*/
func GetContinentData() (ContinentDataResponse, error) {
	body, err := getRequestNoParameter("continent")
	if err != nil {
		return ContinentDataResponse{}, err
	}

	var responseData ContinentDataResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return ContinentDataResponse{}, errUnmarshal
	}
	return responseData, nil
}

// GetAllCountriesData get request for api/countries endpoint
// Response example:
/*
  {
		"data": [
			{
				"country": "Zimbabwe",
				"cases": 7,
				"todayCases": 0,
				"deaths": 1,
				"todayDeaths": 0,
				"recovered": 0,
				"active": 6,
				"critical": 0,
				"casesPerOneMillion": 0.5
			}
		]
	}

*/
func GetAllCountriesData() (AllCountriesDataResponse, error) {
	body, err := getRequestNoParameter("countries")
	if err != nil {
		return AllCountriesDataResponse{}, err
	}

	var responseData AllCountriesDataResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return AllCountriesDataResponse{}, errUnmarshal
	}
	return responseData, nil
}

// GetWorldTotalData get request for api/total endpoint
// Response example:
/*
	{
		"todayPerCentOfTotalCases": 7,
		"todayPerCentOfTotalDeaths": 6,
		"totalCases": 1188489,
		"totalDeaths": 64103,
		"todayTotalCases": 71846,
		"todayTotalDeaths": 4933
	}

*/
func GetWorldTotalData() (TotalResponse, error) {
	body, err := getRequestNoParameter("total")
	if err != nil {
		return TotalResponse{}, err
	}

	var responseData TotalResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return TotalResponse{}, errUnmarshal
	}
	return responseData, nil
}

// CountryData post request for api/country endpoint
/*
Request example:
	{
		"country" : "Greece",
	}

Response example:

	{
		"country": "Greece",
		"cases": 1061,
		"todayCases": 0,
		"deaths": 37,
		"todayDeaths": 5,
		"recovered": 52,
		"active": 972,
		"critical": 66,
		"casesPerOneMillion": 102
	}

*/
func CountryData(country string) (CountryResponse, error) {
	type rj struct {
		Country string `json:"country"`
	}

	requestJSON := rj{Country: country}
	jsonString, _ := json.Marshal(requestJSON)

	body, err := post("country", string(jsonString))
	if err != nil {
		return CountryResponse{}, err
	}

	var responseData CountryResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return CountryResponse{}, errUnmarshal
	}
	return responseData, nil
}

// SortedData post request for api/sort endpoint
/*
	Request example
		{
			"type" : "deaths"
		}

	Types:
		[
			deaths,
			cases,
			todayCases,
			todayDeaths,
			recovered,
			active,
			critical,
			casesPerOneMillion
		]

	Response example
	{
		"data": [{
			"country": "Italy",
			"cases": 124632,
			"todayCases": 4805,
			"deaths": 15362,
			"todayDeaths": 681,
			"recovered": 20996,
			"active": 88274,
			"critical": 3994,
			"casesPerOneMillion": 2061
		},
		{
			"country": "Spain",
			"cases": 124736,
			"todayCases": 5537,
			"deaths": 11744,
			"todayDeaths": 546,
			"recovered": 34219,
			"active": 78773,
			"critical": 6416,
			"casesPerOneMillion": 2668
		}]
	}
*/
func SortedData(sortType string) (SortResponse, error) {
	type rj struct {
		Type string `json:"type"`
	}

	requestJSON := rj{Type: sortType}
	jsonString, _ := json.Marshal(requestJSON)

	body, err := post("sort", string(jsonString))
	if err != nil {
		return SortResponse{}, err
	}

	var responseData SortResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return SortResponse{}, errUnmarshal
	}
	return responseData, nil
}

// CompareCountriesData post request for api/compare/all endpoint
/*
	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "dataDeaths": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsFromFirst": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsPerDay": [
            110,
            52,
            688,
            50,
            74
        ],
        "dataRecoverd": [
            150376,
            150376,
            150376,
            150376,
            150376
        ],
        "dataCases": [
            15,
            32,
            45,
            84,
            120,
            165,
            222,
            259,
        ],
        "dataCasesFromFirst": [
            159,
            294,
            394,
            334,
            318,
            332,
            240
        ]
    },
    "countryTwo": {
        "country": "Spain",
        "dataDeaths": [
            0,
            0,
            0,
            0,
            0,
        ],
        "dataDeathsFromFirst": [
            33530,
            33601,
            33689,
            33774,
            33846,
            33899
        ],
        "dataDeathsPerDay": [
            1,
            1,
            1,
            4,
            3,
            2,
        ],
        "dataRecoverd": [
            0,
            1,
            1,
            1,
            2,
            3,
            45,
            46,
            46,
        ],
        "dataCases": [
            232664,
            232997,
            233197,
            233515,
            233836,
            234013,
            234531,
            234801,
            234998
        ],
        "dataCasesFromFirst": [
            200,
            318,
            321,
            177,
            518,
            270,
            197
        ]
    }
}
*/
func CompareCountriesData(countryOne string, countryTwo string) (CompareResponse, error) {
	type rj struct {
		One string `json:"countryOne"`
		Two string `json:"countryTwo"`
	}

	requestJSON := rj{One: countryOne, Two: countryTwo}
	jsonString, _ := json.Marshal(requestJSON)

	body, err := post("compare/all", string(jsonString))
	if err != nil {
		return CompareResponse{}, err
	}

	var responseData CompareResponse
	if errUnmarshal := json.Unmarshal(body, &responseData); errUnmarshal != nil {
		return CompareResponse{}, errUnmarshal
	}
	return responseData, nil
}

func getRequestWithParameter(endpoint string, parameter string) ([]byte, error) {
	requestURL := url + endpoint + parameter
	return get(requestURL)
}

func getRequestNoParameter(endpoint string) ([]byte, error) {
	requestURL := url + endpoint
	return get(requestURL)
}

func get(getURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	return body, bodyErr
}

func post(endpoint string, payload string) ([]byte, error) {
	postURL := url + endpoint
	req, err := http.NewRequest("POST", postURL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	return body, bodyErr
}
