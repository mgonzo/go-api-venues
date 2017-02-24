package model

type Venue struct {
	Id       int
	Uuid	 string  `json:"uuid"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	City	 City `json:"city"`
	Phone    Phone
}
