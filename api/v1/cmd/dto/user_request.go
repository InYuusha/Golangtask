package dto


type UserRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Dob         string `json:"dob"`
	Description string `json:"description"`
}