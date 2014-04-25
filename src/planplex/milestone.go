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
    "time"
    "strings"
    "fmt"
)

type Milestone struct {
    Object
    Deadline Time 
    Requirements []int 
    Met bool
}

func (milestone *Milestone) String() string {
    var fields = []string {
        "Type: Milestone",
        milestone.Object.String(),
        "Deadline: " + time.Time(milestone.Deadline).String(),
        "Requiresments: " + fmt.Sprintf("%v", milestone.Requirements),
        "Met: " + strconv.FormatBool(milestone.Met),
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Milestones() (milestones []*Milestone, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/milestones", nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&milestones); error != nil {
                return nil, error
            } else {
                return milestones, nil
            }
        }
    }
}

func (project *Project) Milestone(id uint32) (milestone *Milestone, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/milestones/" + strconv.FormatUint(uint64(id), 10), nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&milestone); error != nil {
                return nil, error
            } else {
                return milestone, nil
            }
        }
    }
}

