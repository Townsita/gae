package gae

import (
	"github.com/Townsita/data-adapter-dummy"
	"github.com/Townsita/townsita"
	"net/http"
)

func init() {
	da := dummy.New()
	c := townsita.NewConfig()
	handler := townsita.New(c, da).GetHTTPHandler()
	http.Handle("/", handler)
}
