package router

import "testing"

func TestAppRouter_InitializeRouters(t *testing.T) {
	AppRouter.InitializeRouters()

	if AppRouter.Router == nil {
		t.Fatal("Mux router is supposed to be initialized")
	}
}
