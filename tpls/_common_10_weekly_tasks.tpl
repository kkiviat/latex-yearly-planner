{{- $days := .Body.Week.Days -}}
{{- $day1 := index $days 0 -}}
{{- $day2 := index $days 1 -}}
{{- $day3 := index $days 2 -}}
{{- $day4 := index $days 3 -}}
{{- $day5 := index $days 4 -}}
{{- $day6 := index $days 5 -}}
{{- $day7 := index $days 6 -}}

\Repeat{3}{%
{{`\begin{tabularx}{\linewidth}{p{0.455\textwidth}|*{7}{m{2pt}}|llll}`}}
  \hline
{{`\rowcolor{black}{\textcolor{white}{\bfseries Task}} \myLineHeightButLine{}& \multicolumn{7}{l|}{{\textcolor{white}{\bfseries Worked On}}} & \multicolumn{4}{l}{{\textcolor{white}{\bfseries Status}}}        \\`}}
\hline
     \myLineHeightButLine{}&  {{- $day1.WeekLetterLink -}} & {{- $day2.WeekLetterLink -}}  & {{- $day3.WeekLetterLink -}}  & {{- $day4.WeekLetterLink -}}  & {{- $day5.WeekLetterLink -}} & {{- $day6.WeekLetterLink -}} & {{- $day7.WeekLetterLink -}} & {\tiny Started} & {\tiny Halfway}& {\tiny Almost}& {\tiny Done}\\%
     \hline
\multicolumn{12}{l}{\cellcolor{lightgray}Subtasks\myLineHeightButLine{}}\\%
\hline

  \Repeat{\myNumWeeklyTaskTodos}{%
\multicolumn{12}{l}{$\square$\myLineHeightButLine{}}\\%
\hline}%
\end{tabularx}%

\bigskip
}


\vfill
