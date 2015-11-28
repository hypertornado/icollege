package server

import (
	"github.com/hypertornado/prago"
)

func Start(port int, developmentMode bool) {

	var err error
	app := prago.NewApp()
	//cookieStore := sessions.NewCookieStore([]byte("rand"))

	translations := prago.NewI18N()

	templates, err := prago.LoadTemplates(
		[]string{
			"server/templates/*.tmpl",
		}, translations)
	if err != nil {
		panic(err)
	}

	app.AddTemplates(templates)

	app.Route(prago.GET, "/", app.MainController(), HomeRedirect)

	app.Route(prago.POST, "/newsletter", app.MainController(), Newsletter)

	err = app.ListenAndServe(port, developmentMode)
	if err != nil {
		panic(err)
	}
}

func HomeRedirect(request prago.Request) {
	prago.Redirect(request, "/index.html")
}

func Newsletter(request prago.Request) {
	email := request.Params().Get("email")
	println(email)
	prago.Redirect(request, "/index.html?submitted")
}
