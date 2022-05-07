package factories

import (
	"digimer-api/src/app/mainpage"
	mainpageHandler "digimer-api/src/app/mainpage/handlers"
)

type webHandler struct {
	InitialPageHandler mainpageHandler.Handler
}

func WebInit() webHandler {
	return webHandler{
		InitialPageHandler: mainpage.InitialPageFactory(),
	}
}
