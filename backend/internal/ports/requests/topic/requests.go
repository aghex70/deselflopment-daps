package requests

type CreateTopicRequest struct {
	Name string `json:"name" validate:"required"`
}

type DeleteTopicRequest struct {
	TopicID uint `json:"topic_id"`
}

type GetTopicRequest struct {
	TopicID uint `json:"topic_id"`
}

type UpdateTopicRequest struct {
	Name    string `json:"name"`
	TopicID uint   `json:"topic_id"`
}
