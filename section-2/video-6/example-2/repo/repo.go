package repo

//MOCK
func GetSessionData(session string) map[string]string {
	sessionData := make(map[string]string)
	sessionData["Username"] = "JaneDoe"
	sessionData["Name"] = "Jane"
	sessionData["LastName"] = "Doe"
	return sessionData
}
