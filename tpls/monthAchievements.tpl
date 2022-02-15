
\parbox[t][][t]{0.48\myLenTwoCol}{
\vspace{0.1em}
\begin{tabular}{ r l }
{{- range $i, $action := .Month.ActionsLeft }}
{{$action.Action}}
{{- end -}}
\end{tabular}
}
\parbox[t][][t]{0.48\myLenTwoCol}{
\vspace{0.1em}
\begin{tabular}{ r l }
{{- range $i, $action := .Month.ActionsRight }}
{{$action.Action}}
{{- end -}}
\end{tabular}
}
