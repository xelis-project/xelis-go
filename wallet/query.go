package wallet

import "github.com/xelis-project/xelis-go-sdk/extra_data"

type QueryNumber struct {
	Greater        uint `json:"greater,omitempty"`
	GreaterOrEqual uint `json:"greater_or_equal,omitempty"`
	Lesser         uint `json:"lesser,omitempty"`
	LesserOrEqual  uint `json:"lesser_or_equal,omitempty"`
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
	HasKey          *QueryHasKey           `json:"has_key,omitempty"`
	AtKey           *QueryAtKey            `json:"at_key,omitempty"`
	Len             QueryNumber            `json:"len,omitempty"`
	ContainsElement extra_data.Element     `json:"contains_element,omitempty"`
	AtPosition      *QueryPosition         `json:"at_position,omitempty"`
	ElementType     extra_data.ElementType `json:"type,omitempty"`
}

type QueryValue struct {
	Equal         interface{}          `json:"equal,omitempty"`
	StartsWith    interface{}          `json:"starts_with,omitempty"`
	EndsWith      interface{}          `json:"ends_with,omitempty"`
	ContainsValue interface{}          `json:"contains_value,omitempty"`
	IsOfType      extra_data.ValueType `json:"is_of_type,omitempty"`
	Matches       string               `json:"matches,omitempty"`
	*QueryNumber
}

type Query struct {
	Not *Query  `json:"not,omitempty"`
	And []Query `json:"and,omitempty"`
	Or  []Query `json:"or,omitempty"`
	*QueryElement
	*QueryValue
}
