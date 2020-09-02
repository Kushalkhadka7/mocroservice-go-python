package sample

import (
	laptop "manager/pb"
	"math/rand"
)

func getRandomKeyboardLayout() laptop.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return laptop.Keyboard_AZWETY

	case 2:
		return laptop.Keyboard_QWERTY

	case 3:
		return laptop.Keyboard_QWETRZ

	default:
		return laptop.Keyboard_QWERTY
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

func getRandomMemory() laptop.Memory_Unit {
	switch rand.Intn(3) {
	case 1:
		return laptop.Memory_GITABYTE

	case 2:
		return laptop.Memory_MEGABYTE

	case 3:
		return laptop.Memory_KILOBYTE

	default:
		return laptop.Memory_KILOBYTE
	}
}

func getStorageDriver() laptop.Storage_Driver {
	switch rand.Intn(2) {
	case 1:
		return laptop.Storage_HDD

	case 2:
		return laptop.Storage_SSD

	default:
		return laptop.Storage_SSD
	}
}

func getScreenPanel() laptop.Screen_Panel {
	switch rand.Intn(2) {
	case 1:
		return laptop.Screen_IPS

	case 2:
		return laptop.Screen_OLED

	default:
		return laptop.Screen_IPS
	}
}

func getRandomYear() uint32 {
	year := []uint32{2020, 2015, 2014, 2013, 2012}
	index := rand.Intn(4)

	return year[index]
}
