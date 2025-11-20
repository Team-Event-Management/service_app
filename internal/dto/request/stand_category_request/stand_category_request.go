package standcategoryrequest

type CreateStandCategoryRequest struct {
    Name        string `form:"name" json:"name"`
    Description string `form:"description" json:"description"`
}

type UpdateStandCategoryRequest struct {
    Name        string `form:"name" json:"name"`
    Description string `form:"description" json:"description"`
}
