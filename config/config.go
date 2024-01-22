package BaseConfig

var PG_USER = "POSTGRES_USER"
var PG_HOST = "POSTGRES_HOST"
var PG_PWD = "POSTGRES_PASSWORD"
var PG_PORT = "POSTGRES_PORT"
var PG_DB = "POSTGRES_DB"

var TOKEN_LETTERS = "TOKEN_LETTERS"
var TOKEN_LENGTH = "TOKEN_LENGTH"
var CHAT_GRPC_URL = "CHAT_GRPC_URL"
var CALENDAR_GRPC_URL = "CALENDAR_GRPC_URL"
var FILESTORAGE_PATH = "FILESTORAGE_PATH"
var CHAT_FILES_PATH = "CHAT_FILES_PATH"
var HOMEWORK_FILES_PATH = "HOMEWORK_FILES_PATH"
var SOLUTION_FILES_PATH = "SOLUTION_FILES_PATH"
var URL_DOMAIN = "URL_DOMAIN"

var Port = ":8080"
var BasePath = "/api"
var PathDocs = BasePath + "/docs"

var PathAttach = BasePath + "/attach"

var PathProfile = BasePath + "/profile"
var PathSignUp = BasePath + "/signup"
var PathLogin = BasePath + "/login"
var PathLogout = BasePath + "/logout"
var PathAuth = BasePath + "/auth"

var PathChats = BasePath + "/chats"
var PathChatByID = BasePath + "/chats/{id}"

var PathCalendar = BasePath + "/calendar"
var PathAddEvent = BasePath + "/calendar/addevent"
var PathGetEvents = BasePath + "/calendar/events"
var PathEvent = BasePath + "/calendar/event"

var PathClasses = BasePath + "/classes"
var PathClassByID = BasePath + "/classes/{id}"
var PathClassStudents = BasePath + "/classes/{id}/students"
var PathClassFeed = BasePath + "/classes/{id}/feed"
var PathClassHomeworks = BasePath + "/classes/{id}/homeworks"
var PathClassSolutions = BasePath + "/classes/{id}/solutions"

var PathHomework = BasePath + "/homeworks"
var PathHomeworkByID = BasePath + "/homeworks/{id}"
var PathHomeworkSolutions = BasePath + "/homeworks/{id}/solutions"

var PathTasks = BasePath + "/tasks"
var PathTaskByID = BasePath + "/tasks/{id}"

var PathSolution = BasePath + "/solutions/{id}"

var PathStudent = BasePath + "/students/{id}"

var Headers = map[string]string{
	"Access-Control-Allow-Origin":      "http://127.0.0.1:8001",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
