package ui

import (
	"github.com/carbonyde/tungsten"
)

var HotReloadScript = tungsten.InlineScript(`
let socket;
let reconnectionTimerId;

const requestUrl = window.location.origin.replace("http", "ws") + "/hot-reload";

function connect() {
  if (socket) {
    socket.close();
  }

  socket = new WebSocket(requestUrl);

  socket.addEventListener("open", () => {
    socket.send("ping");
  });

  socket.addEventListener("message", (event) => {
    if (event.data !== "refresh") {
      return;
    }

    console.info("refreshing...");

    window.location.reload();
  });

  socket.addEventListener("close", () => {
    console.info("connection closed");

    clearTimeout(reconnectionTimerId);

    reconnectionTimerId = setTimeout(connect, 300);
  });
}

connect();
`)
