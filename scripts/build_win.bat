cd ..
@ECHO OFF

REM по стандарту должна уже стоять
REM SET GOROOT=C:\GO
SET GOPATH=%CD%

REM путь к бинарникам проекта
SET PATH=%GOPATH%\BIN;%PATH%;

REM SET GOOS=linux
REM SET GOARCH=amd64
REM SET CGO_ENABLED=0

REM go get github.com/gorilla/websocket
REM go get github.com/robfig/cron
REM go get github.com/drgomesp/gorm
REM go get github.com/go-sql-driver/mysql
REM go get github.com/pilu/fresh


REM fresh - hot reloader
go build -o build/maker_win.exe index.go
pause
