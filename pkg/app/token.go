package app

import (
	"TicketSales/global"
	"TicketSales/pkg/cache"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func GenerateToken(userid uint64) (*TokenDetails, error) {
	var err error
	//Creating Access Token
	td := &TokenDetails{}
	td.AtExpires = time.Now().Local().Unix() + global.JWTSetting.AtExpires
	td.AccessUuid = uuid.New().String()
	td.RtExpires = time.Now().Local().Unix() + global.JWTSetting.RtExpires
	td.RefreshUuid = uuid.New().String()
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["refresh_uuid"] = td.RefreshUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(global.JWTSetting.AccessSecret))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(global.JWTSetting.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now().Local()
	errAccess := cache.Store.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now).Milliseconds()/1000)
	if errAccess != nil {
		return errAccess
	}
	errRefresh := cache.Store.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now).Milliseconds()/1000)
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	return token
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.JWTSetting.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

type AccessDetails struct {
	AccessUuid  string
	RefreshUuid string
	UserId      uint64
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid:  accessUuid,
			RefreshUuid: refreshUuid,
			UserId:      userId,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, flag := cache.Store.Get(authD.AccessUuid)
	if !flag {
		return 0, errors.New("auth is expired")
	}
	userID, _ := strconv.ParseUint(userid.(string), 10, 64)
	return userID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	n, err := cache.Store.Delete(givenUuid)
	if err != nil {
		return 0, err
	}
	return n, nil
}
