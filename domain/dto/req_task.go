package dto

type ReqTask struct {
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Status      string   `json:"status" bson:"status"`
	Priority    string   `json:"priority" bson:"priority"`
	Tags        []string `json:"tags" bson:"tags"`
}
