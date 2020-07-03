package main

import ( 
  "fmt"
  "math/rand"
  "time"
)

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You have no fuel at all. None. You are going to crash. Bye!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus":
    fuel = 300000
    case "Mercury": 
    fuel = 500000
    case "Mars":
    fuel = 700000
    case "Jupiter":
    fuel = 1000000
    case "Saturn":
    fuel = 2000000
    default :
    fuel = 0
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("You're destination is", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have enough fuel for your greedy selves who are inconsiderate.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining int
  var fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else if fuelCost > fuelRemaining {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  rand.Seed(time.Now().UnixNano())

  var fuel int
  var planetChoice string
  fuel = 1000000
  randomPlanet := rand.Intn(4)
  switch randomPlanet {
    case 0:
    planetChoice = "Saturn"
    case 1:
    planetChoice = "Venus"
    case 2:
    planetChoice = "Mercury"
    case 3:
    planetChoice = "Mars"
    case 4:
    planetChoice = "Jupiter"
  }
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  // And then liftoff!
  
}