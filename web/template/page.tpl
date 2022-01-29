<!DOCTYPE html>
<html lang="en">
	<link rel="icon" type="image/png" href="/static/logo.png" />
	<link rel="apple-touch-icon-precomposed" type="image/png" href="/static/logo.png" />

	<head>
		<meta charset="UTF-8">
		<title>Ergo</title>

		<style>
		body {
			background-color: #eef;
		}

		hr {
			border: 0;
			border-top: 1px dashed #aae;
			margin: 4px 0;
		}

		.content {
			position: relative;
			width: 960px;
			margin: 0 auto;
			padding: 4px;
			background-color: #fff;
			box-shadow: 2px 2px 2px #ddf;
		}

		.h-card img {
			height: 1.1em;
			vertical-align: -0.2em;
		}

		.external img {
			height: 1.1em;
			vertical-align: -0.2em;
		}

		.section_list {
			padding-left: 1em;
			margin: 0;
		}
		</style>
	</head>
	<body>
		<div class="content">
			/<a href="/">home</a>{{template "section_title" .}}

			<hr />

			<div class="section_content">
				{{template "section_content" .}}
			</div>

			<hr />

			<p class="h-card">
				<img class="u-photo" src="/static/me.png" alt="photo of Per Lönn Wege" />
				<a class="p-name u-url" href="https://perlw.se">Per Lönn Wege</a>
				<a class="u-email" href="mailto:per@perlw.se">per@perlw.se</a>
			</p>

			<p class="external">
				<a href="https://github.com/perlw" rel="me"><img src="/static/external/GitHub-Mark-32px.png" alt="github logo" /></a>
				<a href="https://keybase.io/perlw" rel="me"><img src="/static/external/Keybase_logo_official.png" alt="keybase logo" /></a>
				<a href="https://twitter.com/perlwege" rel="me"><img src="/static/external/Twitter social icons-circle-blue.png" alt="twitter logo" /></a>
			</p>

		</div>
	</body>
</html>
