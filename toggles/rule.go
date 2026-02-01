package toggles

// Comparators 支持的所有匹配符
var Comparators = map[string]func(value1, value2 any) bool{
	"equals": Equals,
	"=":      Equals,
}

// Rule 规则
type Rule struct {
	Id        string `json:"id"`
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

func (r *Rule) GetReferenceValue() any {
	return r.Value
}

func (r *Rule) GetActualValue(m map[string]any) any {
	return m[r.Attribute]
}

func (r *Rule) Compare(m map[string]any) bool {
	return Comparators[r.Operator](r.GetReferenceValue(), r.GetActualValue(m))
}

func Equals(value1, value2 any) bool {
	return value1 == value2
}
