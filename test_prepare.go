package tournament

import (
	"fmt"
	"math/rand"
	"time"
)

var teamNames = []string{
	"Awesome Antelopes", "Agile Alligators", "Artistic Armadillos", "Amazing Aardvarks", "Astonishing Alpacas",
	"Brilliant Baboons", "Bouncy Bears", "Brawny Beavers", "Bold Bison", "Bright Badgers",
	"Crafty Coyotes", "Cunning Cheetahs", "Creative Cranes", "Clever Cats", "Curious Chameleons",
	"Daring Dragons", "Dynamic Dolphins", "Dazzling Deer", "Determined Dogs", "Delightful Ducks",
	"Energetic Eagles", "Elegant Elephants", "Enthusiastic Echidnas", "Excellent Emus", "Excited Elk",
	"Friendly Foxes", "Fast Falcons", "Fierce Ferrets", "Funny Frogs", "Funky Flamingos",
	"Gallant Gorillas", "Graceful Gazelles", "Gentle Giraffes", "Giant Geckos", "Generous Gophers",
	"Happy Hippos", "Hardy Hedgehogs", "Heroic Hares", "Hilarious Hamsters", "Hopeful Hyenas",
	"Incredible Iguanas", "Intelligent Impalas", "Inquisitive Insects", "Inspiring Ibexes", "Imaginative Ibis",
	"Jolly Jaguars", "Jubilant Jackrabbits", "Joyful Jackals", "Jaunty Jellyfish", "Jovial Jerboas",
	"Keen Kangaroos", "Kind Koalas", "Kooky Kittens", "Kingly Kingfishers", "Knowledgeable Kestrels",
	"Lively Lemurs", "Lucky Lions", "Lovely Llamas", "Lightning Lizards", "Lazy Lynxes",
	"Majestic Moose", "Mischievous Monkeys", "Mighty Mice", "Mystical Magpies", "Merry Meerkats",
	"Nimble Newts", "Nifty Narwhals", "Nice Nudibranchs", "Noble Numbats", "Naughty Nuthatches",
	"Optimistic Otters", "Outgoing Opossums", "Observant Ocelots", "Obedient Octopuses", "Overjoyed Ospreys",
	"Precious Penguins", "Powerful Panthers", "Plucky Puffins", "Proud Peacocks", "Playful Platypuses",
	"Quick Quokkas", "Quiet Quails", "Quizzical Quetzals", "Quirky Quaggas", "Qualified Quokkas",
	"Radiant Raccoons", "Resourceful Rhinos", "Reliable Ravens", "Romantic Reindeer", "Restless Rabbits",
	"Silly Sloths", "Sneaky Snakes", "Spectacular Sparrows", "Speedy Squirrels", "Sassy Stingrays",
	"Trusty Turtles", "Talented Tigers", "Thundering Turkeys", "Tough Tarsiers", "Talkative Toucans",
	"Unique Unicorns", "Understanding Uakaris", "Uptight Uguisu", "Upbeat Umbrellabirds", "Unwavering Urutus",
	"Vibrant Voles", "Vengeful Vultures", "Vigorous Vicuñas", "Vivacious Vaquitas", "Valiant Vampire Bats",
	"Whimsical Wombats", "Wild Wolves", "Wise Warthogs", "Wandering Weasels", "Wavy Whales",
	"X-treme Xerus", "Xenophobic Xenarthrans", "Xenial Xoloitzcuintles", "Xenodochial Xerus", "Xenophilic Xenopuses",
	"Young Yellow Jackets", "Yawning Yaks", "Yearning Yellowjackets", "Yellow-bellied Yabbies", "Yodeling Yetis",
	"Zany Zebras", "Zealous Zebus", "Zesty Zorillas", "Zigzagging Zebrasses", "Zombie-like Zanzibar Gems",
}

// GenerateTournamentResults generates tournament results for the specified number of teams
func GenerateTournamentResults(numTeams int) []string {
	teams := make([]string, 0, numTeams)
	for _, index := range GenerateNonRepeatingRandomNumbers(numTeams, len(teamNames)-1) {
		teams = append(teams, teamNames[index])
	}

	// Generate match results
	matches := []string{}
	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			result := ""
			switch rand.Intn(3) {
			case 0:
				result = "win"
			case 1:
				result = "draw"
			case 2:
				result = "loss"
			}
			matches = append(matches, fmt.Sprintf("%s;%s;%s", teams[i], teams[j], result))
		}
	}

	return matches
}

func GenerateNonRepeatingRandomNumbers(n int, max int) []int {
	if n > max {
		n = max
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numsMap := make(map[int]struct{})

	for i := 0; i < n; i++ {
		num := r.Intn(max) + 1
		for contains(numsMap, num) {
			num = rand.Intn(max) + 1
		}

		numsMap[num] = struct{}{}
	}

	nums := make([]int, 0, n)
	for i := range numsMap {
		nums = append(nums, i)
	}

	return nums
}

// contains checks if a slice contains a given value
func contains(nums map[int]struct{}, num int) bool {
	_, ok := nums[num]
	return ok
}
