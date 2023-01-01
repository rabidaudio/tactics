package game

type Custscene struct {
	BattleScene
	Actions []Command
}

func (c *Custscene) Tick() {
	c.BattleScene.Tick()
	// for each action, execute
}
