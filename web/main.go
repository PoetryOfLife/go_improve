package web

func main() {
	server := NewHttpServer("test")
	server.Route("POST", "/", nil)
	server.Start(":8080")

	defer func() {
		if data := recover(); data != nil {

		}
	}()
}
