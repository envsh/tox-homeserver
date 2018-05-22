package store

import (
	"gopp"
	"strconv"
)

func (this *Storage) GetContactByPubkey(pubkey string) *Contact {
	ct := &Contact{}
	ct.Pubkey = pubkey
	_, err := this.dbh.Get(ct)
	gopp.ErrPrint(err)
	return ct
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
	ct := this.GetContactByPubkey(pubkey)
	err := this.dbh.Where("contact_id=?", ct.Id).Limit(10).Desc("id").Find(&msgs)
	gopp.ErrPrint(err, pubkey)
	return msgs
}

func (this *Storage) GetMessagesFrom(pubkey string, sfrom string) []Message {
	msgs := []Message{}
	from, err := strconv.ParseInt(sfrom, 10, 64)
	gopp.ErrPrint(err, sfrom, from)

	ct := this.GetContactByPubkey(pubkey)
	err = this.dbh.Where("contact_id=? AND id<?", ct.Id, from).Limit(13).Desc("id").Find(&msgs)
	gopp.ErrPrint(err, ct)
	return msgs
}

func (this *Storage) GetLastMessage(pubkey string, msg string) *Message {
	msgr := Message{}
	ct := this.GetContactByPubkey(pubkey)
	_, err := this.dbh.Where("contact_id=? AND content=?", ct.Id, msg).Limit(1).Desc("id").Get(&msgr)
	gopp.ErrPrint(err, pubkey, msg)
	return &msgr
}
