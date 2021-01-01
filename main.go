package main

import (
	"fmt"
	"net/http"

	_ "myproject/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tus/tusd/cmd/tusd/cli/hooks"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {

	hookHandler := &hooks.FileHook{
		Directory: "/home/kajol/go/src/myproject/hooks",
	}

	composer := tusd.NewStoreComposer()
	config := tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	}
	//hookHandler.InvokeHook(hooks.HookPreCreate, tusd.Handler., true)

	store := filestore.FileStore{
		Path: "/opt/SLB-uploads",
	}
	store.UseIn(composer)

	handler, err := tusd.NewHandler(config)
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s (%d bytes) finished\n", event.Upload.ID, event.Upload.Size)
			hookHandler.InvokeHook(hooks.HookPostFinish, event, false)
		}
	}()

	beego.Handler("/files/", http.StripPrefix("/files/", handler), true)
	beego.Handler("/metrics", promhttp.Handler(), true)
	beego.Run()
}
