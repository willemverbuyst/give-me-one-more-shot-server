package responses

type HTTPError struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"messasge" example:"bad request"`
}
