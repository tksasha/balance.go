package components

import (
	"net/url"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/core/common/component"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

const yearFrom = 2015

type YearsComponent struct {
	*component.Component

	years []int
}

func NewYearsComponent(component *component.Component) *YearsComponent {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return &YearsComponent{
		Component: component,
		years:     years,
	}
}

func (c *YearsComponent) Years(values url.Values) Node {
	return Div(
		ID("years"),
		Class("row mt-2 text-center"),
		htmx.SwapOOB("true"),
		Map(c.years, func(year int) Node {
			return c.year(year, values)
		}),
	)
}

func (c *YearsComponent) year(year int, values url.Values) Node {
	val, err := strconv.Atoi(values.Get("year"))
	current := err == nil && year == val

	classes := components.Classes{
		"col":            true,
		"link-secondary": !current,
		"link-success":   current,
	}

	return A(
		classes,
		Href(c.ListItems(year, 0, values)),
		Text(strconv.Itoa(year)),
		htmx.Get(c.ListItems(year, 0, values)),
		htmx.Target("#items"),
	)
}
