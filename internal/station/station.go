package station

type Station struct {
    ID   int    `json:"nid"`
    Name string `json:"title"`
}

type StationResponse struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}