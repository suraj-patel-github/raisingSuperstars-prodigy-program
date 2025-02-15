package types

type Activity struct {
	ID        int    `json:"id"`
	Category  string `json:"category"`
	Time      string `json:"time"`
	Frequency string `json:"freq"`
	Activity  string `json:"activity"`
}
