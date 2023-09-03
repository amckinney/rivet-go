// This file was auto-generated by Fern from our API Definition.

package game

import (
	fmt "fmt"
	uuid "github.com/gofrs/uuid/v5"
	rivetgo "github.com/rivet-gg/rivet-go"
)

type Handle struct {
	GameId      uuid.UUID           `json:"game_id"`
	NameId      rivetgo.Identifier  `json:"name_id"`
	DisplayName rivetgo.DisplayName `json:"display_name"`
	// The URL of this game's logo image.
	LogoUrl *string `json:"logo_url,omitempty"`
	// The URL of this game's banner image.
	BannerUrl *string `json:"banner_url,omitempty"`
}

// A game statistic.
type Stat struct {
	Config *StatConfig `json:"config,omitempty"`
	// A single overall value of the given statistic.
	OverallValue float64 `json:"overall_value"`
}

// A value denoting the aggregation method of a game statistic.
type StatAggregationMethod string

const (
	// Summation aggregation.
	StatAggregationMethodSum StatAggregationMethod = "sum"
	// Average aggregation.
	StatAggregationMethodAverage StatAggregationMethod = "average"
	// Minimum value aggregation.
	StatAggregationMethodMin StatAggregationMethod = "min"
	// Maximum value aggregation.
	StatAggregationMethodMax StatAggregationMethod = "max"
)

func NewStatAggregationMethodFromString(s string) (StatAggregationMethod, error) {
	switch s {
	case "sum":
		return StatAggregationMethodSum, nil
	case "average":
		return StatAggregationMethodAverage, nil
	case "min":
		return StatAggregationMethodMin, nil
	case "max":
		return StatAggregationMethodMax, nil
	}
	var t StatAggregationMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s StatAggregationMethod) Ptr() *StatAggregationMethod {
	return &s
}

// A game statistic config.
type StatConfig struct {
	RecordId    uuid.UUID             `json:"record_id"`
	IconId      uuid.UUID             `json:"icon_id"`
	Format      StatFormatMethod      `json:"format,omitempty"`
	Aggregation StatAggregationMethod `json:"aggregation,omitempty"`
	Sorting     StatSortingMethod     `json:"sorting,omitempty"`
	Priority    int                   `json:"priority"`
	DisplayName rivetgo.DisplayName   `json:"display_name"`
	// A string appended to the end of a singular game statistic's value. Example: 1 **dollar**.
	PostfixSingular *string `json:"postfix_singular,omitempty"`
	// A string appended to the end of a game statistic's value that is not exactly 1. Example: 45 **dollars**.
	PostfixPlural *string `json:"postfix_plural,omitempty"`
	// A string appended to the beginning of a singular game statistic's value. Example: **value** 1.
	PrefixSingular *string `json:"prefix_singular,omitempty"`
	// A string prepended to the beginning of a game statistic's value that is not exactly 1. Example: **values** 45.
	PrefixPlural *string `json:"prefix_plural,omitempty"`
}

// A value denoting the format method of a game statistic.
type StatFormatMethod string

const (
	// An integer with no decimals (1,234).
	StatFormatMethodInteger StatFormatMethod = "integer"
	// A float with 1 decimal (1,234.5).
	StatFormatMethodFloat1 StatFormatMethod = "float_1"
	// A float with 2 decimals (1,234.56).
	StatFormatMethodFloat2 StatFormatMethod = "float_2"
	// A float with 3 decimals (1,234.567).
	StatFormatMethodFloat3 StatFormatMethod = "float_3"
	// A duration with minute precision (1d 2h 45m).
	StatFormatMethodDurationMinute StatFormatMethod = "duration_minute"
	// A duration with second precision (1d 2h 45m 21s).
	StatFormatMethodDurationSecond StatFormatMethod = "duration_second"
	// A duration with hundredth-second precision (1d 2h 45m 21.46s). It is important to notice that this custom format is not a standard way to define Enums in OpenAPI and it will be not understood by the majority of tools that parse OpenAPI files. It is important to check with the tools or libraries that you are using to make sure they support this custom format.
	StatFormatMethodDurationHundredthSecond StatFormatMethod = "duration_hundredth_second"
)

func NewStatFormatMethodFromString(s string) (StatFormatMethod, error) {
	switch s {
	case "integer":
		return StatFormatMethodInteger, nil
	case "float_1":
		return StatFormatMethodFloat1, nil
	case "float_2":
		return StatFormatMethodFloat2, nil
	case "float_3":
		return StatFormatMethodFloat3, nil
	case "duration_minute":
		return StatFormatMethodDurationMinute, nil
	case "duration_second":
		return StatFormatMethodDurationSecond, nil
	case "duration_hundredth_second":
		return StatFormatMethodDurationHundredthSecond, nil
	}
	var t StatFormatMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s StatFormatMethod) Ptr() *StatFormatMethod {
	return &s
}

// A value denoting the sorting method of a game statistic.
type StatSortingMethod string

const (
	// Descending sorting.
	StatSortingMethodDesc StatSortingMethod = "desc"
	// Ascending sorting.
	StatSortingMethodAsc StatSortingMethod = "asc"
)

func NewStatSortingMethodFromString(s string) (StatSortingMethod, error) {
	switch s {
	case "desc":
		return StatSortingMethodDesc, nil
	case "asc":
		return StatSortingMethodAsc, nil
	}
	var t StatSortingMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s StatSortingMethod) Ptr() *StatSortingMethod {
	return &s
}

// A game statistic summary.
type StatSummary struct {
	Game  *Handle `json:"game,omitempty"`
	Stats []*Stat `json:"stats,omitempty"`
}
