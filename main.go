package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/vitush93/go-web/app/model"
)

var templates map[string]*template.Template

type TemplateData struct {
	Title string
	Data  interface{}
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	includes, err := filepath.Glob("templates/includes/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range includes {
		name := filepath.Base(tpl)

		templates[name] = template.Must(template.ParseFiles(
			tpl,
			"templates/layout.html",
		))
	}
}

func loadConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic("failed to load configuration")
	}
}

func main() {
	loadConfig()
	loadTemplates()

	dbhost := viper.GetString("dev.mysql.host")
	dbport := viper.GetInt("dev.mysql.port")
	dbuser := viper.GetString("dev.mysql.user")
	dbpass := viper.GetString("dev.mysql.password")
	dbname := viper.GetString("dev.mysql.database")

	model.Connect(dbhost, dbport, dbuser, dbpass, dbname)
	defer model.DB.Close()

	model.Migrate()

	router := httprouter.New()
	router.GET("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts := model.FindAllPosts()

	render(w, "posts.html", &TemplateData{Title: "Hello World", Data: posts})
}

func render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		log.Fatal("compiled template " + name + " not found")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Fatal(err)
	}
}
