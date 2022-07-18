// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Сервис принимает xml документ
type XmlAnalitics interface {
	SendXML()
}

// Структура xml документа
type XmlDoc struct {
}

// Реализация интерфейса отправки xml документа
func (doc XmlDoc) SendXML() {
	fmt.Println("Отправка XML документа")
}

// Структура json
type JsonDocument struct {
}

// Преобразование json в xml
func (JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

// Адаптер который будет реализовать интерфейс отправки xml
type AdapterJson struct {
	JsonDocument *JsonDocument
}

// Реализация интерфеса отправки xml документа
func (adapter AdapterJson) SendXML() {
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Отправка XML документа после конвертации из json")
}

func main() {
	xmlFile := &XmlDoc{}
	xmlFile.SendXML()
	jsonFile := &JsonDocument{}
	AdapterJson{jsonFile}.SendXML()
}
