package model

type student struct {
	Name  string
	score float64
}

type Xiao struct {
	student
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
