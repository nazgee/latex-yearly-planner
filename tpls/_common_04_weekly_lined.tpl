{{- $days := .Body.Week.Days -}}
{{- $day1 := index $days 0 -}}
{{- $day2 := index $days 1 -}}
{{- $day3 := index $days 2 -}}
{{- $day4 := index $days 3 -}}
{{- $day5 := index $days 4 -}}
{{- $day6 := index $days 5 -}}
{{- $day7 := index $days 6 -}}

\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day1.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}%
\hspace{\myLenTriColSep}%
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day2.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}%
\hspace{\myLenTriColSep}%
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day3.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}
\vfill
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day4.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}%
\hspace{\myLenTriColSep}%
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day5.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}%
\hspace{\myLenTriColSep}%
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day6.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}
\vfill
\parbox{\myLenTriCol}{%
  \userActionsUnderline{ {{- $day7.WeekLink -}} }{\userActions}{0.55\myLenTriCol}{0.45\myLenTriCol}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}%
\hspace{\myLenTriColSep}%
\parbox{\dimexpr2\myLenTriCol+\myLenTriColSep}{%
  \myUnderline{Notatki\textcolor{white}{g}}\Repeat{\myNumWeeklyLines}{\myLineGrayVskipTop}%
}
