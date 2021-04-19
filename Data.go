package atinternet

import (
	"fmt"
	"net/url"
	"strings"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// generic structs
//
type params struct {
	Columns    []string            `json:"columns"`
	Sort       *[]string           `json:"sort,omitempty"`
	Filter     *FilterSet          `json:"filter,omitempty"`
	Space      Space               `json:"space"`
	Period     map[string][]period `json:"period"`
	MaxResults *int64              `json:"max-results,omitempty"`
	PageNum    *int64              `json:"page-num,omitempty"`
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
	Sites      *[]int64    `json:"s,omitempty"`
	Level2Site *Level2Site `json:"l2s,omitempty"`
}

type Level2Site struct {
	SiteID       int64 `json:"s"`
	Level2SiteID int64 `json:"l2"`
}

type period struct {
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

type Period struct {
	periods map[string][]period
}

func (p *Period) AddDay(date civil.Date) {
	if p.periods == nil {
		p.periods = make(map[string][]period)
	}

	p.periods[fmt.Sprintf("p%v", len(p.periods)+1)] = []period{period{
		Type:  "D",
		Start: date.String(),
		End:   date.String(),
	}}
}

// GetData
//
type GetDataParams struct {
	Properties Properties
	Metrics    Metrics
	Sort       *Sort
	Filter     *FilterSet
	Space      Space
	Period     Period
	MaxResults *int64
	PageNum    *int64
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
	pa.Space = p.Space
	pa.Options = p.Options
	pa.Period = p.Period.periods
	pa.MaxResults = p.MaxResults
	pa.PageNum = p.PageNum
	pa.Options = p.Options

	if p.Sort != nil {
		pa.Sort = &((*p.Sort).sort)
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

func (service *Service) GetData2(params *GetDataParams) (*DataFeed, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	values := url.Values{}
	columns := []string{}
	if params.MaxResults != nil {
		values.Set("max-results", fmt.Sprintf("%v", *params.MaxResults))
	}
	for _, metric := range params.Metrics.String() {
		columns = append(columns, metric)
	}
	if params.PageNum != nil {
		values.Set("page-num", fmt.Sprintf("%v", *params.PageNum))
	}
	for _, period := range params.Period.periods {
		periods := []string{}
		for _, p := range period {
			periods = append(periods, fmt.Sprintf("%s:{start:'%s',end:'%s'}", p.Type, p.Start, p.End))
		}
		values.Set("period", fmt.Sprintf("{%s}", strings.Join(periods, ",")))
	}
	for _, property := range params.Properties.String() {
		columns = append(columns, property)
	}
	if params.Sort != nil {
		sorts := []string{}
		for _, sort := range params.Sort.sort {
			sorts = append(sorts, sort)
		}
		values.Set("sort", fmt.Sprintf("{%s}", strings.Join(sorts, ",")))
	}
	spaces := []string{}
	if params.Space.Sites != nil {
		if len(*params.Space.Sites) == 1 {
			spaces = append(spaces, fmt.Sprintf("s:%v", (*params.Space.Sites)[0]))
		} else {
			sites := []string{}
			for _, site := range *params.Space.Sites {
				sites = append(sites, fmt.Sprintf("%v", site))
			}
			spaces = append(spaces, fmt.Sprintf("s:{%s}", strings.Join(sites, ",")))
		}
	}
	if params.Space.Level2Site != nil {
		spaces = append(spaces, fmt.Sprintf("l2s:{s:%v,l2:%v}", params.Space.Level2Site.SiteID, params.Space.Level2Site.Level2SiteID))
	}
	values.Set("columns", fmt.Sprintf("{%s}", strings.Join(columns, ",")))
	values.Set("space", fmt.Sprintf("{%s}", strings.Join(spaces, ",")))

	data := Data{}

	//fmt.Println(service.url2(fmt.Sprintf("getData?%s", values.Encode())))

	requestConfig := go_http.RequestConfig{
		URL:           service.url2(fmt.Sprintf("getData?%s", values.Encode())),
		ResponseModel: &data,
	}
	_, _, e := service.post(&requestConfig)

	return &data.DataFeed, e
}

func (service *Service) GetData3(params *GetDataParams) (*DataFeed, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	data := Data{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url3("getData"),
		BodyModel:     params.Params(),
		ResponseModel: &data,
	}
	_, _, e := service.post(&requestConfig)
	//fmt.Println(requestConfig.URL)
	//b, _ := json.Marshal(requestConfig.BodyModel)
	//fmt.Println(string(b))

	return &data.DataFeed, e
}

// GetRowCount
//
type GetRowCountParams struct {
	Properties Properties
	Metrics    Metrics
	Filter     *FilterSet
	Space      Space
	Period     Period
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
	pa.Space = p.Space
	pa.Options = p.Options
	pa.Period = p.Period.periods

	return &pa
}

type RowCounts struct {
	RowCounts []struct {
		RowCount int `json:"RowCount"`
	} `json:"RowCounts"`
}

func (service *Service) GetRowCount3(params *GetRowCountParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	rowCounts := RowCounts{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url3("getRowCount"),
		BodyModel:     params.Params(),
		ResponseModel: &rowCounts,
	}
	_, _, e := service.post(&requestConfig)

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

func (service *Service) GetTotal3(params *GetTotalParams) (*RowCounts, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	rowCounts := RowCounts{}

	requestConfig := go_http.RequestConfig{
		URL:           service.url3("getTotal"),
		BodyModel:     params.Params(),
		ResponseModel: &rowCounts,
	}
	_, _, e := service.post(&requestConfig)

	return &rowCounts, e
}
