package msg

type R_C_JoinRoom struct {
	RoomId int
	Seat   int
	State  int
}

type SyncPlayers struct {
	ID          int
	Gold        int64
	IpAddr      string
	HeadIconUrl string
	Sex         int
	NickName    string

	Seat  int
	State int
}

type R_C_SyncPlayers struct {
	Players []SyncPlayers
}

type C_R_SitDown struct {
}
