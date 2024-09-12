package postgres

import (
	"auth/api/helper"
	"auth/genproto/auth"
	"auth/token"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
)

type AuthStorage struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (stg *AuthStorage) GetProfile(req *auth.ById) (*auth.GetProfileResponse, error) {
	query := `
		SELECT 
			id, 
			first_name, 
			last_name, 
			date_of_birth, 
			gender, 
			email, 
			role 
		FROM users 
		WHERE id = $1
	`

	var profile auth.GetProfileResponse

	err := stg.db.QueryRow(query, req.Id).Scan(
		&profile.Id,
		&profile.FirstName,
		&profile.LastName,
		&profile.DateOfBirth,
		&profile.Gender,
		&profile.Email,
		&profile.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %s not found", req.Id)
		}
		return nil, fmt.Errorf("failed to get profile: %v", err)
	}

	return &profile, nil
}

func (stg *AuthStorage) UpdateProfile(req *auth.UpdateProfileRequest) (*auth.InfoResponse, error) {
	query := `
		UPDATE users
		SET 
			first_name = $1,
			last_name = $2,
			date_of_birth = $3,
			gender = $4
		WHERE id = $5
	`

	result, err := stg.db.Exec(query, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.Id)
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to update profile: %v", err),
			Success: false,
		}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to get affected rows: %v", err),
			Success: false,
		}, err
	}

	if rowsAffected == 0 {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("user with ID %s not found", req.Id),
			Success: false,
		}, nil
	}

	return &auth.InfoResponse{
		Message: "Profile updated successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) DeleteProfile(req *auth.ById) (*auth.InfoResponse, error) {
	currentUnixTime := time.Now().Unix()

	query := `
		UPDATE users
		SET deleted_at = $1
		WHERE id = $2
	`

	result, err := stg.db.Exec(query, currentUnixTime, req.Id)
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to delete profile: %v", err),
			Success: false,
		}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to get affected rows: %v", err),
			Success: false,
		}, err
	}

	if rowsAffected == 0 {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("user with ID %s not found", req.Id),
			Success: false,
		}, nil
	}

	return &auth.InfoResponse{
		Message: "Profile deleted successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) ListAllProfiles(req *auth.ById) (*auth.GetAllProfilesResponse, error) {
	query := `
		SELECT id, first_name, last_name, date_of_birth, gender, email, role
		FROM users
		WHERE deleted_at = 0
	`

	rows, err := stg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*auth.GetProfileResponse
	for rows.Next() {
		var profile auth.GetProfileResponse
		if err := rows.Scan(
			&profile.Id,
			&profile.FirstName,
			&profile.LastName,
			&profile.DateOfBirth,
			&profile.Gender,
			&profile.Email,
			&profile.Role,
		); err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &auth.GetAllProfilesResponse{
		Result: profiles,
	}, nil
}

func (stg *AuthStorage) DeleteProfileByAdmin(req *auth.DeleteProfileByAdminReq) (*auth.InfoResponse, error) {
	query := `
		UPDATE users
		SET deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)::BIGINT
		WHERE id = $1
	`

	result, err := stg.db.Exec(query, req.UserId)
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to delete profile: %v", err),
			Success: false,
		}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("failed to get affected rows: %v", err),
			Success: false,
		}, err
	}

	if rowsAffected == 0 {
		return &auth.InfoResponse{
			Message: fmt.Sprintf("user with ID %s not found", req.UserId),
			Success: false,
		}, nil
	}

	return &auth.InfoResponse{
		Message: "Profile successfully deleted by admin",
		Success: true,
	}, nil
}

func (stg *AuthStorage) RegisterUser(req *auth.Register) (*auth.InfoResponse, error) {
	ctx := context.Background()

	reqData, err := proto.Marshal(req)
	if err != nil {
		return &auth.InfoResponse{
			Message: "error serializing request data",
			Success: false,
		}, err
	}

	err = stg.rdb.Set(ctx, req.Code, reqData, time.Minute*3).Err()
	if err != nil {
		return &auth.InfoResponse{
			Message: "error storing data in Redis",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "The verification code has been send you email. Please verify it!",
		Success: true,
	}, nil
}

func (stg *AuthStorage) RegisterCourier(req *auth.Register) (*auth.InfoResponse, error) {
	ctx := context.Background()

	reqData, err := proto.Marshal(req)
	if err != nil {
		return &auth.InfoResponse{
			Message: "error serializing request data",
			Success: false,
		}, err
	}

	err = stg.rdb.Set(ctx, req.Code, reqData, time.Minute*3).Err()
	if err != nil {
		return &auth.InfoResponse{
			Message: "error storing data in Redis",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "user registration data saved successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) Login(req *auth.Login) (*auth.TokenResponse, error) {
	query := `
		SELECT id, password, role
		FROM users
		WHERE email = $1 AND deleted_at = 0
	`

	var userId, hashedPassword, role string
	err := stg.db.QueryRow(query, req.Email).Scan(&userId, &hashedPassword, &role)
	if err != nil {
		return &auth.TokenResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return &auth.TokenResponse{}, fmt.Errorf("password or email is incorrect")
	}

	tokens := token.GenerateJWTToken(userId, req.Email, role)

	return &auth.TokenResponse{
		Id:          userId,
		AccessToken: tokens.AccessToken,
		ExpiresAt:   time.Now().Add(1 * time.Hour).Format(time.RFC3339),
	}, nil
}

func (stg *AuthStorage) VerifyEmail(req *auth.VerifyEmailRequest) (*auth.InfoResponse, error) {
	ctx := context.Background()

	reqData, err := stg.rdb.Get(ctx, req.Code).Bytes()
	if err != nil {
		if err == redis.Nil {
			return &auth.InfoResponse{
				Message: "verification code or email is incorrect",
				Success: false,
			}, nil
		}
		return &auth.InfoResponse{
			Message: "error retrieving data from Redis",
			Success: false,
		}, err
	}

	var registerReq auth.Register
	if err := proto.Unmarshal(reqData, &registerReq); err != nil {
		return &auth.InfoResponse{
			Message: "error deserializing request data",
			Success: false,
		}, err
	}

	if registerReq.Email != req.Email {
		return &auth.InfoResponse{
			Message: "verification code or email is incorrect",
			Success: false,
		}, nil
	}

	if err := stg.rdb.Del(ctx, req.Code).Err(); err != nil {
		return &auth.InfoResponse{
			Message: "error deleting verification code from Redis",
			Success: false,
		}, err
	}

	query := `INSERT INTO users (id, first_name, last_name, date_of_birth, gender, email, password, role)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = stg.db.ExecContext(ctx, query, registerReq.Id, registerReq.FirstName, registerReq.LastName,
		registerReq.DateOfBirth, registerReq.Gender, registerReq.Email, registerReq.Password, registerReq.Role)
	if err != nil {
		return &auth.InfoResponse{
			Message: "error inserting user into database",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "email verified successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) RefreshToken(req *auth.ById) (*auth.TokenResponse, error) {
	query := `
		SELECT id, email, role
		FROM users
		WHERE id = $1 AND deleted_at = 0
	`

	var userId, email, role string
	err := stg.db.QueryRow(query, req.Id).Scan(&userId, &email, &role)
	if err != nil {
		return &auth.TokenResponse{}, fmt.Errorf("user not found: %w", err)
	}

	tokens := token.GenerateJWTToken(userId, email, role)

	return &auth.TokenResponse{
		Id:          userId,
		AccessToken: tokens.AccessToken,
		ExpiresAt:   time.Now().Add(1 * time.Hour).Format(time.RFC3339),
	}, nil
}

func (stg *AuthStorage) ChangePassword(req *auth.ChangePasswordRequest) (*auth.InfoResponse, error) {
	var storedHash string
	query := `
		SELECT password
		FROM users
		WHERE id = $1 AND deleted_at = 0
	`

	err := stg.db.QueryRow(query, req.Id).Scan(&storedHash)
	if err != nil {
		return &auth.InfoResponse{
			Message: "user not found or error retrieving user",
			Success: false,
		}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.CurrentPassword))
	if err != nil {
		return &auth.InfoResponse{
			Message: "current password is incorrect",
			Success: false,
		}, err
	}

	newPasswordHash, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return &auth.InfoResponse{
			Message: "error hashing new password",
			Success: false,
		}, err
	}

	updateQuery := `
		UPDATE users
		SET password = $1
		WHERE id = $2
	`

	_, err = stg.db.Exec(updateQuery, newPasswordHash, req.Id)
	if err != nil {
		return &auth.InfoResponse{
			Message: "error updating password",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "password changed successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) ForgetPassword(req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error) {
	ctx := context.Background()

	err := stg.rdb.Set(ctx, req.Code, req.Email, time.Minute*3).Err()
	if err != nil {
		return &auth.InfoResponse{
			Message: "error storing data in Redis",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "Verification code has been sent to your email, please verify it",
		Success: true,
	}, nil
}

func (stg *AuthStorage) ResetPassword(req *auth.ResetPasswordRequest) (*auth.InfoResponse, error) {
	ctx := context.Background()

	storedEmail, err := stg.rdb.Get(ctx, req.Code).Result()
	if err != nil {
		if err == redis.Nil {
			return &auth.InfoResponse{
				Message: "Email or password is incorrect",
				Success: false,
			}, nil
		}
		return nil, err
	}

	if storedEmail != req.Email {
		return &auth.InfoResponse{
			Message: "Email or password is incorrect",
			Success: false,
		}, nil
	}

	hashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE users
		SET password = $1
		WHERE id = $2
	`
	_, err = stg.db.Exec(query, hashedPassword, req.Id)
	if err != nil {
		return nil, err
	}

	if err := stg.rdb.Del(ctx, req.Code).Err(); err != nil {
		return nil, err
	}

	return &auth.InfoResponse{
		Message: "Password reset successfully",
		Success: true,
	}, nil
}

func (stg *AuthStorage) ChangeEmail(req *auth.ChangeEmailRequest) (*auth.InfoResponse, error) {
	ctx := context.Background()
	err := stg.rdb.Set(ctx, req.Code, req.NewEmail, time.Minute*15).Err()
	if err != nil {
		return &auth.InfoResponse{
			Message: "Failed to store verification code",
			Success: false,
		}, err
	}

	return &auth.InfoResponse{
		Message: "Verification code sent successfully, please check your email",
		Success: true,
	}, nil
}

func (stg *AuthStorage) VerifyNewEmail(req *auth.VerifyNewEmailRequest) (*auth.TokenResponse, error) {
	ctx := context.Background()

	storedEmail, err := stg.rdb.Get(ctx, req.Code).Result()
	if err != nil {
		if err == redis.Nil {
			return &auth.TokenResponse{
				Id:          req.Id,
				AccessToken: "",
				ExpiresAt:   "",
			}, fmt.Errorf("verification code is invalid or expired")
		}
		return nil, err
	}

	if storedEmail != req.NewEmail {
		return &auth.TokenResponse{
			Id:          req.Id,
			AccessToken: "",
			ExpiresAt:   "",
		}, fmt.Errorf("email does not match")
	}

	query := "UPDATE users SET email = $1 WHERE id = $2"
	_, err = stg.db.ExecContext(ctx, query, req.NewEmail, req.Id)
	if err != nil {
		return nil, err
	}

	if err := stg.rdb.Del(ctx, req.Code).Err(); err != nil {
		return nil, fmt.Errorf("failed to delete verification code from Redis: %v", err)
	}

	var role string
	roleQuery := "SELECT role FROM users WHERE id = $1"
	err = stg.db.QueryRowContext(ctx, roleQuery, req.Id).Scan(&role)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user role from database: %v", err)
	}

	tokenResponse := token.GenerateJWTToken(req.Id, req.NewEmail, role)

	return &auth.TokenResponse{
		Id:          req.Id,
		AccessToken: tokenResponse.AccessToken,
		ExpiresAt:   time.Now().Add(1 * time.Hour).Format(time.RFC3339),
	}, nil
}

func (stg *AuthStorage) CheckEmailAndPassword(req *auth.CheckEmailAndPasswordRequest) (*auth.InfoResponse, error) {
	var storedHash string
	query := `
		SELECT password
		FROM users
		WHERE email = $1
	`
	err := stg.db.QueryRow(query, req.Email).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return &auth.InfoResponse{
				Message: "Email or password is incorrect",
				Success: false,
			}, nil
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
	if err != nil {
		return &auth.InfoResponse{
			Message: "Email or password is incorrect",
			Success: false,
		}, nil
	}

	return &auth.InfoResponse{
		Message: "Email and password are correct",
		Success: true,
	}, nil
}
