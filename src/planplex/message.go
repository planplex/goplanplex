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
    "fmt"
    "sort"
    "bytes"
)

type Message struct {
    Id uint32
    User string
    Text string `json:"message"`
    Timestamp Time `json:"date"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
}

func (message *Message) String() string {
    return fmt.Sprintf("[%s %s (@%s)] (%s) %s", message.FirstName, message.LastName, message.User, time.Time(message.Timestamp), message.Text)
}

type MessageArray []*Message

func (messages MessageArray) Less(i, j int) bool {
    return time.Time(messages[i].Timestamp).Before(time.Time(messages[j].Timestamp))
}

func (messages MessageArray) Len() int {
    return len(messages)
}

func (messages MessageArray) Swap(i, j int) {
    messages[i], messages[j] = messages[j], messages[i] 
}

func (task *Task) Messages() (messages []*Message, error error) {
    if request, error := task.project.session.makeRequest("GET", "/api/projects/" + task.project.Identifier + "/chats/" + strconv.FormatUint(uint64(task.Id), 10), nil); error != nil {
        return nil, error
    } else {
        if response, error := task.project.session.client.Do(request); error != nil {
            return nil, error
        } else if response.StatusCode != http.StatusOK {
            return nil, errors.New(response.Status)
        } else {
            if error := json.NewDecoder(response.Body).Decode(&messages); error != nil {
                return nil, error
            } else {
                sort.Sort(MessageArray(messages))

                return messages, nil
            }
        }
    }
}

func (task *Task) PostMessage(message string) (error error) {
    data, _ := json.Marshal(map[string]string { "message": message })
    url := "/api/projects/" + task.project.Identifier + "/chats/" + strconv.FormatUint(uint64(task.Id), 10)
    
    if request, error := task.project.session.makeRequest("POST", url, bytes.NewReader(data)); error != nil {
        return error
    } else {
        if response, error := task.project.session.client.Do(request); error != nil {
            return error
        } else if response.StatusCode != http.StatusOK {
            return errors.New(response.Status)
        } else {
            return nil
        }
    }
}
