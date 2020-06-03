package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
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

func CreateToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}

	accessToken, err := createAccessToken(userid, td)
	refreshToken, err := createRefreshToken(userid, td)

	if err != nil {
		return nil, err
	}

	td.AccessToken = accessToken
	td.RefreshToken = refreshToken

	return td, nil
}

func createRefreshToken(userid uint64, td *TokenDetails) (string, error) {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return rt.SignedString([]byte("REFRESH_SECRET"))
}

func createAccessToken(userid uint64, td *TokenDetails) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte("ACCESS_SECRET"))
}
