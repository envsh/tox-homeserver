package gofia

import (
	"fmt"
	"gopp"
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
	nickName  string
	netStatus uint32
}

func (this *MainView) registerEvents() {
	vtc := appctx.vtcli
	// appctx.mvst.nickName = "Tofia User"

	vtc.CallbackBaseInfo(func(_ *thscli.LigTox, bi *thspbs.BaseInfo, ud interface{}) {
		log.Println("hehrereh")
		appctx.mvst.nickName = bi.GetName() + "." + bi.GetId()[:5]
		log.Println("hehrereh", appctx.mvst.nickName)
		appctx.mvst.netStatus = bi.Status

		for fn, frnd := range bi.Friends {
			//
			ctis := newContactItemState()
			ctis.cnum = fn
			ctis.ctid = frnd.GetPubkey()
			ctis.ctname = frnd.GetName()
			ctis.stmsg = frnd.GetStmsg()
			ctis.status = frnd.GetStatus()
			// cti.ctis = ctis

			// cf := NewChatFormView()
			// cf.cfst = ctis
			// appctx.cfvs.Put(cti.ctid, cf)
			// _, _ = cti, cf

			appctx.contactStates.Put(ctis.ctid, ctis)
			cfst := &*ctis
			appctx.chatFormStates.Put(ctis.ctid, cfst)

			_, err := appctx.store.AddFriend(frnd.Pubkey, fn, frnd.Name, frnd.Stmsg)
			gopp.ErrPrint(err, frnd)
		}

		for gn, grp := range bi.Groups {
			ctis := newContactItemState()
			ctis.group = true
			ctis.cnum = gn
			ctis.ctid = fmt.Sprintf("%d", gn)
			ctis.ctid = grp.GroupId
			ctis.ctname = grp.Title
			ctis.stmsg = grp.Stmsg
			ctis.stmsg = grp.Title + "SS"

			appctx.contactStates.Put(ctis.ctid, ctis)
			ctsf := &*ctis
			appctx.chatFormStates.Put(ctis.ctid, ctsf)

			_, err := appctx.store.AddGroup(grp.GroupId, gn, grp.Title)
			gopp.ErrPrint(err, grp)
		}

		this.Signal()

	}, nil)
}
