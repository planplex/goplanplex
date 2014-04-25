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

type Calendar struct {
    Object
    In Time `json:"_in"`
    Out Time
    Days []bool
}

func (calendar *Calendar) String() string {
    var fields = []string {
        "Type: Calendar",
        calendar.Object.String(),
        "In: " + time.Time(calendar.In).String(),
        "Out: " + time.Time(calendar.Out).String(),
        "Days: " + fmt.Sprintf("%v", calendar.Days),
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Calendars() (calendars []*Calendar, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/calendars", nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&calendars); error != nil {
                return nil, error
            } else {
                return calendars, nil
            }
        }
    }
}

func (project *Project) Calendar(id uint32) (calendar *Calendar, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/calendars/" + strconv.FormatUint(uint64(id), 10), nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&calendar); error != nil {
                return nil, error
            } else {
                return calendar, nil
            }
        }
    }
}

