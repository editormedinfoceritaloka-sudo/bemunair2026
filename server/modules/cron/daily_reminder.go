package cron

import (
	"log"
	"time"

	"bemunair2026/server/config"
	content "bemunair2026/server/modules/content_submission"
	letter "bemunair2026/server/modules/letter_submission"
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
	contentRepo := content.NewRepository(db)
	letterRepo := letter.NewRepository(db)
	_, _ = c.AddFunc("0 12 * * *", func() {
		contentRows, err := contentRepo.ListPendingOlderThan(24 * time.Hour)
		if err != nil {
			log.Printf("cron content reminder error: %v", err)
			return
		}
		letterRows, err := letterRepo.ListPendingOlderThan(24 * time.Hour)
		if err != nil {
			log.Printf("cron letter reminder error: %v", err)
			return
		}
		items := make([]wa_notification.ReminderItem, 0, len(contentRows)+len(letterRows))
		for _, row := range contentRows {
			items = append(items, wa_notification.ReminderItem{Type: row.SubmissionType, Ministry: row.Ministry, Deadline: row.Deadline, PJ: row.AssignedPJ})
		}
		for _, row := range letterRows {
			items = append(items, wa_notification.ReminderItem{Type: row.LetterType, Ministry: row.Ministry, Deadline: row.Deadline, PJ: row.AssignedPJ})
		}
		wa_notification.NotifyDailyPendingReminder(items, wa)
		wa_notification.NotifyGroupDailyReminder(items, wa, cfg.WAGroupJID1, cfg.WAGroupJID2)
	})
	c.Start()
}
