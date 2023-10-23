package BaseConfig

var Port = ":8080"
var BasePath = "/api"
var PathDocs = BasePath + "/docs"

var PathSignUp = BasePath + "/register"
var PathProfile = BasePath + "/profile"
var PathAddStudent = BasePath + "/addstudent"
var PathSend = BasePath + "/send"
var PathChats = BasePath + "/chats/{teacherID}"

var Headers = map[string]string{
	//"Access-Control-Allow-Origin":      "http://89.208.198.137:8081",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
