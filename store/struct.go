package store

type Contact struct {
	Id         int    `xorm:"pk autoincr INTEGER"`
	Pubkey     string `xorm:"not null unique TEXT"`
	Name       string `xorm:"index TEXT"`
	Stmsg      string `xorm:"TEXT"`
	LastSeen   string `xorm:"TEXT"`
	Status     int    `xorm:"INTEGER"`
	ConnStatus int    `xorm:"INTEGER"`
	Created    string `xorm:"TEXT"`
	Updated    string `xorm:"TEXT"`
	IsGroup    int    `xorm:"index INTEGER"`
	RtId       int    `xorm:"INTEGER"`
	IsPeer     int    `xorm:"index INTEGER"`
	ChatType   int    `xorm:"INTEGER"`
	IsFriend   int    `xorm:"index INTEGER"`
}

type Message struct {
	Id        int    `xorm:"pk autoincr INTEGER"`
	Updated   string `xorm:"TEXT"`
	Created   string `xorm:"TEXT"`
	Content   string `xorm:"TEXT"`
	Mtype     int    `xorm:"INTEGER"`
	ContactId int    `xorm:"index INTEGER"`
	RoomId    int    `xorm:"index INTEGER"`
	EventId   int64  `xorm:"unique INTEGER"`
}

type Device struct {
	Id        int    `xorm:"pk autoincr INTEGER"`
	ContactId int    `xorm:"index INTEGER"`
	Uuid      string `xorm:"unique TEXT"`
	Created   string `xorm:"TEXT"`
	Updated   string `xorm:"TEXT"`
}

type Idgen struct {
	Id int64 `xorm:"pk autoincr INTEGER"`
}

type SyncInfo struct {
	Id        int    `xorm:"pk autoincr INTEGER"`
	CtId      int    `xorm:"unique INTEGER"`
	NextBatch int    `xorm:"INTEGER"`
	PrevBatch int    `xorm:"INTEGER"`
	Updated   string `xorm:"TEXT"`
}
