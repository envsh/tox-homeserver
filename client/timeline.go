package client

import (
	"encoding/json"
	"fmt"
	"gopp"
	"sort"
	"tox-homeserver/store"
	"tox-homeserver/thspbs"
)

func (this *LigTox) PullEventsByContactId(pubkey string, prev_batch int64) ([]store.MessageJoined, error) {
	args := &thspbs.Event{}
	args.EventName = "PullEventsByContactId"
	args.Args = []string{pubkey, fmt.Sprintf("%d", prev_batch)}

	rsp, err := this.rmtCall(args)
	gopp.ErrPrint(err)
	if err != nil {
		return nil, err
	}

	rets := []store.MessageJoined{}
	err = json.Unmarshal([]byte(rsp.Args[0]), &rets)
	gopp.ErrPrint(err)
	if err != nil {
		return nil, err
	}
	return rets, nil
}

/////
type TimeLine struct {
	NextBatch int64 // not found now, so must bigger than PrevBatch, not equal.
	PrevBatch int64
}

func (this *TimeLine) Equal(that *TimeLine) bool {
	return this.NextBatch == that.NextBatch &&
		this.PrevBatch == that.PrevBatch
}
func (this *TimeLine) IsValid() bool { return this.NextBatch > this.PrevBatch }
func (this *TimeLine) Merge(that *TimeLine) (mrgtm *TimeLine, can bool) {
	if !this.IsValid() || !that.IsValid() {
		return
	}

	points := []int64{this.NextBatch, this.PrevBatch, that.NextBatch, that.PrevBatch}
	sort.SliceStable(points, func(i int, j int) bool { return points[i] > points[j] })

	if points[0] == this.NextBatch && points[1] == this.PrevBatch {
		if points[1] == points[2] {
		} else {
			return
		}
	} else if points[0] == that.NextBatch && points[1] == that.PrevBatch {
		if points[1] == points[2] {
			// can
		} else {
			return
		}
	}

	return &TimeLine{points[0], points[3]}, true
}

func MergeTimeLines(tls ...*TimeLine) (mrgtm *TimeLine, can bool) {
	if len(tls) == 0 {
		return
	}
	if len(tls) == 1 {
		mrgtm = &*tls[0]
		can = true
		return
	}
	this := tls[0]
	for i := 1; i < len(tls); i++ {
		mrgtm, can = this.Merge(tls[i])
	}
	return
}

// merge as many as possible, return merged timeline, and mrgcnt
// no touch of params
func MergeTimeLinesCount(btl *TimeLine, tls []*TimeLine) (mrgtl *TimeLine, mrgcnt int) {
	mrgtl = btl
	for _, tl := range tls {
		newtl, can := mrgtl.Merge(tl)

		if can {
			mrgtl = newtl
			mrgcnt += 1
		} else {
			break
		}
	}
	return
}

func SyncInfo2TimeLine(si *store.SyncInfo) *TimeLine {
	if si == nil {
		return nil
	}
	return &TimeLine{int64(si.NextBatch), int64(si.PrevBatch)}
}

func SyncInfos2TimeLines(sis []store.SyncInfo) (tms []*TimeLine) {
	for _, si := range sis {
		tms = append(tms, SyncInfo2TimeLine(&si))
	}
	return
}
