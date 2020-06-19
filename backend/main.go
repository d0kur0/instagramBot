package main

import "github.com/d0kur0/instagramBot/utilites/consoleWriter"

func main() {
	consoleWriter.Header("Начинаю авторизацию")
	consoleWriter.Info("Жду пока отрисуется элемент .root")
	consoleWriter.Success("Элемент отрисован")
	consoleWriter.Info("Жду пока отрисуется элемент .root-2")
	consoleWriter.Error("Элемент не был отрисован")

	consoleWriter.Header("Начинаю авторизацию")
	consoleWriter.Info("Жду пока отрисуется элемент .root")
	consoleWriter.Success("Элемент отрисован")
	consoleWriter.Info("Жду пока отрисуется элемент .root-2")
	consoleWriter.Error("Элемент не был отрисован")
}
