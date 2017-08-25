package app

import (
   "os"
   "fmt"
)








        /* ----------------------------------------------
        |  Создание файла
        |
        |------------------------------------------------
        */
        func MakeFile(name, text string){
            file, err := os.Create(name)
            checkErr(err)
            defer file.Close()
            file.WriteString(text)
        }




        /* ----------------------------------------------
        |  Решала ошибок
        |
        |------------------------------------------------
        */
        func checkErr(err error){
          if err != nil {
              fmt.Println(err)
              // os.Exit(1)
          }
        }
