package main

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func (t *terroristDress) setColor(color string) {
	t.color = color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}
