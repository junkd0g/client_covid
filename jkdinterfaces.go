package jkdcovid

type CSSEResponse struct {
	Country string `json:"country"`
	Data    []struct {
		County    string `json:"county"`
		Province  string `json:"province"`
		Cases     int    `json:"cases"`
		Deaths    int    `json:"deaths"`
		Recovered int    `json:"recovered"`
	} `json:"data"`
}

type HotSpotResponse struct {
	MostCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"mostCases"`
	SecondCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"secondCases"`
	ThirdCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"thirdCases"`
	MostDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"mostDeaths"`
	SecondDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"secondDeaths"`
	ThirdDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"thirdDeaths"`
}

type WorldHistoryResponse struct {
	Cases          []int `json:"cases"`
	Deaths         []int `json:"deaths"`
	Recovered      []int `json:"recovered"`
	CasesDaily     []int `json:"casesDaily"`
	DeathsDaily    []int `json:"deathsDaily"`
	RecoveredDaily []int `json:"recoveredDaily"`
}

type ContinentDataResponse []struct {
	Updated                int64    `json:"updated"`
	Cases                  int      `json:"cases"`
	TodayCases             int      `json:"todayCases"`
	Deaths                 int      `json:"deaths"`
	TodayDeaths            int      `json:"todayDeaths"`
	Recovered              int      `json:"recovered"`
	TodayRecovered         int      `json:"todayRecovered"`
	Active                 int      `json:"active"`
	Critical               int      `json:"critical"`
	CasesPerOneMillion     float64  `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float64  `json:"deathsPerOneMillion"`
	Tests                  int      `json:"tests"`
	TestsPerOneMillion     float64  `json:"testsPerOneMillion"`
	Population             int      `json:"population"`
	Continent              string   `json:"continent"`
	ActivePerOneMillion    float64  `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64  `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float64  `json:"criticalPerOneMillion"`
	Countries              []string `json:"countries"`
}

type AllCountriesDataResponse struct {
	Data []struct {
		Country            string `json:"country"`
		Cases              int    `json:"cases"`
		TodayCases         int    `json:"todayCases"`
		Deaths             int    `json:"deaths"`
		TodayDeaths        int    `json:"todayDeaths"`
		Recovered          int    `json:"recovered"`
		Active             int    `json:"active"`
		Critical           int    `json:"critical"`
		CasesPerOneMillion int    `json:"casesPerOneMillion"`
		Tests              int    `json:"tests"`
		TestsPerOneMillion int    `json:"testsPerOneMillion"`
	} `json:"data"`
}

type TotalResponse struct {
	TodayPerCentOfTotalCases  int `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int `json:"todayPerCentOfTotalDeaths"`
	TotalCases                int `json:"totalCases"`
	TotalDeaths               int `json:"totalDeaths"`
	TodayTotalCases           int `json:"todayTotalCases"`
	TodayTotalDeaths          int `json:"todayTotalDeaths"`
}

type CountryResponse struct {
	Country            string `json:"country"`
	Cases              int    `json:"cases"`
	TodayCases         int    `json:"todayCases"`
	Deaths             int    `json:"deaths"`
	TodayDeaths        int    `json:"todayDeaths"`
	Recovered          int    `json:"recovered"`
	Active             int    `json:"active"`
	Critical           int    `json:"critical"`
	CasesPerOneMillion int    `json:"casesPerOneMillion"`
	Tests              int    `json:"tests"`
	TestsPerOneMillion int    `json:"testsPerOneMillion"`
}

type SortResponse struct {
	Data []struct {
		Country            string `json:"country"`
		Cases              int    `json:"cases"`
		TodayCases         int    `json:"todayCases"`
		Deaths             int    `json:"deaths"`
		TodayDeaths        int    `json:"todayDeaths"`
		Recovered          int    `json:"recovered"`
		Active             int    `json:"active"`
		Critical           int    `json:"critical"`
		CasesPerOneMillion int    `json:"casesPerOneMillion"`
		Tests              int    `json:"tests"`
		TestsPerOneMillion int    `json:"testsPerOneMillion"`
	} `json:"data"`
}

type CompareResponse struct {
	CountryOne struct {
		Country             string `json:"country"`
		DataDeaths          []int  `json:"dataDeaths"`
		DataDeathsFromFirst []int  `json:"dataDeathsFromFirst"`
		DataDeathsPerDay    []int  `json:"dataDeathsPerDay"`
		DataRecoverd        []int  `json:"dataRecoverd"`
		DataCases           []int  `json:"dataCases"`
		DataCasesFromFirst  []int  `json:"dataCasesFromFirst"`
	} `json:"countryOne"`
	CountryTwo struct {
		Country             string `json:"country"`
		DataDeaths          []int  `json:"dataDeaths"`
		DataDeathsFromFirst []int  `json:"dataDeathsFromFirst"`
		DataDeathsPerDay    []int  `json:"dataDeathsPerDay"`
		DataRecoverd        []int  `json:"dataRecoverd"`
		DataCases           []int  `json:"dataCases"`
		DataCasesFromFirst  []int  `json:"dataCasesFromFirst"`
	} `json:"countryTwo"`
}
