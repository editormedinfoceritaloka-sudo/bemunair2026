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
	submitterPhone := "6282"
	submitter := &entities.User{ID: 2, Name: "Submitter", Phone: &submitterPhone}
	now := time.Date(2026, time.July, 30, 10, 30, 0, 0, time.FixedZone("WIB", 7*60*60))
	code := "MED-2026-000123"
	sub := &entities.ContentSubmission{RequestCode: &code, ServiceType: "CONTENT", SubmissionType: "FEED_INSTAGRAM", Title: "Konten Kegiatan", Ministry: "MEDINFO", Status: "SUBMITTED", Deadline: &now}
	wa := &mockWA{}
	errs := NotifyContentSubmissionCreated(sub, nil, submitter, wa)
	if len(errs) != 0 || len(wa.phones) != 1 {
		t.Fatalf("wa calls = %d errs=%d", len(wa.phones), len(errs))
	}
	for _, expected := range []string{"*BEM UNAIR 2026 — Pengajuan Diterima*", "MED-2026-000123", "Pengajuan Konten", "Menunggu Peninjauan"} {
		if !strings.Contains(wa.messages[0], expected) {
			t.Fatalf("message missing %q:\n%s", expected, wa.messages[0])
		}
	}
}

func TestBuildAssignmentMessageIsCompleteAndWhatsAppFriendly(t *testing.T) {
	message := BuildAssignmentMessage("Alya", AssignmentItem{RequestCode: "MED-2026-000124", Service: "Pengajuan Artikel", Title: "Airlangga Festival 2026", Ministry: "SENORA", SubmitterName: "Bima", Status: "SUBMITTED", Deadline: time.Date(2026, time.July, 29, 15, 0, 0, 0, time.FixedZone("WIB", 7*60*60)), DetailURL: "https://admin.bem.unair.ac.id/admin/content-submissions/124"})
	for _, expected := range []string{"*BEM UNAIR 2026 — Penugasan PJ*", "Halo, Alya.", "• Kode: MED-2026-000124", "• Layanan: Pengajuan Artikel", "• Judul/Perihal: Airlangga Festival 2026", "• Pengaju: Bima", "• Kementerian: SENORA", "Rabu, 29 Juli 2026 pukul 15.00 WIB", "*Tindak Lanjut*", "🔗 Detail: https://admin.bem.unair.ac.id/admin/content-submissions/124"} {
		if !strings.Contains(message, expected) {
			t.Errorf("message missing %q:\n%s", expected, message)
		}
	}
	if strings.Contains(message, "|") {
		t.Fatalf("message must not use table-style separators:\n%s", message)
	}
}

func TestApproachingDeadlineUsesThreeDayWindowAndOneDayGrace(t *testing.T) {
	now := time.Date(2026, time.July, 26, 12, 0, 0, 0, time.UTC)
	items := []ReminderItem{{RequestCode: "OVERDUE-TOO-LONG", Deadline: now.Add(-25 * time.Hour)}, {RequestCode: "OVERDUE", Deadline: now.Add(-12 * time.Hour)}, {RequestCode: "SOON", Deadline: now.Add(48 * time.Hour)}, {RequestCode: "TOO-FAR", Deadline: now.Add(73 * time.Hour)}}
	filtered := ApproachingDeadline(items, now)
	if len(filtered) != 2 || filtered[0].RequestCode != "OVERDUE" || filtered[1].RequestCode != "SOON" {
		t.Fatalf("unexpected filtered reminders: %#v", filtered)
	}
}

func TestNotifyDeadlineReminderTargetsEachPJAndSkipsEmptyGroup(t *testing.T) {
	phone := "628111"
	pj := &entities.User{ID: 7, Name: "Nadia", Phone: &phone}
	now := time.Date(2026, time.July, 26, 12, 0, 0, 0, time.UTC)
	items := []ReminderItem{{RequestCode: "MED-2026-000125", Service: "Pengajuan Konten", Title: "Open Recruitment", Ministry: "PSDM", Status: "PENDING_REVIEW", Deadline: now.Add(20 * time.Hour), PJ: pj}}
	wa := &mockWA{}
	NotifyDailyPendingReminder(items, now, wa)
	NotifyGroupDailyReminder(items, now, wa, "", "medinfo@g.us")
	if len(wa.phones) != 1 || wa.phones[0] != phone {
		t.Fatalf("personal recipients = %#v", wa.phones)
	}
	if len(wa.groups) != 1 || wa.groups[0] != "medinfo@g.us" {
		t.Fatalf("group recipients = %#v", wa.groups)
	}
	for _, expected := range []string{"*BEM UNAIR 2026 — Pengingat Deadline*", "20 jam lagi", "Sedang Ditinjau"} {
		if !strings.Contains(wa.messages[0], expected) {
			t.Errorf("personal reminder missing %q:\n%s", expected, wa.messages[0])
		}
	}
}

func TestNoReminderMessageIsSentForEmptyItems(t *testing.T) {
	wa := &mockWA{}
	NotifyDailyPendingReminder(nil, time.Now(), wa)
	NotifyGroupDailyReminder(nil, time.Now(), wa, "medinfo@g.us")
	if len(wa.messages) != 0 {
		t.Fatalf("empty reminder sent %d messages", len(wa.messages))
	}
}
