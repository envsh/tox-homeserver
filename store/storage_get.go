package store

import (
	"gopp"
	"strconv"
)

func (this *Storage) GetContactByPubkey(pubkey string) (*Contact, error) {
	ct := &Contact{}
	ct.Pubkey = pubkey
	_, err := this.dbh.Get(ct)
	gopp.ErrPrint(err)
	return ct, err
}

func (this *Storage) GetContactById(ct_id int) *Contact {
	ct := &Contact{}
	ct.Id = ct_id
	_, err := this.dbh.Get(ct)
	gopp.ErrPrint(err)
	return ct
}

func (this *Storage) GetAllContacts() []*Contact {
	var contacts []Contact
	err := this.dbh.Find(&contacts)
	gopp.ErrPrint(err)
	pcontacts := []*Contact{}
	for _, ct := range contacts {
		nct := ct
		pcontacts = append(pcontacts, &nct)
	}
	return pcontacts
}

func (this *Storage) GetLatestMessages(pubkey string) []Message {
	msgs := []Message{}
	ct, _ := this.GetContactByPubkey(pubkey)
	err := this.dbh.Where("contact_id=?", ct.Id).Limit(10).Desc("id").Find(&msgs)
	gopp.ErrPrint(err, pubkey)
	return msgs
}

func (this *Storage) GetMessagesFrom(pubkey string, sfrom string) []Message {
	msgs := []Message{}
	from, err := strconv.ParseInt(sfrom, 10, 64)
	gopp.ErrPrint(err, sfrom, from)

	ct, _ := this.GetContactByPubkey(pubkey)
	err = this.dbh.Where("contact_id=? AND id<?", ct.Id, from).Limit(13).Desc("id").Find(&msgs)
	gopp.ErrPrint(err, ct)
	return msgs
}

func (this *Storage) GetLastMessage(pubkey string, msg string) *Message {
	msgr := Message{}
	ct, _ := this.GetContactByPubkey(pubkey)
	_, err := this.dbh.Where("contact_id=? AND content=?", ct.Id, msg).Limit(1).Desc("id").Get(&msgr)
	gopp.ErrPrint(err, pubkey, msg)
	return &msgr
}

func (this *Storage) GetTimeLinesByPubkey(pubkey string) (rets []SyncInfo, err error) {
	c, err := this.GetContactByPubkey(pubkey)
	gopp.ErrPrint(err, pubkey)
	if err != nil {
		return
	}

	err = this.dbh.Where("ct_id = ?", c.Id).Desc("next_batch").Find(&rets)
	gopp.ErrPrint(err, pubkey, c)
	return
}

func (this *Storage) DeleteTimeLinesByPubkey(pubkey string) (err error) {
	c, err := this.GetContactByPubkey(pubkey)
	gopp.ErrPrint(err, pubkey)
	if err != nil {
		return
	}

	n, err := this.dbh.Where("ct_id = ?", c.Id).Delete(&SyncInfo{})
	gopp.ErrPrint(err, n, pubkey, c)
	return
}
