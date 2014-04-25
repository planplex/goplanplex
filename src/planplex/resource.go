/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

import(
    "net/http"
    "encoding/json"
    "errors"
    "strconv"
    "strings"
)

type Resource struct {
    Object
    Calendar uint32
    AssignedUser string `json:"assigned_user"`
    Role Role
}

func (resource *Resource) String() string {
    var fields = []string {
        "Type: Resource",
        resource.Object.String(),
        "Calendar: " + strconv.FormatUint(uint64(resource.Calendar), 10),
        "Assigned user: " + resource.AssignedUser,
        "Role: " + resource.Role.String(),
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Resources() (resources []*Resource, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/resources", nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&resources); error != nil {
                return nil, error
            } else {
                return resources, nil
            }
        }
    }
}

func (project *Project) Resource(id uint32) (resource *Resource, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/resources/" + strconv.FormatUint(uint64(id), 10), nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&resource); error != nil {
                return nil, error
            } else {
                return resource, nil
            }
        }
    }
}

