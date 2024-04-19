package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct { //Структура Пользователя
	Name                string
	Age                 uint8
	Influence           int16
	Power, Intelligence float32
}

func (u *User) GetAllinfo() string {
	return fmt.Sprintf("User name: [%s]. User is [%d] years old and he has power: [%f] ", u.Name, u.Age, u.Power)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(page http.ResponseWriter, _ *http.Request) { //Стартовая страница
	LorgarAurelian := User{"Lorgar", 56, 2600, 42.4, 112.3}
	//LorgarAurelian.setNewName("Lorgar Lider")
	//_, _ = fmt.Fprintf(page, `<h1>warhammer 40'000 </h1>`)
	//_, _ = fmt.Fprintf(page, "<b1>Lord: "+Lorg.Name+"</b1>")
	//_, _ = fmt.Fprintf(page, "<b2> "+Lorg.getAllinfo()+"</b2>")
	tmpl, _ := template.ParseFiles("templates/homePage.html") //templates(tmpl) = шаблон
	_ = tmpl.Execute(page, LorgarAurelian)
}

func ImperiaPage(page http.ResponseWriter /*Автор ответов*/, _ *http.Request /*Запрос*/) { //Страница Империя
	_, _ = fmt.Fprintf(page, "In the bleak darkness of the distant future there is only war")
}

func handleRequest() { //обрабатывать запрос
	http.HandleFunc("/", homePage)
	http.HandleFunc("/imperia/", ImperiaPage)
	_ = http.ListenAndServe(":4290", nil)
}

func main() { //Главная функция
	handleRequest()
}
