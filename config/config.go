package BaseConfig

var UrlDB = "URL_DB"
var TokenLetters = "TOKEN_LETTERS"
var TokenLenght = "TOKEN_LENGTH"

var Port = ":8080"
var BasePath = "/api"
var PathDocs = BasePath + "/docs"

var PathProfile = BasePath + "/profile"
var PathSignUp = BasePath + "/register"

var PathChats = BasePath + "/chats"
var PathChatByID = BasePath + "/chats/{id}"

var PathClasses = BasePath + "/classes"
var PathClassByID = BasePath + "/classes/{id}"
var PathClassStudents = BasePath + "/classes/{id}/students"

var Headers = map[string]string{
	"Access-Control-Allow-Origin":      "http://127.0.0.1:8001",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
