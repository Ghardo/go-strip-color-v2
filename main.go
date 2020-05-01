package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Vars injected via ldflags by bundler
var (
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	debug    = false
	w        *astilectron.Window
	appData  map[string]interface{}
	dataPath = path.Join(os.Getenv("APPDATA"), "ArduinoStripColorChanger")
	dataFile = "data.json"
)

func main() {
	// Parse flags
	flag.Parse()

	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	appData = map[string]interface{}{}
	loadData()

	url := "index.html"

	// Debug
	if debug {
		url = "http://localhost:8080"
	}

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            "ArduinoStripColorChanger",
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: "0.37.0",
			VersionElectron:    "8.2.0",
		},
		Debug:       debug,
		Logger:      l,
		MenuOptions: nil,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(2 * time.Second)
				if debug {
					w.OpenDevTools()
				}
			}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       url,
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333333"),
				Center:          astikit.BoolPtr(true),
				Resizable:       astikit.BoolPtr(true),
				Height:          astikit.IntPtr(755),
				Width:           astikit.IntPtr(930),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}

	saveData()
}
