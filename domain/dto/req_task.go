package dto

type ReqTask struct {
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	CategoryId  string   `json:"category_id" bson:"category_id"`
	Status      string   `json:"status" bson:"status"`
	Priority    string   `json:"priority" bson:"priority"`
	Tags        []string `json:"tags" bson:"tags"`
}

type ReqPatchTask struct {
	Title       *string   `json:"title" bson:"title" binding:"omitempty"`
	Description *string   `json:"description" bson:"description" binding:"omitempty"`
	CategoryId  *string   `json:"category_id" bson:"category_id" binding:"omitempty"`
	Status      *string   `json:"status" bson:"status" binding:"omitempty"`
	Priority    *string   `json:"priority" bson:"priority" binding:"omitempty"`
	Tags        *[]string `json:"tags" bson:"tags" binding:"omitempty"`
}

type TaskResp struct {
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	CategoryId  string   `json:"category_id" bson:"category_id"`
	Status      string   `json:"status" bson:"status"`
	Priority    string   `json:"priority" bson:"priority"`
	Tags        []string `json:"tags" bson:"tags"`
}

type TaskByCategoryGroupResp struct {
	CategoryId   string `json:"category_id" bson:"category_id"`
	CategoryName string `json:"category_name" bson:"category_name"`
	Tasks        []TaskResp
}
