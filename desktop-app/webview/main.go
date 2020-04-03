package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zserge/webview"
	"log"
)

const addr = "127.0.0.1:8080"

/*
	# Windows requires special linker flags for GUI apps.
	# It's also recommended to use TDM-GCC-64 compiler for CGo.
	# http://tdm-gcc.tdragon.net/download
	#
	#
	$ go build -ldflags="-H windowsgui" -o myapp.exe # build for windows
	$ ./myapp.exe # run
	#
	#
	# Note: if you see "use option -std=c99 or -std=gnu99 to compile your code"
	# please refer to: https://github.com/zserge/webview/issues/188
*/
func main() {
	//go runServer()
	//showAndWaitWindow()
	Example()
}

func runServer() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello Desktop</h1>")
	})
	app.Run(iris.Addr(addr))
}

func showAndWaitWindow() {
	//webview.Open("My App", addr, 800, 600, true)
}

func Example() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Hello")
	w.Bind("noop", func() string {
		log.Println("hello")
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Navigate(`data:text/html,
		<!doctype html>
		<html>
			<body>hello</body>
			<script>
				window.onload = function() {
					document.body.innerText = ` + "`hello, ${navigator.userAgent}`" + `;
					noop().then(function(res) {
						console.log('noop res', res);
						add(1, 2).then(function(res) {
							console.log('add res', res);
							quit();
						});
					});
				};
			</script>
		</html>
	)`)
	w.Run()
}
