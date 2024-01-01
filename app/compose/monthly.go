package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

func Monthly(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 12)

	for _, quarter := range year.Quarters {
		for _, month := range quarter.Months {
 	    	       notesLink := header.NewTextItem("Notes").RefText("Notes Index")
		       extra := header.Items{notesLink}
		       extra = append(extra, month.PrevNext()...)
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[0],
				Body: map[string]interface{}{
					"Year":         year,
					"Quarter":      quarter,
					"Month":        month,
					"Breadcrumb":   month.Breadcrumb(cfg),
					"HeadingMOS":   month.HeadingMOS(),
					"SideQuarters": year.SideQuarters(quarter.Number),
					"SideMonths":   year.SideMonths(month.Month),
					"Extra":        extra.WithTopRightCorner(cfg.ClearTopRightCorner),
					"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, nil, 0),
				},
			})
		}
	}

	return modules, nil
}
