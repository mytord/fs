package middleware

import (
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const HeaderUserId = "X-User-Id"
const HeaderSetUserId = "X-Set-User-Id"
const HeaderTokenExpires = "X-Token-Expires"
const HeaderSetToken = "X-Set-Token"

type AuthCredentialsClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type Auth struct {
	jwtSecret     string
	tokenLifetime time.Duration
}

func NewAuth(jwtSecret string, tokenLifetime time.Duration) *Auth {
	return &Auth{
		jwtSecret:     jwtSecret,
		tokenLifetime: tokenLifetime,
	}
}

func (auth *Auth) ApplyToken() ResponseMiddleware {
	return func(resp *http.Response) error {
		userIdHeader := resp.Header.Get(HeaderSetUserId)

		if userIdHeader == "" {
			return nil
		}

		userId, err := strconv.Atoi(userIdHeader)

		if err != nil {
			log.Printf("failed to parse userId: %s", err)
			return nil
		}

		expiresAt := time.Now().Add(auth.tokenLifetime)

		claims := &AuthCredentialsClaims{
			userId,
			jwt.StandardClaims{
				ExpiresAt: expiresAt.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(auth.jwtSecret))

		if err != nil {
			log.Printf("failed to create signed string for token: %s", err)
			return nil
		}

		resp.Header.Del(HeaderSetUserId)
		resp.Header.Set(HeaderSetToken, tokenString)
		resp.Header.Set(HeaderTokenExpires, expiresAt.UTC().Format(time.RFC1123))

		log.Printf("token `%s` was applied to response", tokenString)

		return nil
	}
}

func (auth *Auth) Authorize(inner RequestMiddleware) RequestMiddleware {
	return func(req *http.Request) {
		inner(req)

		req.Header.Del(HeaderUserId)

		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" {
			return
		}

		splitToken := strings.Split(tokenHeader, "Bearer ")
		token := splitToken[1]

		parsed, err := jwt.ParseWithClaims(
			token,
			&AuthCredentialsClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.jwtSecret), nil
			},
		)

		if err != nil {
			log.Printf("failed to parse token: %s (%s)", err, token)
			return
		}

		claims, ok := parsed.Claims.(*AuthCredentialsClaims)

		if !ok || !parsed.Valid {
			log.Print("invalid token")
			return
		}

		req.Header.Set(HeaderUserId, strconv.Itoa(claims.UserId))

		log.Printf("request was authorized by user %d", claims.UserId)
	}
}
