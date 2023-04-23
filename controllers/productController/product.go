package productcontroller

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	products := productModel.GetAll()
	data := map[string]any{
		"products": products,
	}
	temp, err := template.ParseFiles("views/products/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
func Detail(w http.ResponseWriter, r *http.Request) {

}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
