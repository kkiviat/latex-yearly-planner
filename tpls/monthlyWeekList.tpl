{%
\setlength{\tabcolsep}{\myLenTabColSep}%
%
{{ .Month.DefineWeekTable .TableType }}
  {{- range $i, $week := .Month.Weeks }}
  {{$week.WeekNumberLink -}}&
  {{ end }}
  {{ .Month.EndTable .TableType -}}
}
