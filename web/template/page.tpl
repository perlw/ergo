<!DOCTYPE html>
<html lang="en">
	<link rel="icon" type="image/png" href="/static/logo.png" />
	<link rel="apple-touch-icon-precomposed" type="image/png" href="/static/logo.png" />
	<meta name="viewport" content="width=device-width, user-scalable=yes">

	<head>
		<meta charset="UTF-8">
		<title>Ergo</title>

		<style>
		body {
			background-color: #fafaff;
			margin: 0;
			padding: 0;
		}

		hr {
			border: 0;
			border-top: 1px dashed #aae;
			margin: 4px 0;
		}

		.deco-bar {
			position: absolute;
			width: 100%;
			height: 64px;
			background-color: #aaf;
			padding: 0;
			margin: 0;
		}

		.content {
			position: relative;
			padding: 32px 16px;
			font-size: 16px;
			font-family: sans-serif, arial;
		}

		.content-inner {
			position: relative;
			max-width: 960px;
			margin: 0 auto;
			padding: 16px;
			background-color: #fff;
			box-shadow: 2px 2px 1px rgba(0, 0, 0, 0.1);
			border-radius: 8px;
			border: 1px solid #ddf;
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

		.wakastats {
			font-size: 0.75em;
		}
		</style>
	</head>
	<body>
		<div class="deco-bar"></div>
		<div class="content">
			<div class="content-inner">
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
		</div>
	</body>
</html>
