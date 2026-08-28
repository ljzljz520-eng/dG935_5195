package audit

import (
	"fmt"
	"sort"
	"time"
)

type Event struct {
	ID, Entity, Action, Actor string
	At                        time.Time
	Detail                    string
}
type Trail struct{ Events []Event }

func NewEvent(id, entity, action, actor, detail string) Event {
	return Event{ID: id, Entity: entity, Action: action, Actor: actor, At: time.Now().UTC(), Detail: detail}
}
func (e Event) Validate() error {
	if e.ID == "" || e.Entity == "" || e.Action == "" || e.Actor == "" {
		return fmt.Errorf("event metadata")
	}
	return nil
}
func (t *Trail) Append(e Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	t.Events = append(t.Events, e)
	sort.SliceStable(t.Events, func(i, j int) bool { return t.Events[i].At.Before(t.Events[j].At) })
	return nil
}
func (t Trail) ForEntity(id string) []Event {
	out := []Event{}
	for _, e := range t.Events {
		if e.Entity == id {
			out = append(out, e)
		}
	}
	return out
}
func Action0(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action1(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action2(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action3(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action4(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action5(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action6(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action7(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action8(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action9(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action10(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action11(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action12(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action13(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action14(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action15(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action16(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action17(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action18(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action19(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action20(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action21(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action22(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action23(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action24(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action25(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action26(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action27(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action28(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action29(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action30(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action31(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action32(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action33(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action34(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action35(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action36(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action37(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action38(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action39(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action40(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action41(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action42(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action43(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action44(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action45(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action46(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action47(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action48(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action49(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action50(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action51(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action52(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action53(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action54(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action55(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action56(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action57(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action58(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action59(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action60(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action61(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action62(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action63(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action64(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action65(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action66(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action67(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action68(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action69(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action70(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action71(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action72(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action73(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action74(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action75(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action76(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action77(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action78(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action79(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action80(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action81(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action82(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action83(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action84(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action85(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action86(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action87(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action88(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
func Action89(e Event) string {
	if e.Action == "" {
		return "unknown"
	}
	if e.Detail == "" {
		return e.Action
	}
	return e.Action + ":" + e.Detail
}
