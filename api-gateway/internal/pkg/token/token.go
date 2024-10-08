package tokens

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	pb "flashSale_gateway/internal/pkg/genproto"
)
const signingKey = "secret_key"
func VerifyToken(tokenString string) error {
	secretKey := []byte("secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ExtractClaims(tokenStr, key string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	log.Println("token : ", err)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
func GenerateJWTToken(user *pb.User) (string, string) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(180 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token: ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	claims["email"] = user.Email
	claims["role"] = user.Role
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(48 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token: ", err)
	}

	return access, refresh
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
