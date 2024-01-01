{{- $days := .Body.Week.Days -}}
{{- $day1 := index $days 0 -}}
{{- $day2 := index $days 1 -}}
{{- $day3 := index $days 2 -}}
{{- $day4 := index $days 3 -}}
{{- $day5 := index $days 4 -}}
{{- $day6 := index $days 5 -}}
{{- $day7 := index $days 6 -}}

\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day1.WeekShortLink -}} }}%
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day2.WeekShortLink -}} }}%
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day3.WeekShortLink -}} }}
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day4.WeekShortLink -}} }}%
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day5.WeekShortLink -}} }}%
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day6.WeekShortLink -}} }}
\hspace{\myLenTriColSep}%
\parbox{\dimexpr\linewidth/7-\myLenTriColSep}{\myUnderline{ {{- $day7.WeekShortLink -}} }}%

{{ template "_common_09_notes.tpl" dict "Cfg" .Cfg "Body" .Body }}
