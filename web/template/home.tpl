{{define "section_title"}}
	<ul class="section_list">
		<li>
			<a href="/musings">Musings</a>
		</li>
		<!--
		<li>
			<a href="/til">TIL</a>
		</li>
		<li>
			<a href="/why">Why?</a>
		</li>
		-->
	</ul>
{{end}}

{{define "section_content"}}
This site is about me. It is me. This is me. I am here and this is where you found me.

{{if .wakaStats}}

<hr />

<div class="wakastats">
Latest 7 days:&nbsp;
{{range $k,$v := .wakaStats}}
{{$k}}:{{$v}}%&nbsp;
{{end}}
</div>
{{end}}

{{end}}
