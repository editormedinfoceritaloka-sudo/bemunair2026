package wa_notification

import (
	"fmt"
	"strings"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/pkg"
)

func NotifyContentSubmissionCreated(s *entities.ContentSubmission, pj, submitter *entities.User, wa pkg.WASender) []error {
	var errs []error
	if wa == nil || submitter == nil {
		return errs
	}
	deadline := formatTime(s.Deadline)
	if pj != nil && pj.Phone != nil {
		msg := fmt.Sprintf("Halo %s\n\nAda pengajuan konten baru yang perlu kamu tangani.\n\nDari: %s (%s)\nPlatform: %s\nJenis: %s\nCaption: %s\nDeadline: %s\nStatus: PENDING\n\nSilakan cek dashboard BEM UNAIR segera.", pj.Name, submitter.Name, s.Ministry, s.Platform, s.SubmissionType, s.Caption, deadline)
		if err := wa.SendTextMessage(*pj.Phone, msg); err != nil {
			errs = append(errs, err)
		}
	}
	if submitter.Phone != nil {
		pjName := "-"
		if pj != nil {
			pjName = pj.Name
		}
		msg := fmt.Sprintf("Halo %s\n\nPengajuan kamu berhasil masuk ke sistem BEM UNAIR!\n\nPlatform: %s\nJenis: %s\nDeadline: %s\nStatus: PENDING\nPJ: %s\n\nKami akan menghubungimu jika ada update. Terima kasih!", submitter.Name, s.Platform, s.SubmissionType, deadline, pjName)
		if err := wa.SendTextMessage(*submitter.Phone, msg); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func NotifyLetterSubmissionCreated(s *entities.LetterSubmission, pj, submitter *entities.User, wa pkg.WASender) []error {
	var errs []error
	if wa == nil || submitter == nil {
		return errs
	}
	if pj != nil && pj.Phone != nil {
		msg := fmt.Sprintf("Halo %s\n\nAda pengajuan surat baru yang perlu kamu tangani.\n\nDari: %s (%s)\nJenis: %s\nSubject: %s\nDeadline: %s\nStatus: PENDING", pj.Name, submitter.Name, s.Ministry, s.LetterType, s.Subject, formatTime(s.Deadline))
		if err := wa.SendTextMessage(*pj.Phone, msg); err != nil {
			errs = append(errs, err)
		}
	}
	if submitter.Phone != nil {
		msg := fmt.Sprintf("Halo %s\n\nPengajuan surat kamu berhasil masuk ke sistem BEM UNAIR!\n\nJenis: %s\nDeadline: %s\nStatus: PENDING", submitter.Name, s.LetterType, formatTime(s.Deadline))
		if err := wa.SendTextMessage(*submitter.Phone, msg); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func NotifySubmissionStatusUpdated(phone, name, status string, wa pkg.WASender) error {
	if wa == nil || phone == "" {
		return nil
	}
	return wa.SendTextMessage(phone, fmt.Sprintf("Halo %s, status pengajuan kamu diperbarui menjadi %s.", name, status))
}

func NotifyAssignedPJ(pj *entities.User, summary string, wa pkg.WASender) error {
	if wa == nil || pj == nil || pj.Phone == nil {
		return nil
	}
	return wa.SendTextMessage(*pj.Phone, summary)
}

type ReminderItem struct {
	Type, Ministry string
	Deadline       time.Time
	PJ             *entities.User
}

func BuildDailyReminder(items []ReminderItem) string {
	lines := make([]string, 0, len(items))
	for _, item := range items {
		lines = append(lines, fmt.Sprintf("- %s dari %s (deadline: %s)", item.Type, item.Ministry, formatTime(item.Deadline)))
	}
	return fmt.Sprintf("Reminder Harian BEM UNAIR\n\nAda %d pengajuan yang masih pending:\n%s\n\nMohon segera ditindaklanjuti via dashboard.", len(items), strings.Join(lines, "\n"))
}

func NotifyDailyPendingReminder(items []ReminderItem, wa pkg.WASender) []error {
	grouped := map[uint64][]ReminderItem{}
	pjs := map[uint64]*entities.User{}
	for _, item := range items {
		if item.PJ != nil && item.PJ.Phone != nil {
			grouped[item.PJ.ID] = append(grouped[item.PJ.ID], item)
			pjs[item.PJ.ID] = item.PJ
		}
	}
	var errs []error
	for id, pending := range grouped {
		if err := wa.SendTextMessage(*pjs[id].Phone, BuildDailyReminder(pending)); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func NotifyGroupDailyReminder(items []ReminderItem, wa pkg.WASender, groups ...string) []error {
	msg := BuildDailyReminder(items)
	var errs []error
	for _, group := range groups {
		if err := wa.SendGroupMessage(group, msg); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func formatTime(t time.Time) string { return t.Format(time.RFC3339) }
