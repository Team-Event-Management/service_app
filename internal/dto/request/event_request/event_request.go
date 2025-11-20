package eventrequest

import "mime/multipart"

type CreateEventRequest struct {
	NameEvent   string    				`form:"name_event" json:"name_event"`
	Description string    				`form:"description" json:"description"`
	Status      int       				`form:"status" json:"status"`
	Location    string    				`form:"location" json:"location"`
	StartDate   string    				`form:"start_date"`
	EventImages []*multipart.FileHeader `form:"event_images" json:"event_images"`
}

type UpdateEventRequest struct {
	NameEvent   string   `form:"name_event" json:"name_event"`
	Description string   `form:"description" json:"description"`
	Status      int      `form:"status" json:"status"`
	Location    string   `form:"location" json:"location"`
	StartDate 	string 	 `form:"start_date" json:"start_date"`
}