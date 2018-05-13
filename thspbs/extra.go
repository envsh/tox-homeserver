package thspbs

// extra for proto auto generated
type ContactInfo = MemberInfo
type ContactType = MemberInfo_MemType

const CTFriend = MemberInfo_FRIEND
const CTGroup = MemberInfo_GROUP
const CTPeer = MemberInfo_PEER

func (this *MemberInfo) IsFriend() bool { return this.Mtype == MemberInfo_FRIEND }
func (this *MemberInfo) IsGroup() bool  { return this.Mtype == MemberInfo_GROUP }
func (this *MemberInfo) IsPeer() bool   { return this.Mtype == MemberInfo_PEER }
