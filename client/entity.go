package client

type CovidResponse struct {
	Name              string `json:"name"`
	CovPositiveCount  int    `json:"cov_positive_count"`
	CovRecoveredCount int    `json:"cov_recovered_count"`
	CovDiedCount      int    `json:"cov_died_count"`
	CovOdpCount       int    `json:"cov_odp_count"`
	CovPdpCount       int    `json:"cov_pdp_count"`
	Code              string `json:"code"`
}

type HereMapsResponse struct {
	Response Response `json:"Response"`
}

type Response struct {
	View []View `json:"View"`
}

type View struct {
	Result []Result `json:"Result"`
}

type Result struct {
	Location Location `json:"Location"`
}

type Location struct {
	Address Address `json:"Address"`
}

type Address struct {
	City string `json:"City"`
}
