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
                var index int

                fmt.Printf("Available projects:\n")
                for index, value := range projects {
                    fmt.Printf("\t[%d] %s\n", index, value.Identifier)
                }
                fmt.Printf("\nChoose project: ")
                fmt.Scanf("%d", &index)

                if error := projects[index].Activate(); error != nil {
                    fmt.Println("Project activation failed: " + error.Error())
                } else {
                    fmt.Println(projects[index])
                    
                    if summary, error := projects[index].Summary(); error == nil {
                        fmt.Println(summary, "\n")
                    }

                    if resources, error := projects[index].Resources(); error == nil {
                        for _, resource := range resources {
                            fmt.Println(resource, "\n")
                        }
                    }
                    
                    if calendars, error := projects[index].Calendars(); error == nil {
                        for _, calendar := range calendars {
                            fmt.Println(calendar, "\n")
                        }
                    }
                    
                    if milestones, error := projects[index].Milestones(); error == nil {
                        for _, milestone := range milestones {
                            fmt.Println(milestone, "\n")
                        }
                    }
                    
                    if page, error := projects[index].Page("start"); error == nil {
                        fmt.Println(page, "\n")
                    }
                    
                    if tasks, error := projects[index].Tasks(); error == nil {
                        for _, task := range tasks {
                            task, _ := projects[index].Task(task.Id)
                            fmt.Println(task, "\n")
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
