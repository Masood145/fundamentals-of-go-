package learning

/*
type Player struct {
	posx float64
	posy float64
}

func (p *Player) move(x, y float64) {
	p.posx += x
	p.posy += y

}
func (p *Player) teloport(x, y float64) {
	p.posx = x
	p.posy = y

}

type Enemy struct {
	posx float64
	posy float64
}

func (e *Enemy) move(x, y float64) {
	e.posx += x
	e.posy += y

}
func (e *Enemy) teloport(x, y float64) {
	e.posx = x
	e.posy = y

}
*/

//Instead of all the above we can use struct embeding more like teh compostion using the same stuct and ha acces to all the function

type Postion struct {
	x float64
	y float64
}
type SpecialPosition struct {
	Postion
}

func (p *Postion) Move(x, y float64) {
	p.x += x
	p.y += y
}
func (p *Postion) Teleport(x, y float64) {
	p.x = x
	p.y = y
}
func (sp *SpecialPosition) MoveSpecial(x, y float64) {

	sp.x += x * x
	sp.y += y * y
}

type Player struct {
	*SpecialPosition
}

func Newplayer() *Player {
	//using this we have acces to the move and teleport func
	/*return &Player{
		Postion : &Postion{}
	}*/
	// we can access to the all the function because the Special position struct have access to postion struct
	return &Player{
		SpecialPosition: &SpecialPosition{},
	}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}

}
