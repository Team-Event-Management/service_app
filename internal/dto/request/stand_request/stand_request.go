package standrequest

import "github.com/google/uuid"

type CreateStandRequest struct {
	StandName       string    `form:"stand_name" json:"stand_name"`
	Lat             float64   `form:"lat" json:"lat"`
	Lng             float64   `form:"lng" json:"lng"`
	Address         string    `form:"address" json:"address"`
	StandCategoryID uuid.UUID `form:"stand_category_id" json:"stand_category_id"`
}

type UpdateStandRequest struct {
	StandName       string    `form:"stand_name" json:"stand_name"`
	Lat             float64   `form:"lat" json:"lat"`
	Lng             float64   `form:"lng" json:"lng"`
	Address         string    `form:"address" json:"address"`
	StandCategoryID uuid.UUID `form:"stand_category_id" json:"stand_category_id"`
}
