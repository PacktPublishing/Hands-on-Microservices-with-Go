package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/repositories"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//Must be at least 256 bytes long
const hmacSampleSecret = "R541QRVVTFFGZ2APJDHFSBBF9DMO6XU9PMBQ3C4CPDH4EII86PC9U5DVTFELM3VNK0OWYLRIDM7ROXGCGF84KVCQPNQK71BJC5PEAL4K7CU8XW8AMVKQL0X33HGOF49FLDA8DR2HEDEG4PMZ2RCK3WVI3LCMT7SM5WSEUW7C1R56NHDOHGN8LR7RG0J02KN178PLVVPM5SI84LZ371VDX24ER7SQWNRXWLMKYY5AJCS0YQ91HKB8CT13S1PG7R89IDPHYCW7CPMZGYQSCWAO9J9VJK5CF1C7MEEBH2SL7CGBNZHSCTRPUSI9U0S0N8IZG3I36QWTJ7KRTMECGOVC6WAVVPU9OW7BTR1XYJ3Y43RDG7Q38E831DK9HOS9X8ZEF1LVDNLT1JUJLN3PGRJN5EHLLMGNC2BYAXE7A9QDPHK6U9KMYRSBFBKOEAA6UDTPVMX41ZWOVW5JI2B0GZG2A51IO7OS0I8SW7RCAO8H01TJRR72M6AUAMPCLFU1ZZRW"

type Handler struct {
	SessionsRepo *repositories.RestSessionsRepository
	UsersRepo    *repositories.RestUsersRepository
}

func (h *Handler) Authorize(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing Username or Password")
		return
	}

	//Retrieve the user from the Users Repository
	user, err := h.UsersRepo.GetUserByUsername(username)
	if err == repositories.Err404OnUserRequest {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	//Verify password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}

	now := time.Now()
	threeMonths := now.AddDate(0, 3, 0)

	jtiRaw := make([]byte, 128)
	_, err = rand.Read(jtiRaw)
	jti := replaceSlashesAndPlus(base64.StdEncoding.EncodeToString(jtiRaw))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	claims := jwt.MapClaims{
		"iss": "packt",
		"sub": username,
		"exp": threeMonths.Unix(),
		"iat": now.Unix(),
		"jti": jti,
		//Private
		"Hello": "World", //This is just an example
	}

	//User is Valid, create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	//Sign JWT
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)

}

type ctxKey int

const (
	Username ctxKey = iota
	UserID
	FirstName
	LastName
	Jti
)

func VerifyJWT(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: No Authorization Header")
			return
		}

		bearerIndex := strings.Index(authHeader, "Bearer")
		if bearerIndex < 0 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: No Bearer on Authorization Header")
			return
		}

		jwtString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret -- Our MacSecret
			return []byte(hmacSampleSecret), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: Wrong JWT alg.")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: Wrong JWT claims.")
			return
		}

		ok = claims.VerifyExpiresAt(time.Now().Unix(), true)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: JWT has expired.")
			return
		}
		ok = claims.VerifyIssuer("packt", true)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: JWT has wrong issuers.")
			return
		}
		username := claims["sub"].(string) //Type assertion neeeded because it is of type interface{}
		if username == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: JWT does not have a valid sub.")
			return
		}
		jti := claims["jti"].(string)
		if jti == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized: JWT does not have a valid jti.")
			return
		}

		//JWT was valid.
		//Add the jti to the context of the Request
		ctx := r.Context()
		ctx = context.WithValue(ctx, Jti, jti)
		ctx = context.WithValue(ctx, Username, username)

		//Continue through the middleware chain
		r = r.WithContext(ctx)
		f.ServeHTTP(w, r)
	}
}

func (h *Handler) AddSessionData(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		jti := r.Context().Value(Jti).(string)
		username := r.Context().Value(Username).(string)
		if jti == "" || username == "" {
			//Something was wrong on the middleware chain
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		//Get Session
		session, err := h.SessionsRepo.GetSession(jti)
		if err == repositories.Err404OnSessionRequest {
			//Create session
			//Get User from User Repo
			user, err := h.UsersRepo.GetUserByUsername(username)
			//VERIFY TYPE OF ERROR
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Unauthorized: JWT does not have a valid username.")
				return
			}
			session = &entities.Session{
				UserID:    user.ID,
				Username:  user.Username,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
			err = h.SessionsRepo.SetSession(jti, session)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, err.Error())
				return
			}
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		//Add the data from the session to the request
		ctx := context.Background()
		ctx = context.WithValue(ctx, FirstName, session.FirstName)
		ctx = context.WithValue(ctx, LastName, session.LastName)

		r = r.WithContext(ctx)

		//Continue the middleware chain
		f.ServeHTTP(w, r)
	}
}

func (h *Handler) Restricted1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ctx := r.Context()
	firstName := ctx.Value(FirstName)
	lastName := ctx.Value(LastName)
	fmt.Fprintln(w, "You have reached Restricted Resource 1")
	fmt.Fprintf(w, "Hello: %s %s\n", firstName, lastName)
}

func (h *Handler) Restricted2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ctx := r.Context()
	firstName := ctx.Value(FirstName)
	lastName := ctx.Value(LastName)
	fmt.Fprintln(w, "You have reached Restricted Resource 2")
	fmt.Fprintf(w, "Hello: %s %s\n", firstName, lastName)
}

func replaceSlashesAndPlus(str string) string {
	str = strings.Replace(str, "/", "0", -1)
	str = strings.Replace(str, "+", "0", -1)
	return str
}
