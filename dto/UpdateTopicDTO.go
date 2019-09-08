package dto

type UpdateTopicDTO struct {
	Offset int64 `json:"offset"`
}

func (dto UpdateTopicDTO) Validate() error {
	return nil
}
