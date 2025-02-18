package components

import (
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func Index(helpers *helpers.Helpers, items item.Items) Node {
	return Table(
		Class("table"),
		THead(
			Tr(
				Th(Class("text-center"), Text("Date")),
				Th(Class("text-center"), Text("Sum")),
				Th(Class("text-center"), Text("Category")),
				Th(Class("text-center"), Text("Description")),
			),
		),
		TBody(
			Map(items, func(item *item.Item) Node {
				return row(helpers, item)
			}),
		),
	)
}
