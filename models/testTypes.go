package models

type User struct {
		ID string `json:"id"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		PictureURL string `json:"pictureURL"`
		PhoneNumber string `json:"phoneNumber"`
		Email float64
		Rating *Rating `json:"rating"`
}
