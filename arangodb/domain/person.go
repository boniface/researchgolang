package domain

// User bind struct
type Person struct {
	Personid string `json:"personid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}
