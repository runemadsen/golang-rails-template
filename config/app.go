package config

import (
  "html/template"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "runtime"
  "path/filepath"
)

func Root() string {
  _, file, _, _ := runtime.Caller(0)
  return filepath.Join(filepath.Dir(file), "../")
}

func CreateApplication() *martini.ClassicMartini {

  m := martini.Classic()

  // Add contrib renderer. See options in docs.
  // https://github.com/martini-contrib/render
  m.Use(render.Renderer(render.Options{
    Directory: filepath.Join(Root(), "app/views"),
    Layout: "layouts/layout",
    Extensions: []string{".html"},
    Funcs: []template.FuncMap{AssetHelpers()},
  }))

  // Call the initializers
  Initialize(m)

  // Call the routes
  Routes(m)

  return m

}