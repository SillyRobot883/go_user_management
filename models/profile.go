package models

import "time"

type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicProfile struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

type ProfileDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewProfile(userID int, avatar, bio string) *Profile {
	return &Profile{
		UserID:    userID,
		Avatar:    avatar,
		Bio:       bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
