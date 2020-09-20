package sample

import (
	pb "client/pb"
	"math/rand"
)

func getRandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_AZWETY

	case 2:
		return pb.Keyboard_QWERTY

	case 3:
		return pb.Keyboard_QWETRZ

	default:
		return pb.Keyboard_QWERTY
	}
}

func getBoolean() bool {
	return rand.Intn(2) == 1
}

func getRandomCPUBrand() string {
	brand := []string{"HP", "LENOVO", "DELL", "APPLE", "ACER"}

	index := rand.Intn(4)

	return brand[index]
}

func getRandomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, 10)
	for i := range randomString {
		randomString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(randomString)
}

func getRandomInt() int {
	return rand.Intn(8)
}

func getRandomFLoat() float64 {
	max := 5.0
	min := 2.0

	return min + rand.Float64()*(max-min)
}

func getGPUBrand() string {
	brand := []string{"NVIDIA", "RADEON", "AMD"}

	index := rand.Intn(2)

	return brand[index]
}

func getRandomMemory() pb.Memory_Unit {
	switch rand.Intn(3) {
	case 1:
		return pb.Memory_GITABYTE

	case 2:
		return pb.Memory_MEGABYTE

	case 3:
		return pb.Memory_KILOBYTE

	default:
		return pb.Memory_KILOBYTE
	}
}

func getStorageDriver() pb.Storage_Driver {
	switch rand.Intn(2) {
	case 1:
		return pb.Storage_HDD

	case 2:
		return pb.Storage_SSD

	default:
		return pb.Storage_SSD
	}
}

func getScreenPanel() pb.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return pb.Screen_IPS

	case 2:
		return pb.Screen_OLED

	default:
		return pb.Screen_IPS
	}
}

func getRandomYear() uint32 {
	year := []uint32{2020, 2015, 2014, 2013, 2012}
	index := rand.Intn(4)

	return year[index]
}
