package msg

type C_S_HeartBeat struct {
}

type S_C_HeartBeat struct {
}

type C_S_Login struct {
	UserID int
}

type S_C_Login struct {
	ErrCode     int
	UserID      int
	Gold        int64
	Coin        int64
	Money       int64
	HeadIconUrl string
	UnderWrite  string
	Sex         int
	ActiveFlag  int64
	NickName    string
}

type C_S_JoinRoom struct {
	RoomId  int
	Service string
}

type S_C_JoinRoom struct {
	ErrCode int
	RoomId  int
}
