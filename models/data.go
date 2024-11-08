package models

type Data struct {
	SwimCount          string `json:"swim_pool_count"`
	SwimPoolLimit      int    `json:"swim_pool_limit"`
	SwimPercentage     string `json:"swim_pool_percentage,omitempty"`
	FitnessCount       string `json:"fitness_center_count"`
	FitnessCenterLimit int    `json:"fitness_center_limit"`
	FitnessPercentage  string `json:"fitness_center_percentage,omitempty"`
}
