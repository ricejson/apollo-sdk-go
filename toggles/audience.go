package toggles

// Audience 人群
type Audience struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Rules []*Rule `json:"rules"`
}
