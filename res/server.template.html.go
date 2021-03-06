package res

const ServerTemplate = `<!DOCTYPE HTML>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Status page</title>
	<style type="text/css">
		div {
			margin: 1em;
		}
	</style>
	<script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.0.2/jquery.min.js"></script>
	<link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.2.2/css/bootstrap-combined.min.css" rel="stylesheet">
</head>
<body>
	<div>
		<h1>{{.Hostname}}</h1>
		<p class="muted">{{.GoVersion}}</p>
	</div>

	<div>
		<h2>pigen.dk</h2>
		<form action="" method="post" id="make_thumbnails">
			<input type="submit" class="btn" value="Gendan billeder" />
		</form>
	</div>

	<script type="text/javascript">
		var checks = {
				landscapeSysinfo: ["/landscape/sysinfo", "landscape sysinfo", 2000, "span8"],
		};

		$(function () {

			var $form = $("form#make_thumbnails"),
				$button = $form.find("input:submit");

			// Handle form submit
			$form.submit(function (e) {

				$button.removeClass("btn-inverse");
				$button.addClass("disabled");
				$button.attr("disabled", "disabled");

				$.post("/action", {action: "make_thumbnails"}).always(function () {
					$button.removeClass("disabled");
					$button.removeAttr("disabled");
					$button.addClass("btn-inverse");
				});
				e.preventDefault()
				return false;
			});

				var $body = $(document.body),
					pre, data;

				for (var id in checks) {
					data = checks[id];
					pre = $('<div id="'+id+'"><h2>'+data[1]+'</h2><pre>loading...</pre></div>').appendTo(document.body).find("pre");

					(function  (p, id, data) {
						var f = function() {
							p.load(data[0]);
						};
						setInterval(f, data[2]);
					})(pre, id, data);
				}
			});

			// setTimeout(function () {window.location.reload();}, 20000);
	</script>
</body>
</html>`
