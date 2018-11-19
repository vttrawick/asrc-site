package router

import(
	"net/http"
	"fmt"
//	"html/template"
//	"github.com/vttrawick/asrc-site/paths"
)

// the cluster page
func Cluster(w http.ResponseWriter, r *http.Request) {
	//	t, err := loadTemplates()
	fmt.Fprintf(w, "this is the cluster page!")
}

//func loadTemplates() (*template.Template, error) {
//
//}
