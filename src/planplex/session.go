/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

import(
    "bytes"
    "net/http"
    "encoding/json"
    "errors"
    "io"
)

type Session struct {
    ServerUrl string
    Id string
    Username string

    client *http.Client
}

func (session *Session) Login(username, password string) error {
    data, _ := json.Marshal(map[string]string {
        "username": username,
        "password": password,
    })

    if response, error := http.Post(session.ServerUrl + "/api/session/login", "text/json", bytes.NewBuffer(data)); error != nil {
        return error
    } else {
        for _, cookie := range response.Cookies() {
            if cookie.Name == "sessionid" {
                session.Id = cookie.Value
                session.Username = username
                session.client = &http.Client{}
                return nil
            }
        }

        return errors.New(response.Status)
    }
}

func (session *Session) makeRequest(verb string, url string, body io.Reader) (request *http.Request, error error) {
    if request, error := http.NewRequest(verb, session.ServerUrl + url, body); error != nil {
        return nil, error
    } else {
        request.AddCookie(&http.Cookie{Name: "sessionid", Value: session.Id})
        return request, nil 
    }
}

func (session *Session) Projects() (projects []*Project, error error) {
    if request, error := session.makeRequest("GET", "/api/users/" + session.Username + "/projects", nil); error != nil {
        return nil, error
    } else {
        if response, error := session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&projects); error != nil {
                return nil, error
            } else {
                for _, project := range projects {
                    project.session = session
                }

                return projects, nil
            }
        }
    }
}

func (session *Session) Logout() error {
    if request, error := session.makeRequest("POST", "/api/session/logout", nil); error != nil {
        return error
    } else {
        if response, error := session.client.Do(request); error != nil {
            return error
        } else if response.StatusCode != http.StatusOK {
            return errors.New(response.Status)
        } else {
            return nil
        }
    }
}
