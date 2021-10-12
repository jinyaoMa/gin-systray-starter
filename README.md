# gin-systray-starter

starter code for a Go system tray application with Gin

- systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

## Additions

- air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- gorm, [https://gorm.io/](https://gorm.io/)

## Setup

``` bash
# install air cli
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# install dependencies
npm install
go mod tidy
```
