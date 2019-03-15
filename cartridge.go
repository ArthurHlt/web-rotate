package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

func cartridge(page *Page) func(context.Context, cdp.Executor) error {
	return func(ctxt context.Context, h cdp.Executor) error {
		// set up parameters
		p := runtime.Evaluate(fmt.Sprintf(`
window.onload = function () {
	document.body.innerHTML += '<div id="chrmodp-cartridge" style="%s">' + 
    	 '<span style="%s">%s</span>'+
     	'</div>';
}
`, page.StyleCartridge.Box.ToCss(), page.StyleCartridge.Text.ToCss(), page.Cartridge))
		// evaluate
		_, _, err := p.Do(ctxt, h)
		if err != nil {
			return err
		}
		// unmarshal
		return nil
	}
}
