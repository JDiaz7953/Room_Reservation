package handlers

import ("fmt"
		"github.com/golang-jwt/jwt/v5"
		"time"
)


var secretKey = []byte("secret-key")

//Creates token
func CreateToken(userID uint, username string) (string, error){
	
	//creates token with hashing method and user related payload 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
	jwt.MapClaims{
		"user_id": userID,
		"username" : username,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	//signes user token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//Verifies Token
func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
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
