package dto

type Transaction struct {
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type Location struct {
	City string `json:"city"`
}

type Details struct {
	Sum   float64
	Avg   float64
	Max   float64
	Min   float64
	Count float64
}
