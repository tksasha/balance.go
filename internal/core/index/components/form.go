package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func form( //nolint:funlen
	item *item.Item,
	categories category.Categories,
	errors validation.Errors,
) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Date"),
			),
			Input(
				If(
					true,
					Class("form-control invalid"),
				),
				If(
					true,
					Class("form-control"),
				),
				Value(
					components.Date(item.Date),
				),
				components.Errors("date", errors),
			),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Sum"),
			),
			Input(
				If(
					true,
					Class("form-control invalid"),
				),
				If(
					true,
					Class("form-control"),
				),
				Value(
					sum(item.Sum),
				),
			),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Category"),
			),
			Select(
				Class("form-select"),
				OptGroup(
					Attr("Label", "Expense"),
					Map(categories.Expense(), option),
				),
				OptGroup(
					Attr("Label", "Income"),
					Map(categories.Income(), option),
				),
			),
		),
	)
}
