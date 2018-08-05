package store

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"runtime"
	thscom "tox-homeserver/common"

	"github.com/go-xorm/xorm"
	// "github.com/hashicorp/go-uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

type Storage struct {
	dbh *xorm.Engine
}

func NewStorage() *Storage {
	this := &Storage{}
	var dsn string
	if runtime.GOOS == "android" {
		dsn = fmt.Sprintf("file:///data/data/io.dnesth.tofia/toxhs.sqlite")
		dsn = fmt.Sprintf("file://%s/toxhs.sqlite", thscom.AndroidGetDataDir())
		log.Println(dsn)
	} else {
		dsn = fmt.Sprintf("toxhs.sqlite?cache=shared&mode=rwc")
	}
	dbh, err := xorm.NewEngine("sqlite3", dsn)
	gopp.ErrPrint(err)
	err = dbh.Ping()
	gopp.ErrPrint(err, dsn)
	if err == nil {
		log.Println("PING DATABASE OK:", dsn)
	}
	this.dbh = dbh
	this.SetWAL(true)

	logger := xorm.NewSimpleLogger2(os.Stdout, thscom.LogPrefix, 0)
	dbh.SetLogger(logger)
	dbh.ShowSQL(false)
	this.initTables()
	return this
}

// -wal when can delete this file? if lost, does it broken database?
// -shm when can delete this file? if lost, does it broken database?
func (this *Storage) SetWAL(enable bool) {
	_, err := this.dbh.Exec("PRAGMA journal_mode=WAL;")
	gopp.ErrPrint(err)
	_, err = this.dbh.Exec(fmt.Sprintf("PRAGMA journal_size_limit=%d;", 3*1000*1000)) // 3MB
	gopp.ErrPrint(err)
	// others: wal_checkpoint, wal_autocheckpoint, synchronous, cache_size
	if true {
		return
	}
	_, err = this.dbh.Exec("PRAGMA locking_mode=EXCLUSIVE;")
	gopp.ErrPrint(err)
}

func (this *Storage) initTables() {
	tbls, err := this.dbh.DBMetas()
	gopp.ErrPrint(err, len(tbls))

	dmrecs := []interface{}{&Contact{}, &Message{}, &Device{}, &Idgen{}, &SyncInfo{}, &Setting{}}
	for _, dmrec := range dmrecs {
		recval := fmt.Sprintf("%+v", dmrec)
		if ok, err := this.dbh.IsTableExist(dmrec); !ok && err == nil {
			err := this.dbh.CreateTables(dmrec)
			gopp.ErrPrint(err, recval)
		} else {
			err := this.dbh.Sync(dmrec)
			gopp.ErrPrint(err, recval)
		}

		if true {
			err = this.dbh.CreateUniques(dmrec)
			gopp.FalsePrint(IsIndexExistErr(err), err)

			err = this.dbh.CreateIndexes(dmrec)
			gopp.FalsePrint(IsIndexExistErr(err), err)
		}
	}
}

func (this *Storage) Close() error {
	return this.dbh.Close()
}

/////
func (this *Storage) AddFriend(pubkey string, rtnum uint32, name, stmsg string) (int64, error) {
	c := &Contact{}
	c.Pubkey = pubkey
	c.RtId = rtnum
	c.Name = name
	c.Stmsg = stmsg
	c.Status = 0
	c.IsFriend = 1
	return this.AddContact(c)
}

func (this *Storage) AddGroup(identify string, rtnum uint32, title string) (int64, error) {
	c := &Contact{}
	c.IsGroup = 1
	c.Pubkey = identify
	c.RtId = rtnum
	c.Name = title
	return this.AddContact(c)
}

func (this *Storage) UpdateGroup(identifier string, rtnum uint32, title string) (int64, error) {
	c := &Contact{}
	c.IsGroup = 1
	c.Pubkey = identifier
	c.RtId = rtnum
	c.Name = title
	return this.UpdateContactByPubkey(c)
}

func (this *Storage) SetGroup(identifier string, rtnum uint32, title string) (ret int64, err error) {
	ret, err = this.AddGroup(identifier, rtnum, title)
	gopp.ErrPrint(err, title)
	if IsUniqueConstraintErr(err) {
		ret, err = this.UpdateGroup(identifier, rtnum, title)
		gopp.ErrPrint(err, title)
	}
	return
}

func (this *Storage) AddPeer(peerPubkey string, rtnum uint32, name string) (int64, error) {
	c := &Contact{}
	c.IsPeer = 1
	c.Pubkey = peerPubkey
	c.RtId = rtnum
	c.Name = name
	return this.AddContact(c)
}

func (this *Storage) UpdatePeer(peerPubkey string, rtnum uint32, name string) (int64, error) {
	c := &Contact{}
	c.Pubkey = peerPubkey
	c.RtId = rtnum
	c.Name = name
	return this.UpdateContactByPubkey(c)
}

func (this *Storage) AddPeerOrUpdateName(peerPubkey string, rtnum uint32, name string) (int64, error) {
	no, err := this.AddPeer(peerPubkey, rtnum, name)
	if IsUniqueConstraintErr(err) {
		c := &Contact{}
		c.IsPeer = 1
		c.RtId = rtnum
		c.Name = name
		return this.dbh.Where("pubkey=? and name != ?", peerPubkey, name).Update(c)
	}
	return no, err
}

func (this *Storage) AddContact(c *Contact) (int64, error) {
	c.Created = thscom.NowTimeStr()
	c.Updated = c.Created

	id, err := this.dbh.InsertOne(c)
	return id, err
}

func (this *Storage) UpdateContactByPubkey(c *Contact) (int64, error) {
	c.Updated = thscom.NowTimeStr()

	id, err := this.dbh.Where("pubkey = ?", c.Pubkey).Update(c)
	gopp.ErrPrint(err, id)
	return id, err
}

// eventId 参数可选，为0表示服务器使用，自动生成
// friendpk => room id, pubkey => peer(contact)
func (this *Storage) AddFriendMessage(msg string, friendpk, pubkey string, eventId int64, userCode int64) (*Message, error) {
	c := &Contact{}
	c.Pubkey = pubkey
	exist, err := this.dbh.Get(c)
	gopp.ErrPrint(err, exist, pubkey)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, xorm.ErrNotExist
	}

	c2 := &Contact{}
	c2.Pubkey = friendpk
	exist, err = this.dbh.Get(c2)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, xorm.ErrNotExist
	}

	m := &Message{}
	m.Content = msg
	m.ContactId = c.Id
	m.RoomId = c2.Id // for friend, room id is contact id. contact is a room
	m.EventId = eventId
	m.UserCode = userCode
	return this.AddMessage(m)
}

// eventId 参数可选，为0表示服务器使用，自动生成
func (this *Storage) AddGroupMessage(msg string, mtype string, identify string, peerPubkey string, eventId int64, userCode int64) (*Message, error) {
	c0 := &Contact{}
	c0.Pubkey = identify
	exist, err := this.dbh.Get(c0)
	gopp.ErrPrint(err, exist, identify)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, xorm.ErrNotExist
	}

	c1 := &Contact{}
	c1.Pubkey = peerPubkey
	exist, err = this.dbh.Get(c1)
	gopp.ErrPrint(err, exist, identify)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, xorm.ErrNotExist
	}

	m := &Message{}
	m.Content = msg
	m.ContactId = c1.Id
	m.RoomId = c0.Id
	m.EventId = eventId
	m.UserCode = userCode
	return this.AddMessage(m)
}

//
func (this *Storage) AddMessage(m *Message) (*Message, error) {
	m.Updated = thscom.NowTimeStr()
	// m.EventId <=0认为是server端，否则客户端
	if m.EventId <= 0 {
		m.Created = m.Updated
		m.EventId = this.NextId()
	}

	id, err := this.dbh.InsertOne(m)
	_ = id
	return m, err
}

func (this *Storage) AddMessageJoined(mj *MessageJoined) (*Message, error) {
	mj.Updated = thscom.NowTimeStr()
	// m.EventId <=0认为是server端，否则客户端
	if mj.EventId <= 0 {
		mj.Created = mj.Updated
		mj.EventId = this.NextId()
	}

	m := &Message{}
	m.EventId = mj.EventId
	m.Content = mj.Content
	m.Mtype = mj.Mtype
	m.Created = mj.Created
	m.Updated = mj.Updated

	ct, _ := this.GetContactByPubkey(mj.PeerPubkey)
	if ct == nil {
		ct = &Contact{}
		ct.Pubkey = mj.PeerPubkey
		ct.Name = mj.PeerName
		_, err := this.AddContact(ct)
		gopp.ErrPrint(err)
	}
	if ct != nil {
		m.ContactId = ct.Id
	}

	room, _ := this.GetContactByPubkey(mj.RoomPubkey)
	if room == nil {
		room = &Contact{}
		room.Pubkey = mj.RoomPubkey
		room.Name = mj.RoomName
		_, err := this.AddContact(room)
		gopp.ErrPrint(err)
	}
	if room != nil {
		m.RoomId = room.Id
	}

	id, err := this.dbh.InsertOne(m)
	_ = id
	return m, err
}

func (this *Storage) SetMessageSent(msgid int64) error {
	mo := &Message{}
	mo.Sent = 1
	mo.Updated = thscom.NowTimeStr()

	_, err := this.dbh.Where("id=?", msgid).Update(mo)
	gopp.ErrPrint(err, msgid)
	return err
}

func (this *Storage) FindUnsentMessages() (unsents []Message, err error) {
	err = this.dbh.Where("sent=?", 0).Asc("id").Find(&unsents)
	return
}

func (this *Storage) FindUnsentMessagesByPubkey(pubkey string) (unsents []Message, err error) {
	c, err := this.GetContactByPubkey(pubkey)
	if err != nil {
		return
	}
	err = this.dbh.Where("room_id=? and sent=?", c.Id, 0).Asc("id").Find(&unsents)
	return
}

func (this *Storage) MaxEventId() (int64, error) {
	r := &Message{}
	exists, err := this.dbh.Desc("event_id").Limit(1).Get(r)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, xorm.ErrNotExist
	}
	return int64(r.EventId), nil
}

func (this *Storage) FindEventsByContactId(pubkey string, prev_batch int64, page_size int) ([]Message, error) {
	c, err := this.GetContactByPubkey(pubkey)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, xorm.ErrNotExist
	}

	r := []Message{}
	err = this.dbh.Where("room_id = ? and event_id <= ?", c.Id, prev_batch).
		Desc("event_id").Limit(page_size).Find(&r)
	gopp.ErrPrint(err)
	return r, err
}

/*
select message.*, ctroom.name as room_name, ctpeer.name as peer_name,ctpeer.pubkey as peer_pubkey from message
left join contact ctroom on message.room_id = ctroom.id
left join contact ctpeer on message.contact_id=ctpeer.id
where ctroom.pubkey = 'FD056CBxxxxxx order by message.event_id desc limit 20
*/
func (this *Storage) FindEventsByContactId2(pubkey string, prev_batch int64, page_size int) ([]MessageJoined, error) {
	sqltpl := `select message.*, ctroom.name as room_name, ctroom.pubkey as room_pubkey,
 ctpeer.name as peer_name,ctpeer.pubkey as peer_pubkey from message
left join contact ctroom on message.room_id = ctroom.id
left join contact ctpeer on message.contact_id=ctpeer.id
where ctroom.pubkey = ? and event_id <= ? order by message.event_id desc limit ?`

	r := []MessageJoined{}
	err := this.dbh.SQL(sqltpl, pubkey, prev_batch, page_size).Find(&r)
	for i := 0; i < len(r); i++ {
		if r[i].PeerName == "" {
			r[i].PeerName = fmt.Sprintf("%s.%s", thscom.DefaultUserName, r[i].PeerPubkey[:7])
		}
	}
	return r, err
}

func (this *Storage) AddDevice() error {
	return this.AddDevice2(gopp.Retn(uuid.NewV4())[0].(string))
}

func (this *Storage) AddDevice2(name string) error {
	dv := Device{}
	dv.Uuid = name
	dv.Created = thscom.NowTimeStr()
	dv.Updated = dv.Created

	id, err := this.dbh.InsertOne(&dv)
	gopp.ErrPrint(err, id)
	return err
}

func (this *Storage) DeviceEmpty() bool {
	dv := &Device{}
	empty, err := this.dbh.IsTableEmpty(dv)
	gopp.ErrPrint(err, empty)
	return err == nil && empty
}

func (this *Storage) GetDevice() *Device {
	dv := &Device{}
	_, err := this.dbh.Get(dv)
	gopp.ErrPrint(err)
	if err != nil {
		return nil
	}
	return dv
}

func (this *Storage) NextId() int64 {
	idv := &Idgen{}
	affected, err := this.dbh.InsertOne(idv)
	gopp.ErrPrint(err, affected)
	return idv.Id
}

func (this *Storage) AddSyncInfo(ct_id int64, next_batch int64, prev_batch int64) error {
	dv := SyncInfo{}
	dv.CtId = ct_id
	dv.NextBatch = next_batch
	dv.PrevBatch = prev_batch
	dv.Updated = thscom.NowTimeStr()

	id, err := this.dbh.InsertOne(&dv)
	gopp.ErrPrint(err, id)
	return err
}

func (this *Storage) FindSyncInfoByCtId(ct_id int64) ([]SyncInfo, error) {
	c := []SyncInfo{}
	err := this.dbh.Where("ct_id = ?", ct_id).Desc("next_batch").Find(c)
	gopp.ErrPrint(err, ct_id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (this *Storage) UpdateSyncInfo(ct_id int64, next_batch int64, prev_batch int64) error {
	c := &SyncInfo{}
	c.CtId = ct_id
	c.NextBatch = next_batch
	c.PrevBatch = prev_batch
	c.Updated = thscom.NowTimeStr()

	_, err := this.dbh.Where("ct_id = ?", ct_id).Update(c)
	gopp.ErrPrint(err, ct_id)
	return err
}

func (this *Storage) DeleteSyncInfoByCtId(ct_id int64) error {
	c := &SyncInfo{}
	c.CtId = ct_id
	_, err := this.dbh.Delete(c)
	gopp.ErrPrint(err, ct_id)
	return err
}

func (this *Storage) DeleteSyncInfoById(id int64) error {
	c := &SyncInfo{}
	c.Id = id
	_, err := this.dbh.Delete(c)
	gopp.ErrPrint(err, id)
	return err
}

func (this *Storage) SetSetting(name, value string) (int64, error) {
	c := &Setting{}

	exist, err := this.dbh.Where("name = ?", name).Get(c)
	if err != nil {
		return 0, err
	}

	c = &Setting{}
	c.Name = name
	c.Value = value
	c.Updated = thscom.NowTimeStr()

	if exist {
		return this.dbh.Where("name = ?", name).Update(c)
	} else {
		c.Created = c.Updated
		return this.dbh.InsertOne(c)
	}
}

func (this *Storage) GetSetting(name string) (setting *Setting, err error) {
	c := &Setting{}
	exist, err := this.dbh.Where("name = ?", name).Get(c)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, xorm.ErrNotExist
	}
	return c, nil
}

func (this *Storage) GetSettings() (settings []Setting, err error) {
	err = this.dbh.Where("1==1").Find(&settings)
	return
}

func init() {
	if false {
		xorm.NewEngine("", "")
		log.Println()
	}
}
