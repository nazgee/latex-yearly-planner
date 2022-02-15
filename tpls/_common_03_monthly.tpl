{{- template "monthTabularV2.tpl" dict "Month" .Body.Month "Large" true -}}
\medskip

{{ if $.Cfg.Dotted -}}
\myUnderline{Notatki}
\vbox to 0pt{\myMash[\myMonthlySpring]{20}{\myNumDotWidthFull}}
{{- else -}}
\parbox[t][][t]{\myLenTwoCol}{
  \myUnderline{Notatki}
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }%
}%
\hspace{\myLenTwoColSep}%
\parbox[t][][t]{\myLenTwoCol}{
  \myUnderline{Treniol}
    {{- template "monthAchievements.tpl" dict "Month" .Body.Month "Large" true -}}
}
%{{- end}}
