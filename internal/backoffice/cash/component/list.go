package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Table(
		THead(
			Tr(
				Text("Name"),
			),
			Tr(
				Text("Sum"),
			),
		),
		TBody(
			Map(cashes, func(cash *cash.Cash) Node {
				return Tr(
					Td(
						Text(cash.Name),
					),
					Td(
						Text(c.Money(cash.Sum)),
					),
				)
			}),
		),
	)
}
