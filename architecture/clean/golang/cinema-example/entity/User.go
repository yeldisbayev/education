package entity

import "time"

// User struct
type User struct {
	ID          int64     `json:"id" bson:"id"`
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"string" bson:"string"`
	Password    string    `json:"password" bson:"password"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	ValidatedAt time.Time `json:"validatedAt" bson:"validatedAt"`
}
