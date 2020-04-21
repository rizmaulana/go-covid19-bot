package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func OverviewStatistic() string {
	var covids, err = FecthCovidData()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return ErrorMessage
	}
	return StatisticInformation(covids)
}
func AvailableLocation() string {
	var covids, err = FecthCovidData()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return ErrorMessage
	}
	result := strings.Builder{}
	for _, cov := range covids {
		result.WriteString("ID LOKASI : ")
		result.WriteString(cov.Code)
		result.WriteString(" - ")
		result.WriteString(cov.Name)
		result.WriteString("\n")
	}
	result.WriteString("\n")
	result.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")
	return result.String()
}

func GetInformationByLoc(loc string) string {
	var covids, err = FecthCovidData()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return ErrorMessage
	}

	var selectedLocation *CovidDto
	for _, cov := range covids {
		if cov.Code == loc {
			fmt.Println("Location assing!", cov.Code)
			selectedLocation = &cov
			break
		}
	}
	if selectedLocation != nil {
		return CovidInformation(selectedLocation)
	}
	return "Mohon maaf, ID LOKASI Anda tidak ditemukan. Lihat ID LOKASI yang tersedia dengan perintah /loc"
}

func FecthCovidData() ([]CovidDto, error) {
	var err error
	var data []CovidDto

	response, err := http.Get(ApiCovidKalsel)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(responseData, &data)
	fmt.Printf("%+v\n", data)

	return data, nil
}

func GetInformationByCoordinate(lat float32, lon float32) string {
	var ress, err = ReverseCoordinate(lat, lon)
	if err != nil {
		return ErrorMessage
	}
	var selectedLocation = ress.Response.View[0].Result[0].Location.Address.City
	return GetInformationByLocName(selectedLocation)

}

func GetInformationByLocName(locName string) string {
	var covids, err = FecthCovidData()
	if err != nil {
		return ErrorMessage
	}
	var selectedLocation *CovidDto
	fmt.Printf("%+v\n", covids)

	for _, cov := range covids {
		if strings.Contains(cov.Name, locName) {
			selectedLocation = &cov
			break
		}
	}
	if selectedLocation != nil {
		return CovidInformation(selectedLocation)
	}
	return "Mohon maaf data pada lokasi Anda tidak ditemukan, silakan coba masukan informasi berdasarkan kode lokasi Anda.\nLokasi Anda : " + locName

}

func ReverseCoordinate(lat float32, lon float32) (*HereMapsResponse, error) {
	var hereMapsResponse HereMapsResponse
	var apiKey = os.Getenv("HEREMAPS_API_KEY")
	req, err := http.NewRequest("GET", ApiHereMaps, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("prox", fmt.Sprintf("%f", lat)+","+fmt.Sprintf("%f", lon))
	q.Add("apiKey", apiKey)
	req.URL.RawQuery = q.Encode()
	response, err := http.Get(req.URL.String())

	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(responseData, &hereMapsResponse)
	return &hereMapsResponse, nil
}
