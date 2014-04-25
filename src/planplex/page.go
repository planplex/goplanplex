/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

import(
    "net/http"
    "encoding/json"
    "errors"
    "strings"
)

type Page struct {
    Content string
}

func (page *Page) String() string {
    var fields = []string {
        "Type: Page",
        "Content: " + page.Content,
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Page(id string) (page *Page, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/wiki/" + id, nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&page); error != nil {
                return nil, error
            } else {
                return page, nil
            }
        }
    }
}
