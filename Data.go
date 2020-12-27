package atinternet

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

// generic structs
//
type params struct {
	Columns    []string            `json:"columns"`
	Sort       *[]string           `json:"sort,omitempty"`
	Filter     *FilterSet          `json:"filter,omitempty"`
	Space      Space               `json:"space"`
	Period     map[string][]Period `json:"period"`
	MaxResults *int                `json:"max-results,omitempty"`
	PageNum    *int                `json:"page-num,omitempty"`
	Options    *Options            `json:"options,omitempty"`
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

type SortOrder string

const SortOrderAscending SortOrder = "ASC"
const SortOrderDescending SortOrder = "DESC"

type Sort struct {
	sort []string
}

func (sort *Sort) AddMetric(metric Metric, sortOrder SortOrder) {
	s := string(metric)

	if sortOrder == SortOrderDescending {
		s = fmt.Sprintf("-%s", s)
	}

	sort.sort = append(sort.sort, s)
}

func (sort *Sort) AddProperty(property Property, sortOrder SortOrder) {
	s := string(property)

	if sortOrder == SortOrderDescending {
		s = fmt.Sprintf("-%s", s)
	}

	sort.sort = append(sort.sort, s)
}

// GetData
//
type GetDataParams struct {
	Properties Properties
	Metrics    Metrics
	Sort       *Sort
	Filter     *FilterSet
	Space      []int
	Period     []Period
	MaxResults *int
	PageNum    *int
	Options    *Options
}

func (p *GetDataParams) Params() *params {
	if p == nil {
		return nil
	}

	pa := params{}

	pa.Columns = append(pa.Columns, p.Properties.String()...)
	pa.Columns = append(pa.Columns, p.Metrics.String()...)
	pa.Filter = p.Filter
	pa.Space = Space{p.Space}
	pa.Options = p.Options
	pa.Period = make(map[string][]Period)
	pa.MaxResults = p.MaxResults
	pa.PageNum = p.PageNum
	pa.Options = p.Options

	if p.Sort != nil {
		pa.Sort = &((*p.Sort).sort)
	}

	for i, period := range p.Period {
		pa.Period[fmt.Sprintf("p%v", i+1)] = []Period{period}
	}

	return &pa
}

type Data struct {
	DataFeed DataFeed `json:"DataFeed"`
}

type DataFeed struct {
	Columns []Column                 `json:"Columns"`
	Rows    []map[string]interface{} `json:"Rows"`
	Context Context                  `json:"Context"`
}

type Column struct {
	Category     string `json:"Category"`
	Name         string `json:"Name"`
	Type         string `json:"Type"`
	CustomerType string `json:"CustomerType"`
	Label        string `json:"Label"`
	Description  string `json:"Description"`
	Filterable   bool   `json:"Filterable"`
}

type Context struct {
	Periods []struct {
		Value string `json:"Value"`
	} `json:"Periods"`
}

func (ai *ATInternet) GetData(params *GetDataParams) (*DataFeed, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	data := Data{}
	_, _, e := ai.Post("getData", params.Params(), &data)

	return &data.DataFeed, e
}

// GetRowCount
//
type GetRowCountParams struct {
	Properties Properties
	Metrics    Metrics
	Filter     *FilterSet
	Space      []int
	Period     []Period
	Options    *Options
}

func (p *GetRowCountParams) Params() *params {
	if p == nil {
		return nil
	}

	pa := params{}

	pa.Columns = append(pa.Columns, p.Properties.String()...)
	pa.Columns = append(pa.Columns, p.Metrics.String()...)
	pa.Filter = p.Filter
	pa.Space = Space{p.Space}
	pa.Options = p.Options
	pa.Period = make(map[string][]Period)

	for i, period := range p.Period {
		pa.Period[fmt.Sprintf("p%v", i+1)] = []Period{period}
	}

	return &pa
}

type RowCounts struct {
	RowCounts []struct {
		RowCount int `json:"RowCount"`
	} `json:"RowCounts"`
}

func (ai *ATInternet) GetRowCount(params *GetRowCountParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getRowCount", params.Params(), &rowCounts)

	return &rowCounts, e
}

// GetTotal
//
type GetTotalParams GetRowCountParams

func (p *GetTotalParams) Params() *params {
	if p == nil {
		return nil
	}

	p1 := GetRowCountParams(*p)
	return p1.Params()
}

func (ai *ATInternet) GetTotal(params *GetTotalParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	rowCounts := RowCounts{}
	_, _, e := ai.Post("getTotal", params.Params(), &rowCounts)

	return &rowCounts, e
}
