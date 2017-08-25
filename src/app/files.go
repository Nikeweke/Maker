package app

// import "fmt"


/* ----------------------------------------------
|  Файлы
|
|------------------------------------------------
*/
// ----------------------> main.go
var maingo = `
package main
import "fmt"

func main(){
 fmt.Println("Hello its me, Mario")
}
`



// ----------------------> run.bat
var runbat = `
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

go run main.go
REM go build main.go
pause
`



// ----------------------> pusher.bat
var pusher = `
git init
git remote add origin https://github.com/.....
git add -A
git commit -m "new coomit"
git push -u origin master
REM git push
`




// ----------------------> .gitignore
var gitignore = `
src/*
!src/app
node_modules/
vendor/
pusher.bat
build/
`



// ----------------------> readme.md
var readmemd = `
## Its my app
* Hello
* goodbye
`
