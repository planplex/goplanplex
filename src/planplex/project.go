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
    "bytes"
)

type Project struct {
    Identifier string
    Name string
    Owned bool
    Description string
    Assignments []int

    session *Session
}

func (project *Project) String() string {
    var fields = []string {
        "Type: Project",
        "Identifier: " + project.Identifier,
        "Name: " + project.Name,
        "Owned: " + strconv.FormatBool(project.Owned),
        "Description: " + project.Description,
        "Assignments: " + fmt.Sprintf("%v", project.Assignments),
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Activate() error {
    data, _ := json.Marshal(map[string]string {
        "identifier": project.session.Username + "/" + project.Name,
    })

    if request, error := project.session.makeRequest("POST", "/api/session/project", bytes.NewReader(data)); error != nil {
        return error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return error
        } else if response.StatusCode != http.StatusOK {
            return errors.New(response.Status)
        } else {
            return nil
        }
    }
}

type ProjectSummaryDates struct {
    Planned Time 
    Estimated Time
}

func (dates *ProjectSummaryDates) String() string {
    var fields = []string {
        time.Time(dates.Planned).String() + " (Planned)",
        time.Time(dates.Estimated).String() + " (Estimated)",
    }
    
    return strings.Join(fields, " ")
}

type ProjectSummaryEffort struct {
    Planned Duration
    Estimated Duration
    Current Duration
}

func (effort *ProjectSummaryEffort) String() string {
    var fields = []string {
        time.Duration(effort.Planned).String() + " (Planned)",
        time.Duration(effort.Estimated).String() + " (Estimated)",
        time.Duration(effort.Current).String() + " (Current)",
    }
    
    return strings.Join(fields, " ")
}

type ProjectSummary struct {
    Start ProjectSummaryDates
    End ProjectSummaryDates
    Effort ProjectSummaryEffort
    Duration ProjectSummaryEffort
}

func (summary *ProjectSummary) String() string {
    var fields = []string {
        "Start: " + summary.Start.String(),
        "End: " + summary.End.String(),
        "Effort: " + summary.Effort.String(),
        "Duration: " + summary.Duration.String(),
    }
    
    return strings.Join(fields, "\n")
}

func (project *Project) Summary() (summary *ProjectSummary, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/summary", nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&summary); error != nil {
                return nil, error
            } else {
                return summary, nil
            }
        }
    }
}

