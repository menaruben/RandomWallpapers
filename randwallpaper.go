package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/reujab/wallpaper"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getRandomImage(path string) (string, error) {
	images, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	// Initialize the random number generator with a different seed value each "time"
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(images)-1)
	randomImagePath := path + "/" + images[randomIndex].Name()

	return randomImagePath, err
}

func setWallpaperLoop(path string, interval time.Duration, crop bool) {
	for {
		randomImagePath, err := getRandomImage(path)
		check(err)

		err = wallpaper.SetFromFile(randomImagePath)
		check(err)

		if crop {
			err = wallpaper.SetMode(wallpaper.Crop)
			check(err)
		}
		time.Sleep(interval)
	}
}

func main() {
	path := flag.String("path", "C:/repos/RandomWallpapers/images", "Path to the image directory")
	interval := flag.Int("interval", 5, "Interval in seconds")
	crop := flag.Bool("crop", true, "set mode of wallpapers to cropped")
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Options:")
		flag.PrintDefaults()
	}
	flag.Parse()
	intervalDuration := time.Duration(*interval) * time.Second
	setWallpaperLoop(*path, intervalDuration, *crop)
}
