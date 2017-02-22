package model

type Venue struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
  Timezone string `json:"timezone"`
  Zip int `json:"zip"`
  Phone Phone
}

