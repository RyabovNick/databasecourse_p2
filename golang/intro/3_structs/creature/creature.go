package creature

type Creature struct {
	x float64
	y float64
}

// New creates new creature
func New() *Creature {
	return &Creature{
		x: 10,
		y: 5,
	}
}

func (c *Creature) Eat() {
	c.x += 10
}
