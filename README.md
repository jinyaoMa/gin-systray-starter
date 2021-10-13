# gin-systray-starter

starter code for a Go system tray application with Gin

- systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

## Additions

- air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- gorm, [https://gorm.io/](https://gorm.io/)

## System

- Windows 10
- Go 1.17
- Npm v8

## Setup

``` bash
# install air cli
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# install go dependencies
go mod tidy
```

## Npm Scripts

Run only the http server with air hot reload,

```
npm run serve
```

Build the project (with flags `-tags=jsoniter -ldflags "-H=windowsgui"`)

```
npm run build
```


Build and run the project

```
npm run build:run
```

## Structure

- database `init database connection`
  - models `map database structure && provide CRUD for each table`
- server `init http server`
  - controllers `actions for each API (request -> services -> response)`
  - middlewares `hooks/filters/interceptors`
  - routers `groups of APIs`
  - services `services (involve multiple models) to be used in actions`
- tray `init system tray`
  - icons `icons for the tray and menu items`
  - locales `language files for the tray`
- .air.toml `air cli config`
- main.go `app entry`
