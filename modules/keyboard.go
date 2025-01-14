package modules

import (
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

func Keyboard() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_C)

	// Set shift to be pressed
	kb.HasSuper(true)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
  // Or you can use Press and Release
	kb.Press()
	time.Sleep(100 * time.Millisecond)
	kb.Release()
}
