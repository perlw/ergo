{{define "section_title"}}/musings{{end}}

{{define "section_content"}}
<ul>
	{{range $date, $links := .musingLinks}}
		<li>
		{{$date}}
			<ul>
			{{range $links}}
				<li><a href="/musing/{{$date}}/{{.}}">{{.}}</a></li>
			{{end}}
			</ul>
		</li>
	{{end}}
</ul>
{{end}}
