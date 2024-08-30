package db

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"profiles_go/models"
	"time"
)

// CreateUserInput describes the required fields to create a new user
type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

func CreateUser(input CreateUserInput) (*models.User, error) {
	// Check if username or email already exists
	existingUser, err := GetUserByUsername(input.Username)
	if err != nil {
		return nil, fmt.Errorf("error checking existing username: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	existingUser, err = GetUserByEmail(input.Email)
	if err != nil {
		return nil, fmt.Errorf("error checking existing email: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("email already exists")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}
	user.PasswordHash = string(hashedPassword)

	query := `INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = DB.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	// Create a profile with default values
	profileInput := CreateProfileInput{
		UserID: user.ID,
		Avatar: "https://www.gravatar.com/avatar/3b3be63a4c2a439b013787725dfce802?d=identicon",
		Bio:    "Hi, I'm new here!",
	}
	_, err = CreateProfile(profileInput)
	if err != nil {
		return nil, fmt.Errorf("error creating profile for user: %w", err)
	}

	return user, nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash FROM users WHERE username = $1`
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1`
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// UpdateUser updates user details in the database
func UpdateUser(username string, input UpdateUserInput) (*models.User, error) {
	user := &models.User{
		Username:  username,
		Email:     input.Email,
		UpdatedAt: time.Now(),
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hashedPassword)
	}

	query := `UPDATE users SET email = $2, password_hash = $3, updated_at = $4 WHERE username = $1 RETURNING id, created_at`
	err := DB.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.UpdatedAt).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user from the database
func DeleteUser(username string) error {
	query := `DELETE FROM users WHERE username = $1`
	_, err := DB.Exec(query, username)
	return err
}
