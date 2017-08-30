package tale

import "net/http"

type (
	GetTaleHandler struct {

	}
)

func NewGetTaleHandler() *GetTaleHandler {
	return &GetTaleHandler{}
}

func (g *GetTaleHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}