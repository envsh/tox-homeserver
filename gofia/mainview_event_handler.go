package gofia

import (
	"log"
	"runtime"
	thscli "tox-homeserver/client"
	"tox-homeserver/thspbs"

	"github.com/kitech/godsts/lists/arraylist"
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
	// appctx.mvst.nickName = "Tofia User"

	vtc.CallbackBaseInfo(func(_ *thscli.LigTox, bi *thspbs.BaseInfo, ud interface{}) {
		log.Println("hehrereh")
		appctx.mvst.nickName = bi.GetName() + "." + bi.GetId()[:5]
		log.Println("hehrereh", appctx.mvst.nickName)

		for fn, frnd := range bi.Friends {
			cti := NewContactItem(false)
			cti.cnum = fn
			cti.ctid = frnd.GetPubkey()
			cti.ctname = frnd.GetName()
			cti.stmsg = frnd.GetStmsg()
			cti.status = frnd.GetStatus()

			// appctx.contacts = append(appctx.contacts, cti)
			// appctx.contactsv = append(appctx.contactsv, cti)

			// appctx.ctvs.Put(cti.ctid, cti)
			cf := NewChatFormView()
			cf.cfst = cti.ContactItemState
			// appctx.cfvs.Put(cti.ctid, cf)
			_, _ = cti, cf

			//
			ctis := &ContactItemState{}
			ctis.cnum = fn
			ctis.ctid = frnd.GetPubkey()
			ctis.ctname = frnd.GetName()
			ctis.stmsg = frnd.GetStmsg()
			ctis.status = frnd.GetStatus()
			ctis.msgs = arraylist.New()
			appctx.contactStates.Put(ctis.ctid, ctis)
			cfst := &*ctis
			appctx.chatFormStates.Put(ctis.ctid, cfst)
		}

		this.Signal()

	}, nil)
}
