package def


type User struct {
    Accid   string  `gorm:"not null;unique_index"`
    Name    string
    Gold    int
	Info    string
}

type NormalCraft struct {
    Id      int
    Author  string
    Rect    int
    Data    string
    Praise  int
}

type GoodCraft struct {
    Id      int
    Author  string
    Rect    int
    Data    string
    Praise  int
}
