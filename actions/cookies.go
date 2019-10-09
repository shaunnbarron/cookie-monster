package actions

import (
	"github.com/gobuffalo/buffalo"
)

func CookieHandler(c buffalo.Context) error {
	cookies := c.Request().Cookies()
	c.Set("cookies", cookies)

	return c.Render(200, r.HTML("cookies.html"))
}
