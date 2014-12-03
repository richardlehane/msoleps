package types

import "time"

// http://msdn.microsoft.com/en-us/library/cc237601.aspx
type Date float64

func (d Date) Time() time.Time {
	start := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	day := float64(time.Hour * 24)
	dur := time.Duration(day * float64(d))
	return start.Add(dur)
}

func (d Date) String() string {
	return d.Time().String()
}
