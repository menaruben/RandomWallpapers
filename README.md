# RandomWallpapers
## Usage
After building the executable using ```go build randwallpaper.go``` and adding it to ```PATH``` you can run it like a CLI application. 
There are several parameters:
- path: 
    - type:         string
    - default:      current working directory
    - description:  specifies the path where all wallpaper images are located
- interval:
    - type:         int
    - default:      5
    - description:  specifies the amount of seconds between wallpaper replacements
- crop:
    - type:         bool
    - default:      true
    - description:  specifies if wallpapers should be cropped or not

### Example using all parameters:
The following command sets the random images inside the path as the wallpaper every 10 seconds in cropped mode. 
```
go run .\randwallpaper.go --path C:\repos\RandomWallpapers\images --interval 10 --crop
```
