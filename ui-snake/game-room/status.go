package gameroom

type SnakeStatus struct {
	Board     Board      `json:"board"`
	PlayerOne []Position `json:"player_one"`
	PlayerTwo []Position `json:"player_two"`
	Fruit     Position   `json:"fruit"`
	GameOver  bool       `json:"game_over"`
	Winner    string     `json:"winner"`
}

type Board struct {
	Length int `json:"lenght"`
	Width  int `json:"width"`
}

type Position struct {
	X int `json:"X"`
	Y int `json:"Y"`
}
