package wa_notification

import (
	"fmt"
	"os"
	"strings"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/pkg"
)

const (
	defaultReminderWindow = 72 * time.Hour
	overdueReminderGrace  = 24 * time.Hour
)

type AssignmentItem struct {
	RequestCode, Service, Title, Ministry, SubmitterName, Status, DetailURL string
	Deadline                                                                time.Time
}

type ReminderItem struct {
	RequestCode, Service, Title, Ministry, Status, DetailURL string
	Deadline                                                 time.Time
	PJ                                                       *entities.User
}

func NotifyContentSubmissionCreated(s *entities.ContentSubmission, pj, submitter *entities.User, wa pkg.WASender) []error {
	var errs []error
	if wa == nil || submitter == nil {
		return errs
	}
	if pj != nil {
		item := AssignmentItem{RequestCode: value(s.RequestCode), Service: contentServiceLabel(s.ServiceType, s.SubmissionType), Title: s.Title, Ministry: s.Ministry, SubmitterName: submitter.Name, Status: s.Status, DetailURL: AdminDetailURL("content-submissions", s.ID)}
		if s.Deadline != nil {
			item.Deadline = *s.Deadline
		}
		if err := NotifyAssignedPJ(pj, item, wa); err != nil {
			errs = append(errs, err)
		}
	}
	if submitter.Phone != nil {
		deadline := "-"
		if s.Deadline != nil {
			deadline = formatTime(*s.Deadline)
		}
		msg := fmt.Sprintf("*BEM UNAIR 2026 — Pengajuan Diterima*\n\nHalo, %s.\nPengajuan Anda telah tercatat dan menunggu peninjauan tim Medinfo.\n\n*Detail Pengajuan*\n• Kode: %s\n• Layanan: %s\n• Judul: %s\n• Jadwal: %s\n• Status: %s\n\nTim Medinfo akan menghubungi Anda melalui WhatsApp apabila diperlukan revisi.\n\n_Pesan otomatis — mohon tidak membalas pesan ini._", submitter.Name, fallback(value(s.RequestCode), "-"), contentServiceLabel(s.ServiceType, s.SubmissionType), fallback(s.Title, "-"), deadline, statusLabel(s.Status))
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
	if pj != nil {
		if err := NotifyAssignedPJ(pj, AssignmentItem{RequestCode: value(s.RequestCode), Service: "Pengajuan Surat", Title: s.Subject, Ministry: s.Ministry, SubmitterName: submitter.Name, Status: s.Status, Deadline: s.Deadline, DetailURL: AdminDetailURL("letter-submissions", s.ID)}, wa); err != nil {
			errs = append(errs, err)
		}
	}
	if submitter.Phone != nil {
		msg := fmt.Sprintf("*BEM UNAIR 2026 — Pengajuan Diterima*\n\nHalo, %s.\nPengajuan surat Anda telah tercatat dan menunggu peninjauan tim terkait.\n\n*Detail Pengajuan*\n• Kode: %s\n• Jenis: %s\n• Perihal: %s\n• Deadline: %s\n• Status: %s\n\nTim terkait akan menghubungi Anda melalui WhatsApp apabila diperlukan revisi.\n\n_Pesan otomatis — mohon tidak membalas pesan ini._", submitter.Name, fallback(value(s.RequestCode), "-"), fallback(s.LetterType, "-"), fallback(s.Subject, "-"), formatTime(s.Deadline), statusLabel(s.Status))
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
	msg := fmt.Sprintf("*BEM UNAIR 2026 — Pembaruan Status*\n\nHalo, %s.\nStatus pengajuan Anda telah diperbarui menjadi *%s*.\n\nSilakan buka dashboard untuk melihat detail dan catatan terbaru.\n\n_Pesan otomatis — mohon tidak membalas pesan ini._", name, statusLabel(status))
	return wa.SendTextMessage(phone, msg)
}

func BuildAssignmentMessage(pjName string, item AssignmentItem) string {
	lines := []string{"*BEM UNAIR 2026 — Penugasan PJ*", "", fmt.Sprintf("Halo, %s.", fallback(pjName, "Rekan PJ")), "Anda ditetapkan sebagai penanggung jawab pengajuan berikut:", "", "*Detail Tugas*", fmt.Sprintf("• Kode: %s", fallback(item.RequestCode, "-")), fmt.Sprintf("• Layanan: %s", fallback(item.Service, "-")), fmt.Sprintf("• Judul/Perihal: %s", fallback(item.Title, "-")), fmt.Sprintf("• Pengaju: %s", fallback(item.SubmitterName, "-")), fmt.Sprintf("• Kementerian: %s", fallback(item.Ministry, "-"))}
	if !item.Deadline.IsZero() {
		lines = append(lines, fmt.Sprintf("• Deadline: %s", formatTime(item.Deadline)))
	}
	lines = append(lines, fmt.Sprintf("• Status: %s", statusLabel(item.Status)), "", "*Tindak Lanjut*", "1. Periksa kelengkapan pengajuan.", "2. Hubungi pengaju jika ada data yang perlu diperbaiki.", "3. Perbarui status dan catatan melalui dashboard.")
	if item.DetailURL != "" {
		lines = append(lines, "", fmt.Sprintf("🔗 Detail: %s", item.DetailURL))
	}
	lines = append(lines, "", "_Pesan otomatis — mohon tidak membalas pesan ini._")
	return strings.Join(lines, "\n")
}

func NotifyAssignedPJ(pj *entities.User, item AssignmentItem, wa pkg.WASender) error {
	if wa == nil || pj == nil || pj.Phone == nil || strings.TrimSpace(*pj.Phone) == "" {
		return nil
	}
	return wa.SendTextMessage(*pj.Phone, BuildAssignmentMessage(pj.Name, item))
}

func ApproachingDeadline(items []ReminderItem, now time.Time) []ReminderItem {
	return approachingDeadline(items, now, defaultReminderWindow, overdueReminderGrace)
}

func approachingDeadline(items []ReminderItem, now time.Time, window, overdueGrace time.Duration) []ReminderItem {
	start := now.Add(-overdueGrace)
	end := now.Add(window)
	filtered := make([]ReminderItem, 0, len(items))
	for _, item := range items {
		if item.Deadline.IsZero() || item.Deadline.Before(start) || item.Deadline.After(end) {
			continue
		}
		filtered = append(filtered, item)
	}
	return filtered
}

func BuildDeadlineReminder(recipientName string, items []ReminderItem, now time.Time) string {
	lines := []string{"*BEM UNAIR 2026 — Pengingat Deadline*", "", fmt.Sprintf("Halo, %s.", fallback(recipientName, "Rekan PJ")), fmt.Sprintf("Ada %d tugas aktif yang perlu segera ditindaklanjuti:", len(items))}
	for index, item := range items {
		lines = append(lines, "", fmt.Sprintf("*%d. %s*", index+1, fallback(item.RequestCode, "Pengajuan")), fmt.Sprintf("• %s", fallback(item.Service, "-")), fmt.Sprintf("• Judul: %s", fallback(item.Title, "-")), fmt.Sprintf("• Kementerian: %s", fallback(item.Ministry, "-")), fmt.Sprintf("• Deadline: %s", formatTime(item.Deadline)), fmt.Sprintf("• Sisa waktu: %s", deadlineDistance(now, item.Deadline)), fmt.Sprintf("• Status: %s", statusLabel(item.Status)))
		if item.DetailURL != "" {
			lines = append(lines, fmt.Sprintf("• Detail: %s", item.DetailURL))
		}
	}
	lines = append(lines, "", "⚠️ Mohon prioritaskan tugas yang paling dekat atau telah melewati deadline, lalu perbarui statusnya di dashboard.", "", "_Pesan otomatis — mohon tidak membalas pesan ini._")
	return strings.Join(lines, "\n")
}

func BuildGroupDeadlineReminder(items []ReminderItem, now time.Time) string {
	lines := []string{"*BEM UNAIR 2026 — Ringkasan Deadline Medinfo*", "", fmt.Sprintf("Ada %d pengajuan aktif yang mendekati deadline:", len(items))}
	for index, item := range items {
		pjName := "Belum ditetapkan"
		if item.PJ != nil {
			pjName = item.PJ.Name
		}
		lines = append(lines, "", fmt.Sprintf("*%d. %s — %s*", index+1, fallback(item.RequestCode, "Pengajuan"), fallback(item.Title, "-")), fmt.Sprintf("• Layanan: %s", fallback(item.Service, "-")), fmt.Sprintf("• Deadline: %s (%s)", formatTime(item.Deadline), deadlineDistance(now, item.Deadline)), fmt.Sprintf("• PJ: %s", pjName), fmt.Sprintf("• Status: %s", statusLabel(item.Status)))
	}
	lines = append(lines, "", "Mohon koordinasikan pengajuan yang belum memiliki PJ dan prioritaskan deadline terdekat.", "", "_Pesan otomatis dari sistem Admin BEM UNAIR 2026._")
	return strings.Join(lines, "\n")
}

func NotifyDailyPendingReminder(items []ReminderItem, now time.Time, wa pkg.WASender) []error {
	if wa == nil || len(items) == 0 {
		return nil
	}
	grouped := map[uint64][]ReminderItem{}
	pjs := map[uint64]*entities.User{}
	for _, item := range items {
		if item.PJ != nil && item.PJ.Phone != nil && strings.TrimSpace(*item.PJ.Phone) != "" {
			grouped[item.PJ.ID] = append(grouped[item.PJ.ID], item)
			pjs[item.PJ.ID] = item.PJ
		}
	}
	var errs []error
	for id, pending := range grouped {
		if err := wa.SendTextMessage(*pjs[id].Phone, BuildDeadlineReminder(pjs[id].Name, pending, now)); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func NotifyGroupDailyReminder(items []ReminderItem, now time.Time, wa pkg.WASender, groups ...string) []error {
	if wa == nil || len(items) == 0 {
		return nil
	}
	msg := BuildGroupDeadlineReminder(items, now)
	var errs []error
	for _, group := range groups {
		if strings.TrimSpace(group) == "" {
			continue
		}
		if err := wa.SendGroupMessage(group, msg); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func AdminDetailURL(resource string, id uint64) string {
	base := strings.TrimRight(strings.TrimSpace(os.Getenv("ADMIN_DASHBOARD_URL")), "/")
	if base == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%d", base, strings.Trim(resource, "/"), id)
}

func contentServiceLabel(serviceType, submissionType string) string {
	switch strings.ToUpper(serviceType) {
	case "ARTICLE":
		return "Pengajuan Artikel"
	case "CONTENT":
		return "Pengajuan Konten"
	}
	return titleWords(submissionType)
}

func statusLabel(status string) string {
	labels := map[string]string{"DRAFT": "Draft", "SUBMITTED": "Menunggu Peninjauan", "PENDING": "Menunggu Peninjauan", "PENDING_REVIEW": "Sedang Ditinjau", "IN_REVIEW": "Sedang Ditinjau", "REVISION_REQUIRED": "Perlu Revisi", "REVISION_SUBMITTED": "Revisi Dikirim", "APPROVED": "Disetujui", "SCHEDULED": "Dijadwalkan", "PUBLISHED": "Dipublikasikan", "COMPLETED": "Selesai", "REJECTED": "Ditolak"}
	if label, ok := labels[strings.ToUpper(strings.TrimSpace(status))]; ok {
		return label
	}
	return titleWords(status)
}

func deadlineDistance(now, deadline time.Time) string {
	distance := deadline.Sub(now)
	if distance < 0 {
		distance = -distance
		if distance < time.Hour {
			return "terlewat kurang dari 1 jam"
		}
		if distance < 24*time.Hour {
			return fmt.Sprintf("terlewat %d jam", int(distance.Round(time.Hour)/time.Hour))
		}
		return fmt.Sprintf("terlewat %d hari", int(distance.Round(24*time.Hour)/(24*time.Hour)))
	}
	if distance < time.Hour {
		return "kurang dari 1 jam"
	}
	if distance < 24*time.Hour {
		return fmt.Sprintf("%d jam lagi", int(distance.Round(time.Hour)/time.Hour))
	}
	return fmt.Sprintf("%d hari lagi", int(distance.Round(24*time.Hour)/(24*time.Hour)))
}

func formatTime(t time.Time) string {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.FixedZone("WIB", 7*60*60)
	}
	t = t.In(loc)
	weekdays := [...]string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	months := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	return fmt.Sprintf("%s, %02d %s %d pukul %02d.%02d WIB", weekdays[t.Weekday()], t.Day(), months[t.Month()-1], t.Year(), t.Hour(), t.Minute())
}

func fallback(value, fallbackValue string) string {
	if strings.TrimSpace(value) == "" {
		return fallbackValue
	}
	return strings.TrimSpace(value)
}

func titleWords(value string) string {
	words := strings.Fields(strings.ReplaceAll(strings.TrimSpace(value), "_", " "))
	for index := range words {
		words[index] = strings.ToUpper(words[index][:1]) + strings.ToLower(words[index][1:])
	}
	return strings.Join(words, " ")
}

func value(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
