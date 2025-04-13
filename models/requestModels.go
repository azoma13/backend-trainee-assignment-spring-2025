package models

type DummyLoginRequest struct {
	Role string `json:"role"`
}

type PVZRequest struct {
	City string `json:"city"`
}

type ReceptionsRequest struct {
	PVZID string `json:"pvzId"`
}

type ProductsRequest struct {
	Type  string `json:"type"`
	PVZID string `json:"pvzId"`
}

type PVZListRequest struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}
