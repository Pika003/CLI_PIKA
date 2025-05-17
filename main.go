package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//go:embed assets/*.mp3
var audioFiles embed.FS

func main() {
	args := os.Args
	if len(args) < 2 {
		showWelcome()  // Show welcome screen instead of just help
		return
	}

	command := strings.ToLower(args[1])
	// Remove the hyphen if present
	command = strings.TrimPrefix(command, "-")

	switch command {
	case "welcome", "w":
		showWelcome()
		playSound("pikamain.mp3")
	case "pika", "p":
		fmt.Println("PIKA PIKA!")
		playSound("pika.mp3")
	case "dance", "d":
		PikaDance()
	case "joke", "j":
		pikaJoke()
		playSound("pikamain.mp3") 
	case "help", "h":
		printHelp()
		playSound("pikamain.mp3") 
	case "info", "i":
		if len(args) < 3 {
			fmt.Println("Please provide a Pokémon name. Example: pika -info pikachu")
			return
		}
		getPokemonInfo(strings.ToLower(args[2]))
		playSound("pikamain.mp3") 
	case "version", "v":
		fmt.Println("Pika CLI v1.0.0")
		playSound("pikamain.mp3")
	default:
		fmt.Println("Unknown command. Try 'pika -help' for available commands.")
	}
}

func showWelcome() {
	welcome := `
    ⚡️ Welcome to Pika CLI! ⚡️
    
    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
    ⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠾⠛⢉⣉⣉⣉⡉⠛⠷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀
    ⠀⠀⠀⠀⠀⠀⢀⣴⠋⣠⣴⣿⣿⣿⣿⣿⣿⣿⣦⣌⠙⣷⣄⠀⠀⠀⠀⠀⠀
    ⠀⠀⠀⠀⠀⡴⠋⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣌⠙⢷⡀⠀⠀⠀⠀
    ⠀⠀⠀⢀⡼⢁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⢹⣆⠀⠀⠀
    ⠀⠀⢀⡞⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⢻⣆⠀⠀
    ⠀⠀⡞⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⢻⡆⠀
    ⠀⢰⢃⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡘⣧⠀
    ⠀⢸⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⣿⠀
    ⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
    ⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀
    ⠀⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀
    ⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀
    ⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀
    ⠀⠀⠀⠀⠈⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠀⠀⠀⠀
    ⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⠿⢿⣿⣿⣿⣿⣿⡿⠿⠟⠛⠉⠀⠀⠀⠀⠀⠀⠀

    Type 'pika -help' to see available commands!
    `
	fmt.Println(welcome)
}

func printHelp() {
	fmt.Println("Pika CLI - A fun Pikachu-themed command line tool")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  -pika, -p    - Displays PIKA PIKA and plays a sound")
	fmt.Println("  -dance, -d   - Shows a dancing Pikachu ASCII animation")
	fmt.Println("  -joke, -j    - Tells a random Pikachu joke")
	fmt.Println("  -info, -i    - Displays information about a Pokémon (usage: pika -info [pokemon_name])")
	fmt.Println("  -help, -h    - Shows this help message")
	fmt.Println("  -version, -v - Shows the current version")
}

func pikaJoke() {
	joke := GetRandomJoke()
	fmt.Println(joke)
	// Removed go keyword from here as sound will be played in main
}

func getPokemonInfo(pokemonName string) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	
	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 404 {
		fmt.Printf("Pokémon '%s' not found! Check the spelling and try again.\n", pokemonName)
		return
	}
	
	if resp.StatusCode != 200 {
		fmt.Printf("Error: API returned status code %d\n", resp.StatusCode)
		return
	}
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}
	
	// Parse JSON
	var pokemonData map[string]interface{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	
	// Display Pokémon information in a cool way
	displayPokemonInfo(pokemonData)
	// Removed go keyword from here as sound will be played in main
}

func displayPokemonInfo(data map[string]interface{}) {
	// Extract basic information
	name := strings.Title(data["name"].(string))
	id := int(data["id"].(float64))
	height := int(data["height"].(float64)) / 10.0 // Convert to meters
	weight := int(data["weight"].(float64)) / 10.0 // Convert to kg
	baseExp := int(data["base_experience"].(float64))
	
	// Extract types
	types := []string{}
	for _, typeData := range data["types"].([]interface{}) {
		typeInfo := typeData.(map[string]interface{})
		typeDetails := typeInfo["type"].(map[string]interface{})
		types = append(types, strings.Title(typeDetails["name"].(string)))
	}
	
	// Extract abilities
	abilities := []string{}
	for _, abilityData := range data["abilities"].([]interface{}) {
		abilityInfo := abilityData.(map[string]interface{})
		abilityDetails := abilityInfo["ability"].(map[string]interface{})
		abilityName := strings.Title(abilityDetails["name"].(string))
		isHidden := abilityInfo["is_hidden"].(bool)
		
		if isHidden {
			abilities = append(abilities, abilityName+" (Hidden)")
		} else {
			abilities = append(abilities, abilityName)
		}
	}
	
	// Extract stats
	stats := map[string]int{}
	for _, statData := range data["stats"].([]interface{}) {
		statInfo := statData.(map[string]interface{})
		baseStat := int(statInfo["base_stat"].(float64))
		statDetails := statInfo["stat"].(map[string]interface{})
		statName := statDetails["name"].(string)
		
		// Convert stat names to more readable format
		switch statName {
		case "hp":
			stats["HP"] = baseStat
		case "attack":
			stats["Attack"] = baseStat
		case "defense":
			stats["Defense"] = baseStat
		case "special-attack":
			stats["Sp. Attack"] = baseStat
		case "special-defense":
			stats["Sp. Defense"] = baseStat
		case "speed":
			stats["Speed"] = baseStat
		}
	}
	
	// Display the information
	fmt.Println("\n╔" + strings.Repeat("═", 50) + "╗")
	fmt.Printf("║ %-48s║\n", "🔍 POKÉMON INFORMATION")
	fmt.Println("╠" + strings.Repeat("═", 50) + "╣")
	fmt.Printf("║ %-15s %-32s ║\n", "Name:", name)
	fmt.Printf("║ %-15s #%-31d ║\n", "Pokédex ID:", id)
	fmt.Printf("║ %-15s %-32s ║\n", "Type:", strings.Join(types, ", "))
	fmt.Printf("║ %-15s %.1fm                             ║\n", "Height:", float64(height))
	fmt.Printf("║ %-15s %.1fkg                            ║\n", "Weight:", float64(weight))
	fmt.Printf("║ %-15s %-32d ║\n", "Base Exp:", baseExp)
	fmt.Println("╠" + strings.Repeat("═", 50) + "╣")
	fmt.Printf("║ %-48s║\n", "⚡ ABILITIES")
	for _, ability := range abilities {
		fmt.Printf("║ %-48s ║\n", "• "+ability)
	}
	fmt.Println("╠" + strings.Repeat("═", 50) + "╣")
	fmt.Printf("║ %-48s║\n", "📊 STATS")
	for stat, value := range stats {
		fmt.Printf("║ %-15s %-32d ║\n", stat+":", value)
	}
	fmt.Println("╚" + strings.Repeat("═", 50) + "╝")
}

func playSound(filename string) {
    data, err := audioFiles.ReadFile("assets/" + filename)
    if err != nil {
        fmt.Printf("Error reading embedded audio file: %v\n", err)
        return
    }

    tmpFile, err := os.CreateTemp("", "*.mp3")
    if err != nil {
        fmt.Printf("Error creating temp file: %v\n", err)
        return
    }
    defer os.Remove(tmpFile.Name())

    if _, err := tmpFile.Write(data); err != nil {
        fmt.Printf("Error writing to temp file: %v\n", err)
        return
    }
    tmpFile.Close()

    var cmd *exec.Cmd
    if runtime.GOOS == "darwin" {
        cmd = exec.Command("afplay", tmpFile.Name())
    } else if runtime.GOOS == "linux" {
        cmd = exec.Command("aplay", tmpFile.Name())
    } else if runtime.GOOS == "windows" {
        cmd = exec.Command("powershell", "-c",
            `Add-Type -AssemblyName PresentationCore; 
            $player = New-Object System.Windows.Media.MediaPlayer;
            $player.Open([System.Uri]"` + tmpFile.Name() + `");
            $player.Play(); Start-Sleep -s 3; $player.Stop(); $player.Close()`)
    }

    // fmt.Println("Playing embedded sound...")
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error playing sound: %v\n", err)
    }
}