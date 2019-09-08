package dto

import "fmt"

type AddTopicDTO struct {
	TopicName string `json:"topic_name"`
}

func (dto AddTopicDTO) Validate() error {
	if dto.TopicName == "" {
		return fmt.Errorf("TopicName cannot be empty")
	}
	return nil
}
