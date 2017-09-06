package app

import (
   "os"
   "fmt"
)

        /* ----------------------------------------------
        |  Старт
        |
        |------------------------------------------------
        */
        func Main(){

           var choice int
           var files []string

           var files_name = []string{".gitignore", "pusher.bat", "readme.md", "main.go", "run.bat"}

           Greeting()
           fmt.Scanln(&choice)
           fmt.Println()


           switch choice {
             case 1:
                    files = append(files, files_name[0], files_name[1])
                    CreateFilesFromArr(files)
             case 2:
                    files = append(files, files_name[0], files_name[1], files_name[2])
                    CreateFilesFromArr(files)
             case 3:
                     files = append(files, files_name[3], files_name[4])
                     CreateFilesFromArr(files)
             case 4:
                    files = append(files, files_name[0], files_name[1], files_name[2], files_name[3], files_name[4])
                    CreateFilesFromArr(files)
             case 8:
                 os.Exit(3)
             default:
                 fmt.Println("OOOOHH IT IS WARNING: You typed something wrong, choose correct answer")
                 fmt.Println()
                 fmt.Println()
                 Main()
           }
        }


        /* ----------------------------------------------
        |  Создание файлов по данному массиву
        |
        |------------------------------------------------
        */
        func CreateFilesFromArr(files []string){

             // создаем папку и переходим в неё
             os.Mkdir("MAKER_files", 1)
             os.Chdir("MAKER_files")

             for _, name := range files {
                //  fmt.Println( v)
                   if name == ".gitignore"{
                      MakeFile(name,   gitignore)

                   } else if name == "readme.md" {
                      MakeFile(name,   readmemd)

                   } else if name == "pusher.bat" {
                      MakeFile("pusher.bat" ,   pusher)

                   } else if name == "main.go" {
                      MakeFile("main.go",   maingo)

                   } else{
                      MakeFile("run.bat",   runbat)
                   }
             }
      }




      /* ----------------------------------------------
      |  Приветсвие
      |
      |------------------------------------------------
      */
      func Greeting(){
        fmt.Println("======================== Hello. I am MAKER ======================= ")
        fmt.Println("type '1' ===> .gitignore + pusher.bat ")
        fmt.Println("type '2' ===> .gitignore + pusher.bat + readme.md")
        fmt.Println("type '3' ===>  main.go + run.bat ")
        fmt.Println("type '4' ===> .gitignore + pusher.bat + readme.md + main.go + run.bat ")
        fmt.Println("type '8' to exit ")
        fmt.Println()
        fmt.Print("CHOICE NOW: ")
      }
