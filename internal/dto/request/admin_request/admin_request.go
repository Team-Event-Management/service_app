package adminrequest

type RegisterAdminRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginAdminRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name"`
}