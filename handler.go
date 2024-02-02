package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handler(c echo.Context) error {
	token := c.FormValue("token")
	if token == "" {
		return c.String(http.StatusForbidden, "token is null")
	}

	rssData, err := fetchRss(token)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}

	data, err := fixDate(rssData)
	if err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationXMLCharsetUTF8, data)
}
