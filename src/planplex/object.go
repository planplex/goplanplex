/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

import(
    "strconv"
    "strings"
)

type Object struct {
    Id uint32
    Name string
    Description string
}

func (object *Object) String() string {
    var fields = []string {
        "Id: " + strconv.FormatUint(uint64(object.Id), 10),
        "Name: " + object.Name,
        "Description: " + object.Description,
    }
    
    return strings.Join(fields, "\n")
}

