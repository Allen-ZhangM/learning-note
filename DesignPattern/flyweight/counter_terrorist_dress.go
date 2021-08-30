package main

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func (c *counterTerroristDress) setColor(color string) {
	c.color = color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}
