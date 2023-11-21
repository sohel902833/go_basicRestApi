package models

import "time"

type Cashier struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	Passcode string `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	ID       string    `json:"_id" bson:"_id,omitempty"`
	Email    string    `json:"email" bson:"email,omitempty"`
	Image    string    `json:"image" bson:"image,omitempty"`
	SignupAt time.Time `json:"signup_at" bson:"signup_at,omitempty"`
}