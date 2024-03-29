package atinternet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type GetDataVersion string

const (
	GetDataVersion2 GetDataVersion = "v2"
	GetDataVersion3 GetDataVersion = "v3"
)

type params struct {
	Columns       []string            `json:"columns"`
	Sort          *[]string           `json:"sort,omitempty"`
	Filter        *FilterSet          `json:"filter,omitempty"`
	FilterString2 *string             `json:"-"`
	Space         Space               `json:"space"`
	Period        map[string][]period `json:"period"`
	MaxResults    *int64              `json:"max-results,omitempty"`
	PageNum       *int64              `json:"page-num,omitempty"`
	Options       *Options            `json:"options,omitempty"`
}

type FilterSet struct {
	Metric   *Filters `json:"metric,omitempty"`
	Property *Filters `json:"property,omitempty"`
}

/*func (filters Filters) AddMetricFilter(fieldName string, operator string, value interface{}) {
	if filters == nil {
		_m := make(map[string]interface{})
		filters = _m
	}

	addFilter(filters, fieldName, operator, value)
}

func (filters Filters) AddPropertyFilter(fieldName string, operator string, value interface{}) {
	if filters == nil {
		_m := make(map[string]interface{})
		filters = _m
	}

	addFilter(filters, fieldName, operator, value)
}

func addFilter(m *map[string]map[string]interface{}, fieldName string, operator string, value interface{}) {
	m1 := make(map[string]interface{})
	m1[operator] = value

	(*m)[fieldName] = m1
}*/

/*type Filters struct {
	Filter    *map[string]map[string]interface{} `json:"-"`
	AndFilter *Filters                           `json:"$AND,omitempty"`
	OrFilter  *Filters                           `json:"$OR,omitempty"`
}*/

type Filters map[string]interface{}

func (filters Filters) AddFilter(fieldName string, operator string, value interface{}) {
	if filters == nil {
		m := make(map[string]interface{})
		filters = m
	}

	m1 := make(map[string]interface{})
	m1[operator] = value

	filters[fieldName] = m1
}

func (filters Filters) AddAnd(f *[]Filters) {
	if filters == nil {
		m := make(map[string]interface{})
		filters = m
	}

	filters["$AND"] = f
}

func (filters Filters) AddOr(f *[]Filters) {
	if filters == nil {
		m := make(map[string]interface{})
		filters = m
	}

	filters["$OR"] = f
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
	Type  Granularity `json:"type"`
	Start string      `json:"start"`
	End   string      `json:"end"`
}

type Granularity string

const (
	GranularityYear     Granularity = "Y"
	GranularitySemester Granularity = "S"
	GranularityQuarter  Granularity = "Q"
	GranularityMonth    Granularity = "M"
	GranularityWeek     Granularity = "W"
	GranularityDay      Granularity = "D"
	GranularityHour     Granularity = "H"
	GranularityMinute   Granularity = "MN"
)

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

	p.periods[fmt.Sprintf("p%v", len(p.periods)+1)] = []period{{
		Type:  GranularityDay,
		Start: date.String(),
		End:   date.String(),
	}}
}

func (p *Period) AddDateRange(startDate civil.Date, endDate civil.Date) {
	if p.periods == nil {
		p.periods = make(map[string][]period)
	}

	p.periods[fmt.Sprintf("p%v", len(p.periods)+1)] = []period{{
		Type:  GranularityDay,
		Start: startDate.String(),
		End:   endDate.String(),
	}}
}

// GetData
//
type GetDataParams struct {
	Properties    Properties
	Metrics       Metrics
	Sort          *Sort
	Filter        *FilterSet
	FilterString2 *string
	Space         Space
	Period        Period
	MaxResults    *int64
	PageNum       *int64
	Options       *Options
}

func (p *GetDataParams) Params() *params {
	if p == nil {
		return nil
	}

	pa := params{}

	pa.Columns = append(pa.Columns, p.Properties.String()...)
	pa.Columns = append(pa.Columns, p.Metrics.String()...)
	pa.Filter = p.Filter
	pa.FilterString2 = p.FilterString2
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

type Data2 struct {
	DataFeed []DataFeed `json:"DataFeed"`
}

type DataFeed struct {
	Columns []Column                 `json:"Columns"`
	Rows    []map[string]interface{} `json:"Rows"`
	Context *Context                 `json:"Context"`
}

type Column struct {
	Category     string          `json:"Category"`
	Name         string          `json:"Name"`
	Type         string          `json:"Type"`
	CustomerType string          `json:"CustomerType"`
	Label        string          `json:"Label"`
	Description  string          `json:"Description"`
	Filterable   bool            `json:"Filterable"`
	Pie          *bool           `json:"Pie"`
	Precision    json.RawMessage `json:"Precision"`
	Summable     *bool           `json:"Summable"`
}

type Column2 struct {
	Name         string  `json:"Name"`
	Label        string  `json:"Label"`
	Category     string  `json:"Category"`
	Type         string  `json:"Type"`
	CustomerType string  `json:"CustomerType"`
	FamilyName   *string `json:"FamilyName"`
	Summable     *bool   `json:"Summable"`
	Pie          *bool   `json:"Pie"`
}

type Context struct {
	Periods []struct {
		Value string `json:"Value"`
	} `json:"Periods"`
}

func (service *Service) GetData(params *GetDataParams, version GetDataVersion) (*[]DataFeed, *errortools.Error) {
	if version == GetDataVersion2 {
		return service.getData2(params)
	} else if version == GetDataVersion3 {
		return service.getData3(params)
	}

	return nil, errortools.ErrorMessagef("Invalid version '%v'", version)
}

func (service *Service) getData2(params *GetDataParams) (*[]DataFeed, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	values := url.Values{}

	columns := []string{}
	if params.FilterString2 != nil {
		values.Set("filter", *params.FilterString2)
	}
	if params.MaxResults != nil {
		values.Set("max-results", fmt.Sprintf("%v", *params.MaxResults))
	}
	columns = append(columns, params.Metrics.String()...)
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
	columns = append(columns, params.Properties.String()...)
	if params.Sort != nil {
		values.Set("sort", fmt.Sprintf("{%s}", strings.Join(params.Sort.sort, ",")))
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

	data := Data2{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url2(fmt.Sprintf("getData?%s", values.Encode())),
		ResponseModel: &data,
	}
	//fmt.Println(requestConfig.Url)
	_, _, e := service.httpRequest(&requestConfig)

	return &data.DataFeed, e
}

func (service *Service) getData3(params *GetDataParams) (*[]DataFeed, *errortools.Error) {
	if params == nil {
		return nil, nil
	}

	data := Data{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url3("getData"),
		BodyModel:     params.Params(),
		ResponseModel: &data,
	}
	//fmt.Println(requestConfig.Url)
	//b, _ := json.Marshal(requestConfig.BodyModel)
	//fmt.Println(string(b))
	_, _, e := service.httpRequest(&requestConfig)

	return &[]DataFeed{data.DataFeed}, e
}

// GetRowCount
//
type GetRowCountParams struct {
	Properties   Properties
	Metrics      Metrics
	FilterString *string
	Space        Space
	Period       Period
	Options      *Options
}

func (p *GetRowCountParams) Params() *params {
	if p == nil {
		return nil
	}

	pa := params{}

	pa.Columns = append(pa.Columns, p.Properties.String()...)
	pa.Columns = append(pa.Columns, p.Metrics.String()...)
	pa.FilterString2 = p.FilterString
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
		Method:        http.MethodPost,
		Url:           service.url3("getRowCount"),
		BodyModel:     params.Params(),
		ResponseModel: &rowCounts,
	}
	_, _, e := service.httpRequest(&requestConfig)

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
		Method:        http.MethodPost,
		Url:           service.url3("getTotal"),
		BodyModel:     params.Params(),
		ResponseModel: &rowCounts,
	}
	_, _, e := service.httpRequest(&requestConfig)

	return &rowCounts, e
}
