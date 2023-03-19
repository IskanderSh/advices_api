package api

type Slip struct {
	Id     int    `json:"id"`
	Advice string `json:"advice"`
	Date   string `json:"date"`
}

type Advice struct {
	Slip         Slip
}

type Searchadvice struct {
	TotalResults string `json:"total_results"`
	Query        string `json:"query"`
	Slips         []Slip
}
