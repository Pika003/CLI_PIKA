package main

import (
	"math/rand"
	"time"
)

// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomJoke returns a random Pikachu joke from the collection
func GetRandomJoke() string {
	jokes := []string{
		"Why did Pikachu go to the gym? To get a little more SPARK-ly!",
		"What's Pikachu's favorite dessert? SHOCK-olate cake!",
		"How does Pikachu greet other Pokémon? 'WATTS up!'",
		"Why don't you want to get into an argument with Pikachu? Because it's always CHARGED up!",
		"What did Pikachu say when it saw something shocking? 'That's re-VOLT-ing!'",
		"What do you call a Pikachu with a cold? An ELECTRIC-ACHOO!",
		"Why was Pikachu so good at baseball? It never strikes out!",
		"How does Pikachu start a car? With a JUMP-START!",
		"What's Pikachu's favorite TV show? POWER Rangers!",
		"Why did Pikachu cross the road? To get to the other VOLT!",
		"What's Pikachu's favorite band? AC/DC!",
		"What did Pikachu say after winning a battle? 'That was ELECTRIFYING!'",
		"Where does Pikachu do its shopping? At VOLT-mart!",
		"What's Pikachu's favorite dance? The ELECTRIC slide!",
		"Why did Pikachu become a DJ? Because it loves dropping the VOLT-age!",
		"What's Pikachu's favorite exercise? CIRCUIT training!",
		"Why was Pikachu such a good student? It had a lot of POTENTIAL!",
		"What's Pikachu's favorite movie? The CURRENT war!",
		"How does Pikachu keep fit? By doing OHM-ing exercises!",
		"What's Pikachu's favorite food? THUNDER-dogs with SHOCK-raut!",
	}

	return jokes[rand.Intn(len(jokes))]
}