package toggles

// Toggle 开关
type Toggle struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	CreateAt    int64       `json:"createAt"`
	UpdateAt    int64       `json:"updateAt"`
	Audiences   []*Audience `json:"audiences"`
}

func (t *Toggle) Init(key string) error {
	return nil
}
