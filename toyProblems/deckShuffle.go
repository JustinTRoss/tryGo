package main

import ("fmt", "crypto/rand")

func shuffleDeck( deck [52]string ) [52]string{
  dLen := deck.length

  for i, value := range deck {
    randI := rand.Int(rand io.Reader, dLen)
    if i !== randI {
      temp := deck[i]
      deck[i] = deck[randI]
      deck[randI] = temp
    }
  }

  return deck
}

func main() {

}