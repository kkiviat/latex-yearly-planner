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
	    hasWeeklyJournal := false
	    for _, page := range cfg.Pages {
		if page.Name == "weekly_tasks" {
		    hasWeeklyTasks = true
		}
		if page.Name == "weekly_journal" {
		    hasWeeklyJournal = true
		}
	    }

 	    notesLink := header.NewTextItem("Notes").RefText("Notes Index")
	    for _, week := range year.Weeks {
	            extra := header.Items{notesLink}
	    	    if prefix == "" {
			if hasWeeklyTasks {
			    extra = append(extra, header.NewTextItem("Tasks").RefText(week.RefText("Tasks")))
			}
			if hasWeeklyJournal {
			    extra = append(extra, header.NewTextItem("Journal").RefText(week.RefText("Journal")))
			}
		    } else {
		            dayLayout := "2"
 			    dayStart := week.Days[0].Time.Format(dayLayout)
 			    dayEnd := week.Days[6].Time.Format(dayLayout)
		       	    dayItemStart := header.NewTextItem(dayStart + "-" + dayEnd)
		            extra = header.Items{dayItemStart}
                    }
		    extra = append(extra, week.PrevNext(prefix, true)...)
		    extra2_other := header.NewCellItem("Tasks").Refer(week.RefText("Tasks")).Selected(prefix == "Tasks")
		    extra2_week := week
		    if prefix == "" {
		       extra2_week = nil
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
				    "Extra2":       extra2(cfg.ClearTopRightCorner, false, false, extra2_week, &extra2_other, 0),
			    },
		    })
	    }

	    return modules, nil
    }
}
