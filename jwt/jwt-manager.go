package jwt

import (
	"auth-micro/model"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTManager struct {
	secretKey string
	tokenDuration time.Duration
}
type UserClaims struct {
	jwt.StandardClaims
	UserEmail string
}
//? Responseibility is create a new jwt manager
func NewJWTManager(secretKey string, tokenDuration time.Duration) (*JWTManager){
	return &JWTManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

// Gnerating a token
func (jwtmanager *JWTManager) GeneratingToken(user *model.User) (string,error) {
	// 1. Prepare the User Claims
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtmanager.tokenDuration).Unix(),
		},
		UserEmail: user.Email,
	}
	// create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims);
	return token.SignedString([]byte(jwtmanager.secretKey))
}
// ? Responsibility :- decode the token and get the userClaims and return the userEmail. (userClaims/userEmail);
func VerifyToken(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func (token *jwt.Token) (interface{},error) {
			_,ok := token.Method.(*jwt.SigningMethodHMAC);
			
			if !ok {
				return nil, fmt.Errorf("Unexpected Token signing Method")
			}
			return [] byte(os.Getenv("SECRET_KEY")),nil;
		},
	);
	if err != nil {
		return nil, fmt.Errorf("Invalid Token %w", err);
	}
	claims, ok := token.Claims.(*UserClaims);
	if !ok {
		return nil, fmt.Errorf("Invalid Token %w", err);
	}
	return claims, nil
};

func AuthorizeJwtToken() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization");
		if len(tokenString) == 0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{"jwt failure":"Authorization token is not provided."})
			ctx.Abort();
		}
		token := strings.Split(tokenString," ")[1];
		claims,err := VerifyToken(token);

		if err != nil {
			ctx.JSON(http.StatusUnauthorized,gin.H{"jwt failure":"Authorization token is not provided."})
			ctx.Abort();
		}
		ctx.Set("usermail", claims.UserEmail);
		ctx.Next();
	}
}