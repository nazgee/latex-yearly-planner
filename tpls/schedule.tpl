\myUnderline{Plan dnia\textcolor{white}{g}}\vskip-\myLenLineThicknessDefault
{{range $hour := .Day.Hours .Cfg.Layout.Numbers.DailyBottomHour .Cfg.Layout.Numbers.DailyTopHour -}}
\myLineHeightButLine%
{{if $.Cfg.AMPMTime -}}
\parbox{9mm}{\hfill\small {{- $hour.FormatHour $.Cfg.AMPMTime -}} }%
{{- else -}}
{\small {{- $hour.FormatHour $.Cfg.AMPMTime -}} }
{{- end}}
\myLineLightGray\vskip\myLenLineHeightButLine\myLineGray
{{- end}}
{{if $.Cfg.AddLastHalfHour}}\vskip\myLenLineHeightButLine\vbox to 0pt{\myLineLightGray}{{end}}
