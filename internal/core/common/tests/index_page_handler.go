package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/balance/internal/core/indexpage"
	"github.com/tksasha/balance/internal/core/indexpage/components"
	"github.com/tksasha/balance/internal/core/indexpage/handler"
)

func NewIndexPageHandler(
	t *testing.T,
	indexPageService indexpage.Service,
	categoryService category.Service,
) *handler.Handler {
	t.Helper()

	component := component.New()

	monthsComonents := components.NewMonthsComponent(component)

	indexPageComponent := components.NewIndexPageComponent(component, monthsComonents)

	return handler.New(common.NewBaseHandler(), indexPageService, categoryService, indexPageComponent)
}
