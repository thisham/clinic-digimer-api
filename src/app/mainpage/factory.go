package mainpage

import "digimer-api/src/app/mainpage/handlers"

func InitialPageFactory() handlers.Handler {
	return *handlers.NewMainPageHandler()
}
