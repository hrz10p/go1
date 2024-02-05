package main

func main() {
	app := NewApplication()
	err := app.Start(":8080")
	if err != nil {
		app.Logger.Error(err.Error())
	}
}
