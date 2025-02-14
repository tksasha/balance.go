package components

import (
	"github.com/tksasha/balance/internal/core/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Edit(cash *cash.Cash) Node {
	return Div(
		form(cash, nil),
	)
}
