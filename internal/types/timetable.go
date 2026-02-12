package types

import (
	"fmt"
	"time"
)

type TimeTable struct {
	Rows  map[int]*TimeTableRow
	Dates []time.Time
}

func (t *TimeTable) GenerateDates(now time.Time) {

	for i := 13; i >= 0; i-- {
		day := truncateToLocalDay(now.AddDate(0, 0, -i))
		t.Dates = append(t.Dates, day)
	}
}

func truncateToLocalDay(t time.Time) time.Time {
	t = t.Local()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func (t *TimeTable) GetRow(projectid int, now time.Time) *TimeTableRow {

	row, exists := t.Rows[projectid]
	if !exists {
		row = &TimeTableRow{}

		for i := 13; i >= 0; i-- {
			day := truncateToLocalDay(now.AddDate(0, 0, -i))
			newcell := &TimeTableCell{}
			newcell.Date = day
			row.Cells = append(row.Cells, newcell)
		}
		t.Rows[projectid] = row

	}

	return row

}

type TimeTableRow struct {
	ProjectName string
	Cells       []*TimeTableCell
	Color       string
}

func (r *TimeTableRow) GetCell(date time.Time) *TimeTableCell {
	date = truncateToLocalDay(date)

	for _, cell := range r.Cells {
		if cell.Date.Equal(date) {
			return cell
		}
	}
	return nil

}

type TimeTableCell struct {
	Date      time.Time
	TotalTime int //seconds
}

func (c *TimeTableCell) Hours() string {
	hours := float64(c.TotalTime) / 3600
	return fmt.Sprintf("%.1f", hours)
}

func (c *TimeTableCell) IsWeekday() bool {
	day := c.Date.Weekday()
	return day != time.Saturday && day != time.Sunday
}
