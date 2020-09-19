package sample

import (
	laptop "client/pb"

	"github.com/golang/protobuf/ptypes"
)

// NewKeyboard initializes new random key board.
func NewKeyboard() *laptop.Keyboard {
	return &laptop.Keyboard{
		Layout:  getRandomKeyboardLayout(),
		Backlit: getBoolean(),
	}

}

// NewCPU initializes sample cpu.
func NewCPU() *laptop.CPU {
	return &laptop.CPU{
		Brand:       getRandomCPUBrand(),
		Name:        getRandomString(),
		NoOfCores:   uint32(getRandomInt()),
		NoOfThreads: uint32(getRandomInt()),
		MinGhz:      getRandomFLoat(),
		MaxGhz:      getRandomFLoat(),
	}
}

// NewGPU initializes new sample gpu.
func NewGPU() *laptop.GPU {
	return &laptop.GPU{
		Brand:  getGPUBrand(),
		Name:   getRandomString(),
		MinGhz: getRandomFLoat(),
		MaxGhz: getRandomFLoat(),
		Memory: &laptop.Memory{
			Unit:  getRandomMemory(),
			Value: uint32(getRandomFLoat()),
		},
	}
}

// NewStorage initializes sample laptop storage
func NewStorage() *laptop.Storage {
	return &laptop.Storage{
		Driver: getStorageDriver(),
		Memory: &laptop.Memory{
			Unit:  getRandomMemory(),
			Value: uint32(getRandomFLoat()),
		},
	}
}

// NewScreen initializes new sample laptop screen
func NewScreen() *laptop.Screen {
	return &laptop.Screen{
		Size: float32(getRandomInt()),
		Resolution: &laptop.Screen_Resolution{
			Width:  200,
			Height: 300,
		},
		Panel:      getScreenPanel(),
		Multitouch: getBoolean(),
	}

}

// NewLaptop generates sampled data for laptop.
func NewLaptop() *laptop.Laptop {
	return &laptop.Laptop{
		XId:   "",
		Brand: getRandomCPUBrand(),
		Name:  getRandomString(),
		Cpu:   NewCPU(),
		Ram: &laptop.Memory{
			Unit:  getRandomMemory(),
			Value: uint32(getRandomFLoat()),
		},
		Gpus: []*laptop.GPU{
			NewGPU(),
			NewGPU(),
		},
		Storage: []*laptop.Storage{
			NewStorage(),
			NewStorage(),
		},
		Keyboard:    NewKeyboard(),
		Weight:      getRandomFLoat(),
		PriceUsd:    getRandomFLoat(),
		ReleaseYear: getRandomYear(),
		UpdatedAt:   ptypes.TimestampNow(),
	}

}
