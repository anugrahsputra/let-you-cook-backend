package model

type Task struct {
	Id          string   `json:"id" bson:"id"`
	UserId      string   `json:"user_id" bson:"user_id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Status      string   `json:"status" bson:"status"`
	Priority    string   `json:"priority" bson:"priority"`
	CreatedAt   int      `json:"created_at" bson:"created_at"`
	UpdatedAt   int      `json:"updated_at" bson:"updated_at"`
	CompletedAt int      `json:"completed_at" bson:"completed_at"`
	Tags        []string `json:"tags" bson:"tags"`
}
