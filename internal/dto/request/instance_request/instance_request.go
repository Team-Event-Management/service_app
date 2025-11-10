package instancerequest

type CreateInstanceRequest struct {
	Name        string  `form:"name" json:"name"`
	Lat         float64 `form:"lat" json:"lat"`
	Lng         float64 `form:"lng" json:"lng"`
	FullAddress string  `form:"full_address" json:"full_address"`
}

type UpdateInstanceRequest struct {
	Name        string  `form:"name" json:"name"`
	Lat         float64 `form:"lat" json:"lat"`
	Lng         float64 `form:"lng" json:"lng"`
	FullAddress string  `form:"full_address" json:"full_address"`
}