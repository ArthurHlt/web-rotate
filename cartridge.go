package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

func cartridge(page *Page) func(context.Context, cdp.Executor) error {
	return func(ctxt context.Context, h cdp.Executor) error {
		elem := fmt.Sprintf(
			`'<div id="chromedp-cartridge" style="%s"><span style="%s">%s</span></div>'`,
			page.StyleCartridge.Box.ToCss(), page.StyleCartridge.Text.ToCss(), page.Cartridge,
		)
		js := fmt.Sprintf(`
window.onload = function () {
	document.body.innerHTML += %s;
}
document.body.innerHTML += %s;
`, elem, elem)
		p := runtime.Evaluate(js)
		// evaluate
		_, _, err := p.Do(ctxt, h)
		if err != nil {
			return err
		}
		// unmarshal
		return nil
	}
}
