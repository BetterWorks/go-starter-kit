package controllers

import (
	"github.com/BetterWorks/gosk-api/internal/types"
	"github.com/hetiansu5/urlquery"
)

func parseQuery(qs []byte) *types.QueryData {
	data := &types.QueryData{}
	urlquery.Unmarshal(qs, data)
	data.Paging = pageSettings(data.Paging)
	data.Sorting = sortSettings(data.Sorting)
	return data
}

func pageSettings(p types.QueryPaging) types.QueryPaging {
	var (
		defaultLimit  = 20 // TODO: move to config
		defaultOffset = 0  // TODO: move to config
	)

	page := types.QueryPaging{
		Limit:  &defaultLimit,
		Offset: &defaultOffset,
	}

	if p.Limit != nil {
		page.Limit = p.Limit
	}
	if p.Offset != nil {
		page.Offset = p.Offset
	}

	return page
}

func sortSettings(s types.QuerySorting) types.QuerySorting {
	var (
		defaultOrder = "desc"       // TODO: move to config
		defaultProp  = "created_on" // TODO: move to config
	)

	sort := types.QuerySorting{
		Order: &defaultOrder,
		Prop:  &defaultProp,
	}

	if s.Order != nil {
		sort.Order = s.Order
	}
	if s.Prop != nil {
		sort.Prop = s.Prop
	}

	return sort
}
