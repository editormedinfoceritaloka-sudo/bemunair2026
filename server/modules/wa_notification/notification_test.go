package wa_notification

import (
	"strings"
	"testing"
	"time"

	"bemunair2026/server/database/entities"
)

type mockWA struct {
	phones   []string
	messages []string
	groups   []string
}

func (m *mockWA) SendTextMessage(phone, message string) error {
	m.phones = append(m.phones, phone)
	m.messages = append(m.messages, message)
	return nil
}
func (m *mockWA) SendGroupMessage(groupJID, message string) error {
	m.groups = append(m.groups, groupJID)
	m.messages = append(m.messages, message)
	return nil
}

func TestNotifyContentSubmissionCreated(t *testing.T) {
	pjPhone, submitterPhone := "6281", "6282"
	pj := &entities.User{ID: 1, Name: "PJ", Phone: &pjPhone}
	submitter := &entities.User{ID: 2, Name: "Submitter", Phone: &submitterPhone}
	now := time.Now()
	sub := &entities.ContentSubmission{SubmissionType: "FEEDS_REELS", Title: "Konten Kegiatan", Ministry: "MEDINFO", Deadline: &now}
	wa := &mockWA{}
	errs := NotifyContentSubmissionCreated(sub, pj, submitter, wa)
	if len(errs) != 0 || len(wa.phones) != 2 {
		t.Fatalf("wa calls = %d errs=%d", len(wa.phones), len(errs))
	}
	if !strings.Contains(wa.messages[0], "FEEDS_REELS") {
		t.Fatalf("message missing submission type: %s", wa.messages[0])
	}
}

func TestGroupReminder(t *testing.T) {
	wa := &mockWA{}
	items := []ReminderItem{{Type: "Feed", Ministry: "MEDINFO", Deadline: time.Now()}}
	NotifyGroupDailyReminder(items, wa, "g1@g.us", "g2@g.us")
	if len(wa.groups) != 2 {
		t.Fatalf("groups = %d", len(wa.groups))
	}
}
