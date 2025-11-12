package eventrequest

type CreateEventRequest struct {
	NameEvent   string   `form:"name_event" json:"name_event"`
	Description string   `form:"description" json:"description"`
	Status      int   	 `form:"status" json:"status"`
	Location    string   `form:"location" json:"location"`
	ImageIDs    []string `form:"image_ids" json:"image_ids"`
}

type UpdateEventRequest struct {
	NameEvent   string   `form:"name_event" json:"name_event"`
	Description string   `form:"description" json:"description"`
	Status      int   	 `form:"status" json:"status"`
	Location    string   `form:"location" json:"location"`
	ImageIDs    []string `form:"image_ids" json:"image_ids"`
}