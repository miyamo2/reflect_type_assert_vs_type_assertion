package reflect_type_assert_vs_type_assertion

import "time"

type ITime interface {
	String() string
	GoString() string
	Format(layout string) string
	AppendFormat(b []byte, layout string) []byte
	IsZero() bool
	After(u time.Time) bool
	Before(u time.Time) bool
	Compare(u time.Time) int
	Equal(u time.Time) bool
	Date() (year int, month time.Month, day int)
	Year() int
	Month() time.Month
	Day() int
	Weekday() time.Weekday
	ISOWeek() (year int, week int)
	Clock() (hour int, min int, sec int)
	Hour() int
	Minute() int
	Second() int
	Nanosecond() int
	YearDay() int
	Add(d time.Duration) time.Time
	Sub(u time.Time) time.Duration
	AddDate(years int, months int, days int) time.Time
	UTC() time.Time
	Local() time.Time
	In(loc *time.Location) time.Time
	Location() *time.Location
	Zone() (name string, offset int)
	ZoneBounds() (start time.Time, end time.Time)
	Unix() int64
	UnixMilli() int64
	UnixMicro() int64
	UnixNano() int64
	AppendBinary(b []byte) ([]byte, error)
	MarshalBinary() ([]byte, error)
	GobEncode() ([]byte, error)
	MarshalJSON() ([]byte, error)
	AppendText(b []byte) ([]byte, error)
	MarshalText() ([]byte, error)
	IsDST() bool
	Truncate(d time.Duration) time.Time
	Round(d time.Duration) time.Time
}
