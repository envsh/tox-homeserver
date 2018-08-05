package thspbs

import (
	"log"
	"time"
)

// extra for proto auto generated
type ContactInfo = MemberInfo
type ContactType = MemberInfo_MemType

const CTFriend = MemberInfo_FRIEND
const CTGroup = MemberInfo_GROUP
const CTPeer = MemberInfo_PEER

func (this *MemberInfo) IsFriend() bool { return this.Mtype == MemberInfo_FRIEND }
func (this *MemberInfo) IsGroup() bool  { return this.Mtype == MemberInfo_GROUP }
func (this *MemberInfo) IsPeer() bool   { return this.Mtype == MemberInfo_PEER }

// TODO mutex
func (this *BaseInfo) UpdatePeerInfo(grpnum uint32, groupId string, pubkey string, name string, rtnum uint32) {
	if this.Groups == nil {
		return
	}
	for _, grpo := range this.Groups {
		if grpo.GroupId == groupId {
			grpo.UpdatePeerInfo(pubkey, name, rtnum)
			return
		}
	}
	// not found
	grpo := NewGroupInfo()
	grpo.AddPeerInfo(pubkey, name, rtnum)
	// if there has grpnum, override it
	if grpo, ok := this.Groups[grpnum]; ok {
		log.Println("Override group info:", grpnum, grpo.Gnum, grpo.Title)
	}
	this.Groups[grpnum] = grpo
}

func (this *BaseInfo) DeletePeerInfo(gnum uint32, groupId, pubkey string) {
	for _, grpo := range this.Groups {
		if grpo.GroupId == groupId {
			grpo.DeletePeerInfo(pubkey)
			return
		}
	}
}

func NewGroupInfo() *GroupInfo {
	grpo := &GroupInfo{}
	grpo.Members = make(map[string]*MemberInfo)
	return grpo
}
func (this *GroupInfo) AddPeerInfo(pubkey string, name string, rtnum uint32) {
	peero := &MemberInfo{}
	peero.Name = name
	peero.Pubkey = pubkey
	peero.Pnum = rtnum
	peero.Joints = time.Now().Unix()
	if peero2, ok := this.Members[pubkey]; ok {
		log.Println("Override peer info:", rtnum, name, peero.Pnum, peero.Name, peero2.Pnum, peero2.Name)
	}
	this.Members[pubkey] = peero
}

func (this *GroupInfo) UpdatePeerInfo(pubkey string, name string, rtnum uint32) {
	for _, memo := range this.Members {
		if memo.Pubkey == pubkey {
			memo.Name = name
			memo.Pnum = rtnum
			return
		}
	}
	// not found
	this.AddPeerInfo(pubkey, name, rtnum)
}

func (this *GroupInfo) DeletePeerInfo(pubkey string) {
	for pnum, peero := range this.Members {
		if peero.Pubkey == pubkey {
			delete(this.Members, pnum)
			return
		}
	}
}

///
func (this *BaseInfo) GetGroupMembers(grpnum uint32) map[string]*MemberInfo {
	if mis, ok := this.Groups[grpnum]; ok {
		return mis.Members
	}
	return nil
}

func (this *BaseInfo) GetGroupMembersByPubkey(groupId string) map[string]*MemberInfo {
	for _, grpo := range this.Groups {
		if grpo.GroupId == groupId {
			return grpo.Members
		}
	}
	return nil
}

func (this *BaseInfo) GetGroup(grpnum uint32) *GroupInfo {
	if this.HasGroup(grpnum) {
		return this.Groups[grpnum]
	}
	return nil
}

func (this *BaseInfo) AddGroup(grpnum uint32, grpid string) bool {
	if this.HasGroup(grpnum) {
		log.Println("why group already exist:", grpnum, grpid)
		return false
	}
	grpo := NewGroupInfo()
	grpo.Gnum, grpo.GroupId = grpnum, grpid
	this.Groups[grpnum] = grpo
	return true
}
func (this *BaseInfo) HasGroup(grpnum uint32) bool {
	_, ok := this.Groups[grpnum]
	return ok
}
func (this *BaseInfo) RemoveGroup(grpnum uint32) bool {
	if _, ok := this.Groups[grpnum]; ok {
		delete(this.Groups, grpnum)
	}
	return true
}
func (this *BaseInfo) PeerCount(grpnum uint32) int {
	if grpo, ok := this.Groups[grpnum]; ok {
		return len(grpo.Members)
	}
	return 0
}
