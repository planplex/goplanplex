/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

type Role int

const (
    Unknown Role = 0
    Manager Role = 1
    Participant Role = 2
    Observer Role = 3
)

func (role Role) String() string {
    switch role {
        case Unknown: return "Unknown"
        case Manager: return "Manager"
        case Participant: return "Participant"
        case Observer: return "Observer"
        default: return "<>"
    }
}

