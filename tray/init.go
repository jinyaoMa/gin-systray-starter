package tray

import (
	_ "embed"
	"this/server"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

var (
	serverMenu      *systray.MenuItem
	serverStart     *systray.MenuItem
	serverStartSwag *systray.MenuItem
	serverStop      *systray.MenuItem
	languageMenu    *systray.MenuItem
	languageEn      *systray.MenuItem
	languageZh      *systray.MenuItem
	quit            *systray.MenuItem

	serverStatusChannel = make(chan bool, 2)
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	makeTray()
	resetLanguage()

	go watchClicks()
	go watchStatus()

	serverStatusChannel <- false
}

func onExit() {
	close(serverStatusChannel)
	server.Stop()
}

func watchClicks() {
	for {
		isServerRunning := server.IsRunning
		select {
		case <-serverStart.ClickedCh:
			server.Run(false)
			isServerRunning = true
		case <-serverStartSwag.ClickedCh:
			server.Run(true)
			isServerRunning = true
		case <-serverStop.ClickedCh:
			server.Stop()
			isServerRunning = false
		case <-languageEn.ClickedCh:
			if setLocale("en") {
				resetLanguage()
			}
		case <-languageZh.ClickedCh:
			if setLocale("zh") {
				resetLanguage()
			}
		case <-quit.ClickedCh:
			systray.Quit()
			return
		}
		serverStatusChannel <- isServerRunning
	}
}

func watchStatus() {
	for {
		isServerRunning, ok := <-serverStatusChannel
		if !ok {
			return
		}

		var serverStatusMessage, languageStatusMessage string
		if isServerRunning {
			serverStart.Disable()
			serverStartSwag.Disable()
			serverStop.Enable()
			if server.HasSwagger {
				serverStatusMessage = getMessage("server-start-swag")
			} else {
				serverStatusMessage = getMessage("server-start")
			}
		} else {
			serverStart.Enable()
			serverStartSwag.Enable()
			serverStop.Disable()
			serverStatusMessage = getMessage("server-stop")
		}
		switch locale {
		case "en":
			languageEn.Check()
			languageZh.Uncheck()
			languageStatusMessage = getMessage("language-en")
		case "zh":
			languageEn.Uncheck()
			languageZh.Check()
			languageStatusMessage = getMessage("language-zh")
		}

		systray.SetTooltip(getMessageWithParams("tooltip", serverStatusMessage, languageStatusMessage))
	}
}

func resetLanguage() {
	systray.SetTitle(getMessage("title"))

	serverMenu.SetTitle(getMessage("server"))
	serverMenu.SetTooltip(getMessage("server"))
	serverStart.SetTitle(getMessage("server-start"))
	serverStart.SetTooltip(getMessage("server-start"))
	serverStartSwag.SetTitle(getMessage("server-start-swag"))
	serverStartSwag.SetTooltip(getMessage("server-start-swag"))
	serverStop.SetTitle(getMessage("server-stop"))
	serverStop.SetTooltip(getMessage("server-stop"))

	languageMenu.SetTitle(getMessage("language"))
	languageMenu.SetTooltip(getMessage("language"))
	languageEn.SetTitle(getMessage("language-en"))
	languageEn.SetTooltip(getMessage("language-en"))
	languageZh.SetTitle(getMessage("language-zh"))
	languageZh.SetTooltip(getMessage("language-zh"))

	quit.SetTitle(getMessage("quit"))
	quit.SetTooltip(getMessage("quit"))
}

func makeTray() {
	systray.SetTemplateIcon(icon, icon)

	serverMenu = systray.AddMenuItem("", "")
	serverStart = serverMenu.AddSubMenuItemCheckbox("", "", false)
	serverStartSwag = serverMenu.AddSubMenuItemCheckbox("", "", false)
	serverStop = serverMenu.AddSubMenuItemCheckbox("", "", false)

	systray.AddSeparator()

	languageMenu = systray.AddMenuItem("", "")
	languageEn = languageMenu.AddSubMenuItemCheckbox("", "", false)
	languageZh = languageMenu.AddSubMenuItemCheckbox("", "", false)

	systray.AddSeparator()

	quit = systray.AddMenuItem("", "")
}
