package main

type Game struct {
	playerName  string
	playerScore int
}

func CreateGame(name string, score int) *Game {
	return &Game{
		playerName:  name,
		playerScore: score,
	}
}

func (g *Game) SetScore(score int) {
	g.playerScore = score
}

func (g *Game) GetScore() int {
	return g.playerScore
}

func (g *Game) RevertState(point *checkPoint) {
	g.playerName = point.playerName
	g.playerScore = point.playerScore
}

func (g *Game) saveState() *checkPoint {
	return createCheckPoint(g.playerName, g.playerScore)
}

type checkPoint struct {
	playerName  string
	playerScore int
}

func createCheckPoint(name string, score int) *checkPoint {
	return &checkPoint{
		playerName:  name,
		playerScore: score,
	}
}

type Caretaker struct {
	saves [10]*checkPoint
}

func (c *Caretaker) Save(g *Game, idx uint8) {
	c.saves[idx] = g.saveState()
}

func (c *Caretaker) Revert(g *Game, idx uint8) {
	g.RevertState(c.saves[idx])
}
