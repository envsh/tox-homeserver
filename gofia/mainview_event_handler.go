package gofia

import (
	"log"
	"runtime"
	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"
)

func init() {
	log.SetPrefix("gofiat ")
	log.SetFlags(log.Flags() | log.Lshortfile)
	if runtime.GOOS == "android" {
		log.SetFlags(log.Flags() ^ log.Ltime)
		log.SetFlags(log.Flags() ^ log.Ldate)
	}
}

type mainViewState struct {
	nickName string
}

func (this *TutorialView) registerEvents() {
	vtc := appctx.vtcli
	this.mvst.nickName = "Tofia User"

	vtc.CallbackBaseInfo(func(_ *thscli.LigTox, bi *thspbs.BaseInfo, ud interface{}) {
		log.Println("hehrereh")
		this.mvst.nickName = bi.GetName() + "." + bi.GetId()[:5]

		for fn, frnd := range bi.Friends {
			cti := NewContactItem(false)
			cti.cnum = fn
			cti.ctid = frnd.GetPubkey()
			cti.ctname = frnd.GetName()
			cti.stmsg = frnd.GetStmsg()
			cti.status = frnd.GetStatus()

			this.contacts = append(this.contacts, cti)
			this.contactsv = append(this.contactsv, cti)

			appctx.cts.Put(cti.ctid, cti)
			cf := NewChatFormView()
			cf.cfst = cti.ContactItemState
			appctx.cfs.Put(cti.ctid, cf)
		}

		this.Signal()

	}, nil)
}
