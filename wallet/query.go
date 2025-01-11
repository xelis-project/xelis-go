package wallet

import "github.com/xelis-project/xelis-go-sdk/extra_data"

type QueryNumber struct {
	Greater        uint `json:"greater"`
	GreaterOrEqual uint `json:"greater_or_equal"`
	Lesser         uint `json:"lesser"`
	LesserOrEqual  uint `json:"lesser_or_equal"`
}

type QueryHasKey struct {
	Key   interface{} `json:"key"`
	Query *Query      `json:"query"`
}

type QueryAtKey struct {
	Key   interface{} `json:"key"`
	Query *Query      `json:"query"`
}

type QueryPosition struct {
	Position uint   `json:"position"`
	Query    *Query `json:"query"`
}

type QueryElement struct {
	HasKey          QueryHasKey            `json:"has_key"`
	AtKey           QueryAtKey             `json:"at_key"`
	Len             QueryNumber            `json:"len"`
	ContainsElement extra_data.Element     `json:"contains_element"`
	AtPosition      QueryPosition          `json:"at_position"`
	EType           extra_data.ElementType `json:"type"`
}

type QueryValue struct {
	Equal         extra_data.Value     `json:"equal"`
	StartsWith    extra_data.Value     `json:"starts_with"`
	EndsWith      extra_data.Value     `json:"ends_with"`
	ContainsValue extra_data.Value     `json:"contains_value"`
	IsOfType      extra_data.ValueType `json:"is_of_type"`
	Matches       string               `json:"matches"`
	NumberOp      QueryNumber          `json:"number_op"`
}

type Query struct {
	Not     *Query       `json:"not,omitempty"`
	And     []Query      `json:"and"`
	Or      []Query      `json:"or"`
	Element QueryElement `json:"element,omitempty"`
	Value   QueryValue   `json:"value"`
}
