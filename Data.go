package atinternet

import (
	errortools "github.com/leapforce-libraries/go_errortools"
)

type RowCounts struct {
	RowCounts []struct {
		RowCount int `json:"RowCount"`
	} `json:"RowCounts"`
}

type Contact struct {
	ATInternetID         int32  `json:"ATInternetId"`
	InsightlyID          int32  `json:"InsightlyId"`
	Email                string `json:"Email"`
	Cellphone            string `json:"Cellphone"`
	Phone                string `json:"Phone"`
	Manual               bool   `json:"Manual"`
	MainContact          bool   `json:"MainContact"`
	MainContactCreditor  bool   `json:"MainContactCreditor"`
	MainContactDebtor    bool   `json:"MainContactDebtor"`
	FunctionName         string `json:"FunctionName"`
	EmploymentTerminated bool   `json:"EmploymentTerminated"`
	OrganizationID       int32  `json:"OrganizationId"`
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
	Columns    []string          `json:"columns"`
	Sort       *[]string         `json:"sort,omitempty"`
	Filter     *FilterSet        `json:"filter,omitempty"`
	Space      Space             `json:"space"`
	Period     map[string]Period `json:"period"`
	MaxResults *int              `json:"max-results"`
	PageNum    *int              `json:"page-num"`
	Options    *Options          `json:"options,omitempty"`
}

func (ai *ATInternet) GetData(params *GetDataParams) (*Contact, *errortools.Error) {

	contact := Contact{}
	_, _, e := ai.Post("getData", params, &contact)

	return &contact, e
}

type GetRowCountParams struct {
	Columns []string          `json:"columns"`
	Filter  *FilterSet        `json:"filter,omitempty"`
	Space   Space             `json:"space"`
	Period  map[string]Period `json:"period"`
	Options *Options          `json:"options,omitempty"`
}

func (ai *ATInternet) GetRowCount(params *GetRowCountParams) (*RowCounts, *errortools.Error) {

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getRowCount", params, &rowCounts)

	return &rowCounts, e
}

type GetTotalParams GetRowCountParams

func (ai *ATInternet) GetTotal(params *GetTotalParams) (*RowCounts, *errortools.Error) {

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getTotal", params, &rowCounts)

	return &rowCounts, e
}
