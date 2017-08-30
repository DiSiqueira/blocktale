package tale

import "net/http"

type (
	PostTaleHandler struct {

	}
)

func NewPostTaleHandler() *PostTaleHandler {
	return &PostTaleHandler{}
}

func (g *PostTaleHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}