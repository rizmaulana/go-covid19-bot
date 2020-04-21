package main

type CovidDto struct {
	Name              string `json:"name"`
	CovPositiveCount  int    `json:"cov_positive_count"`
	CovRecoveredCount int    `json:"cov_recovered_count"`
	CovDiedCount      int    `json:"cov_died_count"`
	CovOdpCount       int    `json:"cov_odp_count"`
	CovPdpCount       int    `json:"cov_pdp_count"`
	Code              string `json:"code"`
}
