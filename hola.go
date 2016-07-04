package hola

import (
      "html/template"
      "net/http"
      "io/ioutil"
)


type Page struct{
	Title string
	Body []byte
}

func loadPage(title string) (*Page, error){
   filename := title + ".txt"
   body, err := ioutil.ReadFile(filename)
   if err != nil {
	return nil, err
   }

 return &Page{Title: title, Body: body},nil
}

func init(){
     http.HandleFunc("/edit/", editHandler)
}


func editHandler(w http.ResponseWriter, r *http.Request){
   title := r.URL.Path[len("/edit/"):]
   p, err := loadPage(title)
  if err != nil {
      p  =  &Page{Title: title}
   }
  t, _ := template.ParseFiles("static/edit.html")
  t.Execute(w,p)

}
