package balance

type Movement struct {
	UserId      string  `json:"userId"`
	Time    	int64   `json:"time"`
	Description string  `json:"description"`
	Value 		float64 `json:"value"`
}

