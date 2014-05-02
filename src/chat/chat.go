package main

import(
    "planplex"
    "fmt"
    "os"
    "code.google.com/p/gopass"
)

func main() {
    if len(os.Args) > 1 {
        var url, username, password string

        url = os.Args[1]
        fmt.Printf("Username: ")
        fmt.Scanf("%s", &username)
        password, _ = gopass.GetPass("Password: ")
        session := planplex.Session{ ServerUrl:  url }

        if error := session.Login(username, password); error != nil {
            fmt.Println("Login failed: " + error.Error())
        } else {
            if projects, error := session.Projects(); error != nil {
                fmt.Println("Project enumeration failed: " + error.Error())
            } else {
                var projectIndex int

                fmt.Printf("Available projects:\n")
                for index, value := range projects {
                    fmt.Printf("\t[%d] %s\n", index, value.Identifier)
                }
                fmt.Printf("\nChoose project: ")
                fmt.Scanf("%d", &projectIndex)

                if error := projects[projectIndex].Activate(); error != nil {
                    fmt.Println("Project activation failed: " + error.Error())
                } else {
                    var taskIndex uint32

                    if tasks, error := projects[projectIndex].Tasks(); error == nil {
                        for _, task := range tasks {
                            fmt.Printf("[%d] %s\n", task.Id, task.Name)
                        }

                        fmt.Printf("\nChoose task: ")
                        fmt.Scanf("%d", &taskIndex)

                        if task, error := projects[projectIndex].Task(taskIndex); error == nil {
                            messages, _ := task.Messages()

                            for _, message := range messages {
                                fmt.Println(message)
                            }
                        }
                    }
                }
            }

            if error := session.Logout(); error != nil {
                fmt.Println("Logout failed: " + error.Error())
            } else {
                fmt.Println("Logout successful.")
            }
        }
    }
}
