package db

import (
	"database/sql"
	"fmt"
	"profiles_go/models"
	"time"
)

// CreateProfileInput describes the required fields to create a new profile
type CreateProfileInput struct {
	UserID int    `json:"user_id" binding:"required"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

// UpdateProfileInput describes the required fields to update a profile
type UpdateProfileInput struct {
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

// CreateProfile creates a new profile in the database
func CreateProfile(input CreateProfileInput) (*models.Profile, error) {
	profile := &models.Profile{
		UserID:    input.UserID,
		Avatar:    input.Avatar,
		Bio:       input.Bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `INSERT INTO profiles (user_id, avatar, bio, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := DB.QueryRow(query, profile.UserID, profile.Avatar, profile.Bio, profile.CreatedAt, profile.UpdatedAt).Scan(&profile.ID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetProfileByUserID retrieves a profile by user ID
func GetProfileByUserID(userID int) (*models.Profile, error) {
	profile := &models.Profile{}
	query := `SELECT id, user_id, avatar, bio, created_at, updated_at FROM profiles WHERE user_id = $1`
	err := DB.QueryRow(query, userID).Scan(&profile.ID, &profile.UserID, &profile.Avatar, &profile.Bio, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return profile, nil
}

// UpdateProfile update basic profile info like avatar and bio
func UpdateProfile(userID int, input UpdateProfileInput) (*models.Profile, error) {
	profile := &models.Profile{
		UserID:    userID,
		Avatar:    input.Avatar,
		Bio:       input.Bio,
		UpdatedAt: time.Now(),
	}

	query := `UPDATE profiles SET avatar=$1, bio=$2, updated_at=$3 WHERE user_id=$4 RETURNING id`
	err := DB.QueryRow(query, profile.Avatar, profile.Bio, profile.UpdatedAt, profile.UserID).Scan(&profile.ID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// this function should use the PublicProfile struct
func GetPublicProfileByUsername(username string) (*models.PublicProfile, error) {
	var publicProfile models.PublicProfile

	publicProfile.Username = username
	fmt.Println("username: ", username)
	// Get user ID from username
	var userID int
	query := `SELECT id FROM users WHERE username = $1`
	err := DB.QueryRow(query, username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Get public profile details using user ID
	query = `SELECT avatar, bio FROM profiles WHERE user_id = $1`
	err = DB.QueryRow(query, userID).Scan(&publicProfile.Avatar, &publicProfile.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &publicProfile, nil
}

// this will be used in other repo functions
func GetUsernameByID(id int) (string, error) {
	var username string
	query := `SELECT username FROM users WHERE id = $1`
	err := DB.QueryRow(query, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

// this is used to get the profile of a user by their username
func GetProfileByUsername(username string) (*models.Profile, error) {
	var profile models.Profile
	query := `SELECT id, user_id, avatar, bio, created_at, updated_at FROM profiles WHERE user_id = (SELECT id FROM users WHERE username = $1)`
	err := DB.QueryRow(query, username).Scan(&profile.ID, &profile.UserID, &profile.Avatar, &profile.Bio, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func UpdateProfileByUsername(username string, input UpdateProfileInput) (*models.Profile, error) {
	var userID int
	query := `SELECT id FROM users WHERE username = $1`
	err := DB.QueryRow(query, username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user ID by username: %w", err)
	}

	return UpdateProfile(userID, input)
}

// DeleteProfileByUsername deletes the profile of a user by their username
func DeleteProfileByUsername(username string) error {
	// Get user ID from username
	var userID int
	query := `SELECT id FROM users WHERE username = $1`
	err := DB.QueryRow(query, username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("error retrieving user ID by username: %w", err)
	}

	// Delete profile by user ID
	query = `DELETE FROM users WHERE id = $1`
	_, err = DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error deleting profile by username: %w", err)
	}

	return nil
}
