package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

type contextKey string

const RoleKey contextKey = "Role"

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type PVZListResponse struct {
	PVZ       PVZ            `json:"pvz"`
	Reception []GetReception `json:"receptions"`
}

type GetReception struct {
	ID       string       `json:"id"`
	DateTime string       `json:"dateTime"`
	PVZID    string       `json:"pvzId"`
	Status   string       `json:"status"`
	Products []GetProduct `json:"products"`
}

type GetProduct struct {
	ID       string `json:"id"`
	DateTime string `json:"dateTime"`
	Type     string `json:"type"`
}
