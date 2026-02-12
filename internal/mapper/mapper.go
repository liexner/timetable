package mapper

import (
	"time"
	"timetable/internal/types"
)

func BuildTimeTable(time_entries []types.TimeEntry, projects []types.Project) types.TimeTable {
	now := time.Now()

	projectMap := make(map[int]types.Project)
	for _, p := range projects {
		projectMap[p.ID] = p
	}

	timetable := types.TimeTable{Rows: make(map[int]*types.TimeTableRow)}
	timetable.GenerateDates(now)

	for _, entry := range time_entries {

		row := timetable.GetRow(entry.ProjectID, now)

		if row.ProjectName == "" {
			if p, ok := projectMap[entry.ProjectID]; ok {
				row.ProjectName = p.Name
				row.Color = p.Color
			}
		}

		cell := row.GetCell(entry.Start)
		if cell != nil {
			cell.TotalTime += entry.Duration
		}
	}

	return timetable

}
