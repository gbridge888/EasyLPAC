package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed assets/DroidSansFallback.ttf
var droidSansFallback []byte

//go:embed assets/DroidSansMono.ttf
var droidSansMono []byte

//go:embed assets/DroidSansBold.ttf
var droidSansBold []byte

//go:embed assets/NotoSansJP.ttf
var notoSansJP []byte

//go:embed assets/NotoSansJP-Bold.ttf
var notoSansJPBold []byte

var resourceDroidSansFallback = &fyne.StaticResource{
	StaticName:    "DroidSansFallback.ttf",
	StaticContent: droidSansFallback,
}

var resourceDroidSansMono = &fyne.StaticResource{
	StaticName:    "DroidSansMono.ttf",
	StaticContent: droidSansMono,
}

var resourceDroidSansBold = &fyne.StaticResource{
	StaticName:    "DroidSansBold.ttf",
	StaticContent: droidSansBold,
}

var resourceNotoSansJP = &fyne.StaticResource{
	StaticName:    "NotoSansJP.ttf",
	StaticContent: notoSansJP,
}

var resourceNotoSansJPBold = &fyne.StaticResource{
	StaticName:    "NotoSansJP-Bold.ttf",
	StaticContent: notoSansJPBold,
}
