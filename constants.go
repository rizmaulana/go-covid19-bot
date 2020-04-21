package main

import (
	"strconv"
	"strings"
)

const (
	ApiCovidKalsel = "" //Put url data source API

	ApiHereMaps = "" //Put HERE Maps API

	ErrorMessage = "Mohon maaf, data sedang tidak tersedia, silakan coba beberapa saat lagi."
)

func Information() string {
	info := strings.Builder{}
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

func CovidInformation(selectedLocation *CovidDto) string {
	result := strings.Builder{}
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

func StatisticInformation(covids []CovidDto) string {
	result := strings.Builder{}
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

	for _, cov := range covids {
		result.WriteString(cov.Code)
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(cov.CovPositiveCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(cov.CovRecoveredCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(cov.CovDiedCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(cov.CovOdpCount))
		result.WriteString("          ")
		result.WriteString(strconv.Itoa(cov.CovPdpCount))
		result.WriteString("          ")
		result.WriteString("\n")
	}
	result.WriteString("\n")
	result.WriteString("with ❤️ : github.com/rizmaulana #DiRumahAja")

	return result.String()
}
