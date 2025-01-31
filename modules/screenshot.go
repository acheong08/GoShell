package modules

import (
  "github.com/vova616/screenshot"
  "image/png"
  "os"
  "log"
)

func Screenshot(){
	img, err := screenshot.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("screenshot.png")
	if err != nil {
		log.Fatal("Failed to create file")
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal("Failed to write to file")
	}
	f.Close()
}
