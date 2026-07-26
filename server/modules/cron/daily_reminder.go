package cron

import (
	"log"
	"time"

	"bemunair2026/server/config"
	contentRepository "bemunair2026/server/modules/content_submission/repository"
	letterRepository "bemunair2026/server/modules/letter_submission/repository"
	"bemunair2026/server/modules/wa_notification"
	"bemunair2026/server/pkg"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func ReminderItemsFromSubmissions(contentRows []contentRow, letterRows []letterRow) []wa_notification.ReminderItem {
	items := make([]wa_notification.ReminderItem, 0, len(contentRows)+len(letterRows))
	for _, row := range contentRows {
		items = append(items, row.ReminderItem())
	}
	for _, row := range letterRows {
		items = append(items, row.ReminderItem())
	}
	return items
}

type contentRow interface {
	ReminderItem() wa_notification.ReminderItem
}
type letterRow interface {
	ReminderItem() wa_notification.ReminderItem
}

func StartDailyCron(db *gorm.DB, wa pkg.WASender, cfg *config.Config) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("Asia/Jakarta", 7*3600)
	}
	c := cron.New(cron.WithLocation(loc))
	contentRepo := contentRepository.NewContentSubmissionRepository(db)
	letterRepo := letterRepository.NewLetterSubmissionRepository(db)
	_, _ = c.AddFunc("0 12 * * *", func() {
		contentRows, err := contentRepo.ListPendingForReminder()
		if err != nil {
			log.Printf("cron content reminder error: %v", err)
			return
		}
		letterRows, err := letterRepo.ListPendingForReminder()
		if err != nil {
			log.Printf("cron letter reminder error: %v", err)
			return
		}
		items := make([]wa_notification.ReminderItem, 0, len(contentRows)+len(letterRows))
		for _, row := range contentRows {
			if row.Deadline == nil {
				continue
			}
			service := "Pengajuan Konten"
			if row.ServiceType == "ARTICLE" {
				service = "Pengajuan Artikel"
			}
			items = append(items, wa_notification.ReminderItem{RequestCode: stringValue(row.RequestCode), Service: service, Title: row.Title, Ministry: row.Ministry, Status: row.Status, Deadline: *row.Deadline, PJ: row.AssignedPJ, DetailURL: wa_notification.AdminDetailURL("content-submissions", row.ID)})
		}
		for _, row := range letterRows {
			items = append(items, wa_notification.ReminderItem{RequestCode: stringValue(row.RequestCode), Service: "Pengajuan Surat", Title: row.Subject, Ministry: row.Ministry, Status: row.Status, Deadline: row.Deadline, PJ: row.AssignedPJ, DetailURL: wa_notification.AdminDetailURL("letter-submissions", row.ID)})
		}
		now := time.Now().In(loc)
		items = wa_notification.ApproachingDeadline(items, now)
		if len(items) == 0 {
			return
		}
		for _, notifyErr := range wa_notification.NotifyDailyPendingReminder(items, now, wa) {
			log.Printf("cron personal reminder error: %v", notifyErr)
		}
		for _, notifyErr := range wa_notification.NotifyGroupDailyReminder(items, now, wa, cfg.WAGroupJID1, cfg.WAGroupJID2) {
			log.Printf("cron group reminder error: %v", notifyErr)
		}
	})
	c.Start()
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
