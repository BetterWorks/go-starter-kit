package controllers

import (
	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/hetiansu5/urlquery"
)

type QueryConfig struct {
	Defaults *QueryDefaults
}

type QueryDefaults struct {
	Paging  *query.QueryPaging  `validate:"required"`
	Sorting *query.QuerySorting `validate:"required"`
}

type queryHandler struct {
	defaults *QueryDefaults
}

func NewQueryHandler(c *QueryConfig) *queryHandler {
	return &queryHandler{
		defaults: c.Defaults,
	}
}

func (q *queryHandler) parseQuery(qs []byte) *query.QueryData {
	data := &query.QueryData{}

	// TODO: validate query
	urlquery.Unmarshal(qs, data)
	// if err := validation.Validate.Struct(data); err != nil {
	// 	return nil, err
	// }
	data.Paging = q.pageSettings(data.Paging)
	data.Sorting = q.sortSettings(data.Sorting)

	return data
}

func (q *queryHandler) pageSettings(p query.QueryPaging) query.QueryPaging {
	page := query.QueryPaging{
		Limit:  q.defaults.Paging.Limit,
		Offset: q.defaults.Paging.Offset,
	}

	if p.Limit != nil {
		page.Limit = p.Limit
	}
	if p.Offset != nil {
		page.Offset = p.Offset
	}

	return page
}

func (q *queryHandler) sortSettings(s query.QuerySorting) query.QuerySorting {
	sort := query.QuerySorting{
		Order: q.defaults.Sorting.Order,
		Attr:  q.defaults.Sorting.Attr,
	}

	if s.Attr != nil {
		sort.Attr = s.Attr
	}
	if s.Order != nil {
		sort.Order = s.Order
	}

	return sort
}
