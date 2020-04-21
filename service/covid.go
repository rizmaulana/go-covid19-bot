package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ivandzf/go-covid19-bot/client"
)

const (
	generalErrorMessage          = "Mohon maaf, data sedang tidak tersedia, silakan coba beberapa saat lagi."
	locationNotFoundErrorMessage = "Mohon maaf, ID LOKASI Anda tidak ditemukan. Lihat ID LOKASI yang tersedia dengan perintah /loc"
)

type covid struct {
	clientSvc client.Client
}

type Covid interface {
	GetInformation() string
	GetInformationByLocation(loc string) string
	GetAvailableLocation() string
	GetOverviewStatistic() string
	GetInformationByCoordinate(lat float32, lon float32) string
}

func NewCovidService(clientSvc client.Client) Covid {
	return &covid{clientSvc}
}

func (c *covid) GetInformation() string {
	var info strings.Builder
	info.WriteString("Selamat datang di covid19kalsel : Telegram bot untuk informasi perkembangan covid19 di Kalimantan Selatan")
	info.WriteString("\n")
	info.WriteString("Gunakan command berikut untuk melihat informasi")
	info.WriteString("\n\n")
	info.WriteString("/info - Menampilkan informasi penggunaan bot")
	info.WriteString("\n")
	info.WriteString("/all - Menampilkan seluruh statistik di Kalimantan Selatan")
	info.WriteString("\n")
	info.WriteString("/loc - Menampilkan informasi ID LOKASI kabupaten / kota yang tersedia")
	info.WriteString("\n")
	info.WriteString("/loc <ID LOKASI> - Menampilkan informasi perkembangan covid19 di lokasi yang dimaksud, contoh : /loc BJM")
	info.WriteString("\n\n")
	info.WriteString("Atau kirimkan lokasimu untuk mendapatkan data covid19 di sekitar lokasimu sekarang")
	info.WriteString("\n\n")
	info.WriteString("Bagikan bot dengan share link t.me/covid19kalselbot")
	info.WriteString("\n")
	info.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")

	return info.String()
}

func (c *covid) GetInformationByLocation(loc string) string {
	covidData, err := c.clientSvc.FetchCovidData()
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return generalErrorMessage
	}

	var selectedLocation *client.CovidResponse
	for _, v := range covidData {
		if v.Code == loc {
			fmt.Println("Location assign : ", v.Code)
			selectedLocation = &v
			break
		}
	}

	if selectedLocation != nil {
		var result strings.Builder
		result.WriteString("Informasi Covid19 di ")
		result.WriteString(selectedLocation.Name)
		result.WriteString("\n")
		result.WriteString("Positif : ")
		result.WriteString(strconv.Itoa(selectedLocation.CovPositiveCount))
		result.WriteString("\n")
		result.WriteString("Sembuh : ")
		result.WriteString(strconv.Itoa(selectedLocation.CovRecoveredCount))
		result.WriteString("\n")
		result.WriteString("Meninggal : ")
		result.WriteString(strconv.Itoa(selectedLocation.CovDiedCount))
		result.WriteString("\n")
		result.WriteString("PDP : ")
		result.WriteString(strconv.Itoa(selectedLocation.CovPdpCount))
		result.WriteString("\n")
		result.WriteString("ODP : ")
		result.WriteString(strconv.Itoa(selectedLocation.CovOdpCount))
		result.WriteString("\n\n")
		result.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")

		return result.String()
	}

	return locationNotFoundErrorMessage
}

func (c *covid) GetAvailableLocation() string {
	covidData, err := c.clientSvc.FetchCovidData()
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return generalErrorMessage
	}

	var result strings.Builder
	for _, v := range covidData {
		result.WriteString("ID LOKASI : ")
		result.WriteString(v.Code)
		result.WriteString(" - ")
		result.WriteString(v.Name)
		result.WriteString("\n")
	}

	result.WriteString("\n")
	result.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")
	return result.String()
}

func (c *covid) GetOverviewStatistic() string {
	covidData, err := c.clientSvc.FetchCovidData()
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return generalErrorMessage
	}

	var result strings.Builder
	result.WriteString("Statistik seluruh Kalsel")
	result.WriteString("\n")
	result.WriteString("Lokasi")
	result.WriteString("\t\t")
	result.WriteString("Positif")
	result.WriteString("\t\t")
	result.WriteString("Sembuh")
	result.WriteString("\t\t")
	result.WriteString("Meninggal")
	result.WriteString("\t\t")
	result.WriteString("ODP")
	result.WriteString("\t\t")
	result.WriteString("PDP")
	result.WriteString("\t\t")
	result.WriteString("\n")

	for _, v := range covidData {
		result.WriteString(v.Code)
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(v.CovPositiveCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(v.CovRecoveredCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(v.CovDiedCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(v.CovOdpCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(v.CovPdpCount))
		result.WriteString("          ")
		result.WriteString("\n")
	}

	result.WriteString("\n")
	result.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")

	return result.String()
}

func (c *covid) GetInformationByCoordinate(lat float32, lon float32) string {
	coordinateData, err := c.clientSvc.ReverseCoordinate(lat, lon)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return generalErrorMessage
	}

	responseView := coordinateData.Response.View
	if len(responseView) == 0 {
		return locationNotFoundErrorMessage
	}

	result := responseView[0].Result
	if len(result) == 0 {
		return locationNotFoundErrorMessage
	}

	return c.GetInformationByLocation(result[0].Location.Address.City)
}
