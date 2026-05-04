package kata
package main
	
func RPS(P1, P2 string) string {
	if P1 == P2 {
		return "Draw!"
	} else if P1 == "rock" && 
		P2 == "scissors" || 
		P1 == "scissors" &&  
		P2 == "paper" ||
		P1 == "paper" &&
		P2 == "rock" {
			return "Jugador 1 gana!"
		} else {
			return "Jugador 2 gana!"
		}
}

func main(){
	RPS("rock", "scissors")
	RPS("paper", "paper")
	RPS("scissors", "rock")
	RPS("rock", "rock")
	RPS("scissors", "scissors")
	RPS("paper", "scissors")
	RPS("scissors", "paper")
	RPS("paper", "rock")
	RPS("rock", "paper")
	RPS("scissors", "rock")
	RPS("paper", "scissors")
	RPS("rock", "scissors")
}
