package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) form(category *category.Category, _ validation.Errors) Node {
	return Form(
		If(category.ID == 0, htmx.Post(path.CreateBackofficeCategory(nil))),
		If(category.ID != 0, htmx.Patch(path.UpdateBackofficeCategory(nil, category.ID))),
		Div(Class("mb-3"),
			Label(Class("form-label"), Text("Валюта")),
			Select(Class("form-select"), Name("currency"),
				c.CurrencyOptions(currency.GetCode(category.Currency))),
		),
		c.Input("Назва", "name", category.Name, nil, nil, AutoFocus()),
		c.CheckBox("Надходження", "is-category-income", category.Income),
		c.CheckBox("Відображається", "is-category-visible", category.Visible),
		c.Input("Група", "supercategory", strconv.Itoa(category.Supercategory), nil, nil),
		c.Submit(category.ID),
	)
}
