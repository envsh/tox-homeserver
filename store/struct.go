package store

type Contact struct {
	Id         int64  `xorm:"pk autoincr INTEGER"`
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
	Id        int64  `xorm:"pk autoincr INTEGER"`
	Updated   string `xorm:"TEXT"`
	Created   string `xorm:"TEXT"`
	Content   string `xorm:"TEXT"`
	Mtype     int    `xorm:"INTEGER"`
	ContactId int64  `xorm:"index INTEGER"`
	RoomId    int64  `xorm:"index INTEGER"` // it's really another ContactId
	EventId   int64  `xorm:"unique INTEGER"`
}

type MessageJoined struct {
	Id        int64  `xorm:"pk autoincr INTEGER"`
	Updated   string `xorm:"TEXT"`
	Created   string `xorm:"TEXT"`
	Content   string `xorm:"TEXT"`
	Mtype     int    `xorm:"INTEGER"`
	ContactId int64  `xorm:"index INTEGER"`
	RoomId    int64  `xorm:"index INTEGER"` // it's really another ContactId
	EventId   int64  `xorm:"unique INTEGER"`

	RoomName   string
	RoomPubkey string
	PeerName   string
	PeerPubkey string
}

// for server
type Device struct {
	Id        int64  `xorm:"pk autoincr INTEGER"`
	ContactId int64  `xorm:"index INTEGER"`
	Uuid      string `xorm:"unique TEXT"`
	Created   string `xorm:"TEXT"`
	Updated   string `xorm:"TEXT"`
}

type Idgen struct {
	Id int64 `xorm:"pk autoincr INTEGER"`
}

// for client
type SyncInfo struct {
	Id        int64  `xorm:"pk autoincr INTEGER"`
	CtId      int64  `xorm:"unique(siu) INTEGER"`
	NextBatch int64  `xorm:"unique(siu) INTEGER"`
	PrevBatch int64  `xorm:"unique(siu) INTEGER"`
	Updated   string `xorm:"TEXT"`
}

const (
	SK_DEVICE_NAME    = "device_name"
	SK_HOMESERVER_URL = "homeserver_url" // it's really last
	SK_SHOW_SQL_INLOG = "show_sql_inlog"
	SK_DEBUG_LEVEL    = "debug_level"
)

var SettingKeys = []string{"device_name", "homeserver_url", "show_sql", "debug_level"}

// for client
type Setting struct {
	Id int64 `xorm:"pk autoincr INTEGER"`
	// case for multiple accounts
	// ContactId int64  `xorm:"index ITNEGER"`
	Name    string `xorm:"unique TEXT"`
	Value   string `xorm:"TEXT"`
	Created string `xorm:"TEXT"`
	Updated string `xorm:"TEXT"`
}
