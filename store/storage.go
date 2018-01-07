package store

import (
	"fmt"
	"gopp"
	"log"
	"os"
	"runtime"
	"time"
	"tox-homeserver/common"

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
	} else {
		dsn = fmt.Sprintf("toxhs.sqlite")
	}
	dbh, err := xorm.NewEngine("sqlite3", dsn)
	gopp.ErrPrint(err)
	err = dbh.Ping()
	gopp.ErrPrint(err, dsn)

	logger := xorm.NewSimpleLogger2(os.Stdout, common.LogPrefix, 0)
	dbh.SetLogger(logger)
	this.dbh = dbh
	dbh.ShowSQL(true)
	this.initTables()
	return this
}

func (this *Storage) initTables() {
	tbls, err := this.dbh.DBMetas()
	gopp.ErrPrint(err, len(tbls))

	ctrec := &Contact{}
	msgrec := &Message{}
	dvrec := &Device{}

	if ok, err := this.dbh.IsTableExist(ctrec); !ok && err == nil {
		this.dbh.CreateTables(ctrec)
	}
	if ok, err := this.dbh.IsTableExist(msgrec); !ok && err == nil {
		this.dbh.CreateTables(msgrec)
	}
	if ok, err := this.dbh.IsTableExist(dvrec); !ok && err == nil {
		this.dbh.CreateTables(dvrec)
	}

	if true {
		err = this.dbh.CreateUniques(ctrec)
		gopp.ErrPrint(err)
		err = this.dbh.CreateUniques(msgrec)
		gopp.ErrPrint(err)
		err = this.dbh.CreateUniques(dvrec)
		gopp.ErrPrint(err)

		err = this.dbh.CreateIndexes(ctrec)
		gopp.ErrPrint(err)
		err = this.dbh.CreateIndexes(msgrec)
		gopp.ErrPrint(err)
		err = this.dbh.CreateIndexes(dvrec)
		gopp.ErrPrint(err)
	}
}

func (this *Storage) AddFriend(pubkey string, num uint32, name, stmsg string) (int64, error) {
	c := &Contact{}
	c.Pubkey = pubkey
	c.RtId = int(num)
	c.Name = name
	c.Stmsg = stmsg
	c.Status = 0
	return this.AddContact(c)
}

func (this *Storage) AddGroup(identify string, num uint32, title string) (int64, error) {
	c := &Contact{}
	c.IsGroup = 2
	c.Pubkey = identify
	c.RtId = int(num)
	c.Name = title
	return this.AddContact(c)
}

func (this *Storage) AddPeer(peerPubkey string, num uint32) (int64, error) {
	c := &Contact{}
	c.IsPeer = 3
	c.Pubkey = peerPubkey
	c.RtId = int(num)
	return this.AddContact(c)
}

func (this *Storage) AddContact(c *Contact) (int64, error) {
	nowt := time.Now().String()
	c.Created = nowt
	c.Updated = nowt

	id, err := this.dbh.InsertOne(c)
	gopp.ErrPrint(err, id)
	return id, err
}

func (this *Storage) AddFriendMessage(msg string, pubkey string) (int64, error) {
	c := &Contact{}
	c.Pubkey = pubkey
	exist, err := this.dbh.Get(c)
	gopp.ErrPrint(err, exist, pubkey)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, fmt.Errorf("not found: %s", pubkey)
	}

	m := &Message{}
	m.Content = msg
	m.ContactId = c.Id
	return this.AddMessage(m)
}

func (this *Storage) AddGroupMessage(msg string, mtype string, identify string, peerPubkey string) (int64, error) {
	c0 := &Contact{}
	c0.Pubkey = identify
	exist, err := this.dbh.Get(c0)
	gopp.ErrPrint(err, exist, identify)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, fmt.Errorf("not found: %s", identify)
	}

	c1 := &Contact{}
	c1.Pubkey = peerPubkey
	exist, err = this.dbh.Get(c1)
	gopp.ErrPrint(err, exist, identify)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, fmt.Errorf("not found: %s", peerPubkey)
	}

	m := &Message{}
	m.Content = msg
	m.ContactId = c1.Id
	m.RoomId = c0.Id
	return this.AddMessage(m)
}

func (this *Storage) AddMessage(m *Message) (int64, error) {
	nowt := time.Now().String()
	m.Created = nowt
	m.Updated = nowt
	id, err := this.dbh.InsertOne(m)
	gopp.ErrPrint(err, id)
	return id, err
}

func (this *Storage) AddDevice() error {
	dv := Device{}
	dv.Uuid = uuid.NewV4().String()
	dv.Created = time.Now().String()
	dv.Updated = time.Now().String()

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

func init() {
	if false {
		xorm.NewEngine("", "")
		log.Println()
	}
}
