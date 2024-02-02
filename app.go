package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func handler(c echo.Context) error {
	token := c.FormValue("token")
	if token == "" {
		return c.String(http.StatusForbidden, "token is null")
	}

	var (
		rssData []byte
		err     error
	)

	for i := 0; i < 3; i++ {
		rssData, err = fetchRss(token)
		if err == nil {
			break
		}

		time.Sleep(time.Second * 3)
	}
	if err != nil {
		return c.String(http.StatusServiceUnavailable,
			fmt.Sprintf("Failed to fetch RSS content after retrying 3 times. The last err: %v", err))
	}

	data, err := fixDate(rssData)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationXMLCharsetUTF8, data)
}
