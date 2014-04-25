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

type TaskDependency struct {
    Id uint32
    Type string
}

func (dependency *TaskDependency) String() string {
    return fmt.Sprintf("[%d] %s", dependency.Id, dependency.Type)
}

func (dependency *TaskDependency) UnmarshalJSON(bytes []byte) (error error) {
    var temp []interface{}

    if error = json.Unmarshal(bytes, &temp); error == nil {
        dependency.Id = uint32(temp[0].(float64));
        dependency.Type = temp[1].(string);
    }

    return error;
}

type TaskResource struct {
    Id uint32
    Dedication float32
}

func (resource *TaskResource) String() string {
    return fmt.Sprintf("[%d] %f", resource.Id, resource.Dedication)
}

func (resource *TaskResource) UnmarshalJSON(bytes []byte) (error error) {
    var temp []interface{}

    if error = json.Unmarshal(bytes, &temp); error == nil {
        resource.Id = uint32(temp[0].(float64));
        resource.Dedication = float32(temp[1].(float64));
    }

    return error;
}

type TaskEffort struct {
    Id uint32
    Timestamp Time
    Effort Duration 
}

func (effort *TaskEffort) String() string {
    return fmt.Sprintf("[%d] %s %s", effort.Id, time.Time(effort.Timestamp), time.Duration(effort.Effort))
}

func (effort *TaskEffort) UnmarshalJSON(bytes []byte) (error error) {
    var temp []interface{}

    if error = json.Unmarshal(bytes, &temp); error == nil {
        effort.Id = uint32(temp[0].(float64));
        effort.Timestamp = Time(time.Unix(int64(temp[1].(float64)), 0));
        effort.Effort = Duration(time.Duration(int64(temp[2].(float64)) * 1000000000));
    }

    return error;
}

type Task struct {
    Object
    Closed bool
    StartsOn Time `json:"starts_on"`
    
    PlannedStart Time `json:"planned_start"`
    PlannedEnd Time `json:"planned_end"`
    PlannedDuration Duration `json:"planned_duration"`

    EstimatedStart Time `json:"estimated_start"`
    EstimatedEnd Time `json:"estimated_end"`
    EstimatedDuration Duration `json:"estimated_duration"`

    PlannedEffort Duration `json:"planned_effort"`
    EstimatedEffort Duration `json:"estimated_effort"`
    CurrentEffort Duration `json:"current_effort"`

    EstimatedProgress float32 `json:"estimated_progress"`
    PlannedProgress float32 `json:"planned_progress"`
    
    OutgoingDependencies []TaskDependency `json:"outgoing_dependencies"`
    Resources []TaskResource `json:"resources"`
    Efforts []TaskEffort `json:"efforts"`
}

func (task *Task) String() string {
    var fields = []string {
        "Type: Task",
        task.Object.String(),
        "Closed: " + strconv.FormatBool(task.Closed),
        "Planned start: " + time.Time(task.PlannedStart).String(),
        "Planned end: " + time.Time(task.PlannedEnd).String(),
        "Planned duration: " + time.Duration(task.PlannedDuration).String(),
        "Estimated start: " + time.Time(task.EstimatedStart).String(),
        "Estimated end: " + time.Time(task.EstimatedEnd).String(),
        "Estimated duration: " + time.Duration(task.EstimatedDuration).String(),
        "Planned effort: " + time.Duration(task.PlannedEffort).String(),
        "Estimated effort: " + time.Duration(task.EstimatedEffort).String(),
        "Current effort: " + time.Duration(task.CurrentEffort).String(),
        "Estimated progress: " + strconv.FormatFloat(float64(task.EstimatedProgress), 'f', 2, 64),
        "Planned progress: " + strconv.FormatFloat(float64(task.PlannedProgress), 'f', 2, 64),
    }
    var dependencies = []string { "Tasks:" }
    var resources = []string { "Assigned resources:" }
    var efforts = []string { "Efforts:" }
    
    for _, value := range task.OutgoingDependencies {
        dependencies = append(dependencies, "\t" + value.String())
    }
    
    for _, value := range task.Resources {
        resources  = append(resources, "\t" + value.String())
    }
    
    for _, value := range task.Efforts {
        efforts  = append(efforts, "\t" + value.String())
    }
    
    return strings.Join(append(append(append(fields, dependencies...), resources...), efforts...), "\n")
}

func (project *Project) Tasks() (tasks []*Task, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/tasks", nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&tasks); error != nil {
                return nil, error
            } else {
                return tasks, nil
            }
        }
    }
}

func (project *Project) Task(id uint32) (task *Task, error error) {
    if request, error := project.session.makeRequest("GET", "/api/projects/" + project.Identifier + "/tasks/" + strconv.FormatUint(uint64(id), 10), nil); error != nil {
        return nil, error
    } else {
        if response, error := project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&task); error != nil {
                return nil, error
            } else {
                return task, nil
            }
        }
    }
}
