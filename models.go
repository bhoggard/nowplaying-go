package main

type Piece struct {
	Title    string
	Composer string
}

func counterstream() Piece {
	return Piece{"The People United Will Never Be Defeated!", "Frederic Rzewski"}
}
