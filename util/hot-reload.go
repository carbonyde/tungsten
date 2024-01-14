package util

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}
var reloaded = false

func HotReload(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		if reloaded {
			continue
		}

		_, message, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		if string(message[:]) == "ping" {
			err := ws.WriteMessage(websocket.TextMessage, []byte("refresh"))
			reloaded = true

			if err != nil {
				c.Logger().Error(err)
			}
		}

	}
}
