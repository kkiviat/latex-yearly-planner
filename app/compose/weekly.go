package compose

import (
        // "time"
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

var Weekly = WeeklyStuff("", "")
var WeeklyTasks = WeeklyStuff("Tasks", "Tasks")
var WeeklyJournal = WeeklyStuff("Journal", "Journal")

func WeeklyStuff(prefix, leaf string) func(cfg config.Config, tpls []string) (page.Modules, error) {
	return func(cfg config.Config, tpls []string) (page.Modules, error) {
	    modules := make(page.Modules, 0, 53)
	    year := cal.NewYear(cfg.WeekStart, cfg.Year)

	    hasWeeklyTasks := false
	    hasNotes := false
	    hasWeeklyJournal := false
	    for _, page := range cfg.Pages {
		if page.Name == "weekly_tasks" {
		    hasWeeklyTasks = true
		}
		if page.Name == "weekly_journal" {
		    hasWeeklyJournal = true
		}
		if page.Name == "notes_indexed" {
		    hasNotes = true
		}
	    }

	    for _, week := range year.Weeks {
	            extra := week.PrevNext(prefix, false)
	    	    if prefix == "" && (hasNotes || hasWeeklyTasks || hasWeeklyJournal) {
			    if hasNotes {
			        extra = append(extra, header.NewTextItem("Notes").RefText("Notes Index"))
			    }
			    if hasWeeklyTasks {
			        extra = append(extra, header.NewTextItem("Tasks").RefText(week.RefText("Tasks")))
		            }
			    if hasWeeklyJournal {
			        extra = append(extra, header.NewTextItem("Journal").RefText(week.RefText("Journal")))
		            }
		    } else if prefix != "" {
			extra = week.PrevNext(prefix, true)
                    }
		    modules = append(modules, page.Module{
			    Cfg: cfg,
			    Tpl: tpls[0],
			    Body: map[string]interface{}{
				    "Year":         year,
				    "Week":         week,
				    "Breadcrumb":   week.Breadcrumb(prefix, leaf, cfg),
				    "HeadingMOS":   week.HeadingMOS(prefix, leaf),
				    "SideQuarters": year.SideQuarters(week.Quarters.Numbers()...),
				    "SideMonths":   year.SideMonths(week.Months.Months()...),
				    "Extra":        extra.WithTopRightCorner(cfg.ClearTopRightCorner),
				    "Extra2":       extra2(cfg.ClearTopRightCorner, false, false, week, 0),
			    },
		    })
	    }

	    return modules, nil
    }
}
