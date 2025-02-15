package components

import (
	"github.com/tksasha/balance/internal/core/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Index(cashes cash.Cashes) Node {
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
			Map(cashes, row),
		),
	)
}
