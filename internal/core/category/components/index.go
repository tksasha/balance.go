package components

import (
	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func Index(categories category.Categories) Node {
	return Table(
		Class("table"),
		THead(
			Tr(
				Th(Text("Name")),
			),
		),
		TBody(
			Map(categories, row),
		),
	)
}
