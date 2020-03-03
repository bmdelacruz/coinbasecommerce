package coinbasecommerce

import (
	"net/url"
	"strconv"
)

// PaginationOrder represents the order of the items in the charges list.
type PaginationOrder string

// PaginationOrder constants.
const (
	PaginationOrderDesc PaginationOrder = "desc"
	PaginationOrderAsc  PaginationOrder = "asc"
)

// PaginationOption contains options for paginated responses.
type PaginationOption struct {
	order         PaginationOrder
	limit         int
	startingAfter string
	endingBefore  string
}

// MakeQueryString makes a query string based on the PaginationOption.
func (options PaginationOption) MakeQueryString() string {
	values := make(url.Values)
	if options.order != PaginationOrderDesc {
		values.Set("order", string(options.order))
	}
	if options.limit != 25 {
		values.Set("limit", string(options.limit))
	}
	if options.startingAfter != "" {
		values.Set("starting_after", options.startingAfter)
	}
	if options.endingBefore != "" {
		values.Set("ending_before", options.endingBefore)
	}
	return values.Encode()
}

// PaginationOptionFunc represents a function that accepts and modifies a
// pagination options object.
type PaginationOptionFunc func(*PaginationOption)

// PaginationOptionOrder creates a PaginationOptionFunc that will set the
// PaginationOrder value of a PaginationOption object.
func PaginationOptionOrder(order PaginationOrder) PaginationOptionFunc {
	if order != PaginationOrderDesc && order != PaginationOrderAsc {
		panic(`invalid pagination order. valid values: "desc" or "asc"`)
	}
	return func(options *PaginationOption) {
		options.order = order
	}
}

// PaginationOptionLimit creates a PaginationOptionFunc that will set the
// limit of a PaginationOption object.
func PaginationOptionLimit(limit int) PaginationOptionFunc {
	if limit > 100 || limit < 0 {
		panic(`invalid pagination limit. valid values: 0 <= limit <= 100`)
	}
	return func(options *PaginationOption) {
		options.limit = limit
	}
}

// PaginationOptionStartingAfter creates a PaginationOptionFunc that will set the
// starting after of a PaginationOption object.
func PaginationOptionStartingAfter(startingAfter string) PaginationOptionFunc {
	if startingAfter == "" {
		panic(`invalid pagination starting after. valid value cannot be empty`)
	}
	return func(options *PaginationOption) {
		options.startingAfter = startingAfter
	}
}

// PaginationOptionEndingBefore creates a PaginationOptionFunc that will set the
// ending before of a PaginationOption object.
func PaginationOptionEndingBefore(endingBefore string) PaginationOptionFunc {
	if endingBefore == "" {
		panic(`invalid pagination ending before. valid value cannot be empty`)
	}
	return func(options *PaginationOption) {
		options.endingBefore = endingBefore
	}
}

// NewPaginationOption creates a new pagination option object.
func NewPaginationOption(optionFuncs ...PaginationOptionFunc) PaginationOption {
	options := PaginationOption{
		order:         PaginationOrderDesc,
		limit:         25,
		startingAfter: "",
		endingBefore:  "",
	}
	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}
	return options
}

// NewPaginationOptionFromQueryString parses the query string, looks for valid values
// and then returns a new pagination option with values from the query string.
func NewPaginationOptionFromQueryString(queryString string) (PaginationOption, error) {
	options := PaginationOption{
		order:         PaginationOrderDesc,
		limit:         25,
		startingAfter: "",
		endingBefore:  "",
	}

	values, err := url.ParseQuery(queryString)
	if err != nil {
		return options, err
	}

	if order := values.Get("order"); order == string(PaginationOrderAsc) || order == string(PaginationOrderDesc) {
		options.order = PaginationOrder(order)
	}
	if limitStr := values.Get("limit"); limitStr != "" {
		if limit32, err := strconv.ParseInt(limitStr, 10, 32); err == nil {
			if limit := int(limit32); limit >= 0 && limit < 101 {
				options.limit = limit
			}
		}
	}
	if startingAfter := values.Get("starting_after"); startingAfter != "" {
		options.startingAfter = startingAfter
	}
	if endingBefore := values.Get("ending_before"); endingBefore != "" {
		options.endingBefore = endingBefore
	}

	return options, nil
}

// Pagination represents the pagination of a list that was retrieved from the
// Coinbase Commerce API.
type Pagination struct {
	Order         string   `json:"order"`
	StartingAfter *string  `json:"starting_after"`
	EndingBefore  *string  `json:"ending_before"`
	Total         int      `json:"total"`
	Yielded       int      `json:"yielded"`
	Limit         int      `json:"limit"`
	PreviousURI   *string  `json:"previous_uri"`
	NextURI       *string  `json:"next_uri"`
	CursorRange   []string `json:"cursor_range"`
}
