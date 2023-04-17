package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/app"
)

func init() {
	flag.Parse()

	err := godotenv.Load()
	if os.Getenv("APP_ENV") == "local" && err != nil {
		log.Fatal("Error loading .env file")
	}
}

func catch(err error) {
	if err != nil {
		log.Fatalf("Something went wrong while preparing app, %+v", err)
	}
}

func main() {
	app := app.NewApp()
	catch(app.InitPostgres())
	catch(app.InitRouter())
	catch(app.InitServices())
	catch(app.Run())

	defer app.Stop()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
