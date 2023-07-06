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
			background-color: #f2f2f2;
			color: #222;
			margin: 0;
			padding: 0;
		}

		hr {
      border: 0;
      border-top: 2px solid #ddf;
      max-width: 75ch;
      margin: 0 auto;
		}

    a {
      text-decoration: none;
      color: #33d;
      border-bottom: 1px solid #33d;
    }

    a:hover, a:visited {
      color: #fa0;
      border-color: #fa0;
    }

		.deco-bar {
			position: absolute;
			width: 100%;
			height: 64px;
			background-color: #88c;
			padding: 0;
			margin: 0;
		}

		.content {
			position: relative;
			padding: 32px 16px;
			font-size: 18px;
			font-family: Inconsolata, sans-serif, arial;
      font-weight: 200;
		}

		.content-inner {
			position: relative;
			max-width: 960px;
			margin: 0 auto;
			padding: 1em;
			background-color: #fefefe;
			box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
			border-radius: 8px 8px 4px 4px;
			border-top: 4px solid #fa0;
		}

		.h-card img {
			height: 1.1em;
			vertical-align: -0.2em;
		}

    .external a {
      border: 0;
    }

		.external img {
			height: 1.1em;
			vertical-align: -0.2em;
		}

    .sections {
      padding-bottom: 1em;
      max-width: 75ch;
      margin: 0 auto;
      line-height: 130%;
    }

		.section_list {
      padding-left: 1em;
			margin: 0;
		}

    .section_content {
      padding: 1em 0;
      max-width: 75ch;
      margin: 0 auto;
      line-height: 130%;
    }

    .section_footer {
      padding-top: 1em;
      max-width: 75ch;
      margin: 0 auto;
      line-height: 130%;
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
        <section class="sections">
          /<a href="/">home</a>{{template "section_title" .}}
        </section>

				<hr />

				<section class="section_content">
					{{template "section_content" .}}
				</section>

				<hr />

        <section class="section_footer">
          <p class="h-card">
            <img class="u-photo" src="/static/me.png" alt="photo of Per Lönn Wege" />
            <a class="p-name u-url" href="https://perlw.se">Per Lönn Wege</a>
            <a class="u-email" href="mailto:per@perlw.se">per@perlw.se</a>
          </p>

          <p class="external">
            <a href="https://github.com/perlw" rel="me"><img src="/static/external/GitHub-Mark-32px.png" alt="github logo" /></a>
          </p>
        </section>
			</div>
		</div>
	</body>
</html>
