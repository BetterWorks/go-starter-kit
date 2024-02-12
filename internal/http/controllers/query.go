package controllers

import (
	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	"github.com/hetiansu5/urlquery"
)

type QueryConfig struct {
	Defaults *QueryDefaults `validate:"required"`
}

type QueryDefaults struct {
	Paging  *query.QueryPaging  `validate:"required"`
	Sorting *query.QuerySorting `validate:"required"`
}

type queryHandler struct {
	defaults *QueryDefaults
}

func NewQueryHandler(c *QueryConfig) (*queryHandler, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	if c.Defaults.Paging.Offset == nil {
		offset := 0
		c.Defaults.Paging.Offset = &offset
	}

	handler := &queryHandler{
		defaults: c.Defaults,
	}

	return handler, nil
}

func (q *queryHandler) parseQuery(qs []byte) *query.QueryData {
	data := &query.QueryData{}

	// TODO: validate query
	urlquery.Unmarshal(qs, data)
	// if err := app.Validator.Validate.Struct(data); err != nil {
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
