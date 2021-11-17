package Bridge_Pattern

import "testing"

func TestBridgePattern(t *testing.T) {
	mac := &Mac{
		OperatingSystem{
			&RmvbDecoder{},
		},
	}
	mac.play("Titanic")

	win := &Windows{
		 OperatingSystem{
			&AviDecoder{},
		 },
	}
	win.play("Titanic")
}
