package main

import (
	"fmt"
	"net/http"

	_ "myproject/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tus/tusd/cmd/tusd/cli"
	"github.com/tus/tusd/cmd/tusd/cli/hooks"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {

	store := filestore.FileStore{
		Path: "/home/SLB/uploads",
	}

	composer := tusd.NewStoreComposer()
	config := tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
		PreUploadCreateCallback: func(hook tusd.HookEvent) error {
			fmt.Println("pre-create event called")
			return nil
		},
		PreFinishResponseCallback: func(hook tusd.HookEvent) error {
			fmt.Println("pre-finish event called")
			return nil
		},
	}
	//hookHandler.InvokeHook(hooks.HookPreCreate, tusd.Handler., true)

	hookHandler := &hooks.FileHook{
		Directory: "/home/SLB/src/myproject/hooks",
	}
	store.UseIn(composer)

	handler, err := tusd.NewHandler(config)
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	cli.Flags.MetricsPath = "/metrics"
	cli.SetupMetrics(handler)
	//prometheus.MustRegister(prometheuscollector.New(handler.Metrics))
	cli.SetupHookMetrics()
	uploadedFileSizes := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "uploaded_file_sizes",
		Help:    "Uploaded file size in bytes",
		Buckets: prometheus.LinearBuckets(1000000, 5000000, 10),
	})

	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s (%d bytes) finished\n", event.Upload.ID, event.Upload.Size)
			hookHandler.InvokeHook(hooks.HookPostFinish, event, false)
			uploadedFileSizes.Observe(float64(event.Upload.Size))
		}
	}()
	
	prometheus.MustRegister(uploadedFileSizes)

	beego.Handler("/files/", http.StripPrefix("/files/", handler), true)
	beego.Handler("/metrics", promhttp.Handler(), true)
	beego.Run()
}
