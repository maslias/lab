package main

import "github.com/maslias/tasks/cmd"

func main() {
    // config.Load()
    // fmt.Printf("config DB_PATH: %v", config.GET().DB_PATH)
    // fmt.Printf("DB_PATH: %v", os.Getenv("DB_PATH"))
    cmd.Execute()
}
