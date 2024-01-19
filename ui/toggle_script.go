package ui

import (
	"github.com/carbonyde/tungsten"
)

var ToggleScript = tungsten.InlineScript(`
const isDark =
  window.localStorage.theme === "dark" ||
  (!("theme" in window.localStorage) &&
    window.matchMedia("(prefers-color-scheme: dark)").matches);

const toggles = document.querySelectorAll("#theme-toggle");
const iconDark = document.querySelector("#theme-toggle-dark");
const iconLight = document.querySelector("#theme-toggle-light");

if (isDark) {
  iconDark?.classList.toggle("hidden");
} else {
  iconLight?.classList.toggle("hidden");
}

for (const toggle of toggles) {
  toggle.addEventListener("click", function () {
    const isDark = document.documentElement.classList.contains("dark");

    if (isDark) {
      window.localStorage.setItem("theme", "");
      document.documentElement.classList.remove("dark");
      iconDark?.classList.toggle("hidden");
      iconLight?.classList.toggle("hidden");
    } else {
      window.localStorage.setItem("theme", "dark");
      document.documentElement.classList.add("dark");
      iconDark?.classList.toggle("hidden");
      iconLight?.classList.toggle("hidden");
    }
  });
}
`)
