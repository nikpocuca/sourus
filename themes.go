package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Tuple[T1 any, T2 any] struct {
	Left  T1
	Right T2
}

type ColorTheme struct {
	ThemeName                string
	HostMemoryGradientColors Tuple[string, string]
	CPULoadGradientColors    Tuple[string, string]
	GPUGradientColors        Tuple[string, string]
	BorderColor              string
	InstructionsColor        string
	HeaderTitleColor         string
	AppWidth                 int
	CPULoadGridColumns       int
}

func GenerateDarkTheme() ColorTheme {

	// host memory color gradient
	AegeanWaterBlue := "#00A5BF"
	CrystalRingMagenta := "#BF008F"

	// cpu load colors
	LimeCandyGreen := "#B0FF00"
	DangerFireRed := "#FF0F00"

	// gpu colors
	SunLemonYellow := "#FFFB00"
	LighterCrystalRingMagenta := "#FF0084"
	WhiteGrey := "#FFFFFF"
	ANSIBlue := "#0099FF"
	ANSILightBlue := "#00FFFF"

	return ColorTheme{
		ThemeName:                "Dark",
		HostMemoryGradientColors: Tuple[string, string]{AegeanWaterBlue, CrystalRingMagenta},
		CPULoadGradientColors:    Tuple[string, string]{LimeCandyGreen, DangerFireRed},
		GPUGradientColors:        Tuple[string, string]{SunLemonYellow, LighterCrystalRingMagenta},
		HeaderTitleColor:         WhiteGrey,
		BorderColor:              ANSIBlue,
		InstructionsColor:        ANSILightBlue,
		AppWidth:                 60,
		CPULoadGridColumns:       2,
	}
}

func GenerateLightTheme() ColorTheme {

	// host memory color gradient
	AegeanSkyBlue := "#87CEEB"
	BlushRosePink := "#FFC0CB"

	// cpu load colors
	MintGreen := "#98FF98"
	CoralRed := "#FF4040"

	// gpu colors
	SunflowerYellow := "#FFD700"
	SoftPeach := "#FFDAB9"
	LightGrey := "#D3D3D3"
	ANSILightCyan := "#E0FFFF"
	ANSIBrightBlue := "#4682B4"

	return ColorTheme{
		ThemeName:                "Light",
		HostMemoryGradientColors: Tuple[string, string]{AegeanSkyBlue, BlushRosePink},
		CPULoadGradientColors:    Tuple[string, string]{MintGreen, CoralRed},
		GPUGradientColors:        Tuple[string, string]{SunflowerYellow, SoftPeach},
		HeaderTitleColor:         LightGrey,
		BorderColor:              ANSIBrightBlue,
		InstructionsColor:        ANSILightCyan,
		AppWidth:                 60,
		CPULoadGridColumns:       2,
	}
}

const settingsDir string = ".sourus"

const settingsFile string = "settings.yml"

func GenerateTheme() ColorTheme {

	currentUser, _ := user.Current()

	settingsPath := filepath.Join("/home", currentUser.Username, settingsDir)

	err := os.MkdirAll(settingsPath, os.ModeDir)

	_, infoErr := os.Stat(settingsPath)

	// file already exists
	if infoErr == nil {

		// read in yml file where  into ColorTheme
		file, openErr := os.Open(filepath.Join(settingsPath, settingsFile))

		if openErr != nil {
			fmt.Println("Failed to open yml file", openErr)
			return GenerateDarkTheme()
		}

		var importedTheme ColorTheme

		decoder := yaml.NewDecoder(file)

		if decoderErr := decoder.Decode(&importedTheme); decoderErr != nil {
			fmt.Println("Failed to import color theme but file exists", decoderErr)
			return GenerateDarkTheme()
		}

		return importedTheme
	}

	// newly created file
	if err == nil {

		chmodErr := os.Chmod(settingsPath, 0777)

		if chmodErr != nil {
			fmt.Println("Failed to set permissions:", err)
			return GenerateDarkTheme()
		}

		file, creationErr := os.Create(filepath.Join(settingsPath, settingsFile))

		defer file.Close()

		if creationErr != nil {
			fmt.Println("Failed to create settings file, resorting to default theme", creationErr)
			return GenerateDarkTheme()
		}

		defaultDarkTheme := GenerateDarkTheme()

		yamlEncoder := yaml.NewEncoder(file)
		defer yamlEncoder.Close()

		// attempt encoding
		if encodingErr := yamlEncoder.Encode(defaultDarkTheme); encodingErr != nil {
			fmt.Println("Failed to encode ColorTheme resorting to default")
			return defaultDarkTheme
		}

		return defaultDarkTheme

	}

	fmt.Println("Failed to create settings directory", err)
	return GenerateDarkTheme()

}
