package main

import (
	"fmt"
	"net/http"

	_ "myproject/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {

	store := filestore.FileStore{
		Path: "/mnt",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	beego.Handler("/files/", http.StripPrefix("/files/", handler), true)
	beego.Run()
}
