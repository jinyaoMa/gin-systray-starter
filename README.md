# gin-systray-starter

Template/Starter code for Go application with Gin, System Tray, Gorm, Air, Swagger, JWT

- [x] systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- [x] gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [x] air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- [x] gorm, [https://gorm.io/](https://gorm.io/)
- [x] swagger, [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- [x] jwt, [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)

## System

- Windows 10
- Go 1.17
- Npm v8

## Setup

``` bash
# install air cli
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# install swag cli
go install github.com/swaggo/swag/cmd/swag@latest
# install go dependencies
go mod tidy
```

## Npm Scripts

Run only the http server with air hot reload,

```
npm run serve
```

Re-generate swagger configuration (do this when APIs changed)

```
npm run swag
```

Build and run the project (without flags)

```
npm run build:run
```

Build the project (with flags, `-tags=jsoniter -ldflags "-H=windowsgui"`)

```
npm run build
```

## Structure

- /air `air configuration and temporary output files`
  - .air.toml `air cli config`
- /database `init database connection`
  - /models `map database structure && provide CRUD for each table`
- /gate `authentication tools`
- /server `init http server`
  - /controllers `actions for each API (request -> services -> response)`
  - /middlewares `hooks/filters/interceptors`
    - auth.go `authentication gateway with jwt`
  - /routers `groups of APIs`
  - /services `services (involve multiple models) to be used in actions`
  - api.go `swag General API annotations with a set of API routers/groups`
- /swagger `auto-generated swagger configuration`
- /tray `init system tray`
  - /icons `icons for the tray and menu items`
  - /locales `language files for the tray`
- main.go `app entry`
