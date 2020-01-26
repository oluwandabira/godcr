package page

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/raedahgroup/godcr-gio/event"
	"github.com/raedahgroup/godcr-gio/ui/units"
)

// LoadingID is the id of the loading page.
const LoadingID = "loading"

// Loading represents the loading page of the app.
type Loading struct {
	label material.Label
}

// Init initializies the page with a label.
func (pg *Loading) Init(theme *material.Theme) {
	pg.label = theme.Label(units.Label, "LOADING")
}

// Draw renders the page widgets.
// It does not react to nor does it generate any event.
func (pg *Loading) Draw(gtx *layout.Context, _ event.Event) event.Event {
	pg.label.Layout(gtx)
	return nil
}