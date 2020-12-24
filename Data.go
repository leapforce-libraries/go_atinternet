package atinternet

import (
	errortools "github.com/leapforce-libraries/go_errortools"
)

type Data struct {
}

type RowCounts struct {
	RowCounts []struct {
		RowCount int `json:"RowCount"`
	} `json:"RowCounts"`
}

type FilterSet struct {
	Metric   *Filters `json:"metric,omitempty"`
	Property *Filters `json:"property,omitempty"`
}

type Filters struct {
	Filter    *Filter
	AndFilter *struct {
		And []Filters `json:"$AND"`
	}
	OrFilter *struct {
		Or []Filters `json:"$OR"`
	}
}

type Filter struct {
	Filter map[string]map[string]interface{}
}

type Space struct {
	S []int `json:"s"`
}

type Period struct {
	Type  string `json:"type"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Options struct {
	IgnoreNullProperties bool `json:"ignore_null_properties"`
}

type GetDataParams struct {
	Properties Properties
	Metrics    Metrics
	columns    []string            `json:"columns"`
	Sort       *[]string           `json:"sort,omitempty"`
	Filter     *FilterSet          `json:"filter,omitempty"`
	Space      Space               `json:"space"`
	Period     map[string][]Period `json:"period"`
	MaxResults *int                `json:"max-results"`
	PageNum    *int                `json:"page-num"`
	Options    *Options            `json:"options,omitempty"`
}

func (ai *ATInternet) GetData(params *GetDataParams) (*Data, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	data := Data{}
	_, _, e := ai.Post("getData", params, &data)

	return &data, e
}

type GetRowCountParams struct {
	Properties Properties
	Metrics    Metrics
	columns    []string            `json:"columns"`
	Filter     *FilterSet          `json:"filter,omitempty"`
	Space      Space               `json:"space"`
	Period     map[string][]Period `json:"period"`
	Options    *Options            `json:"options,omitempty"`
}

func (ai *ATInternet) GetRowCount(params *GetRowCountParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	(*params).columns = append((*params).columns, params.Properties.String()...)

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getRowCount", params, &rowCounts)

	return &rowCounts, e
}

type GetTotalParams GetRowCountParams

func (ai *ATInternet) GetTotal(params *GetTotalParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getTotal", params, &rowCounts)

	return &rowCounts, e
}
