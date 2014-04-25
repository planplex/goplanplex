/*
 * Copyright (C) 2014 Planplex
 * License: http://www.gnu.org/licenses/gpl.html GPL version 2 or higher
 */

package planplex

import(
    "strconv"
    "time"
)

type Time time.Time

func (_time *Time) UnmarshalJSON(b []byte) error {
    if epoch, error := strconv.ParseInt(string(b), 0, 64); error != nil {
        return error
    } else {
        *_time = Time(time.Unix(epoch, 0))

        return nil
    }
}

type Duration time.Duration

func (_duration *Duration) UnmarshalJSON(bytes []byte) error {
    if count, error := strconv.ParseInt(string(bytes), 0, 64); error != nil {
        return error
    } else {
        *_duration = Duration(time.Duration(count * 1000000000))

        return nil
    }
}
