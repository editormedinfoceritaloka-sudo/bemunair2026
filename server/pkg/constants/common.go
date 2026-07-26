package constants

const (
	RoleAdmin        = "ADMIN"
	RoleAdminMedinfo = "ADMIN_MEDINFO"
	// RoleMentri remains a JWT compatibility alias for tokens issued before migration 008.
	RoleMentri = "MENTRI"

	StatusDraft             = "DRAFT"
	StatusSubmitted         = "SUBMITTED"
	StatusPendingReview     = "PENDING_REVIEW"
	StatusRevisionRequired  = "REVISION_REQUIRED"
	StatusRevisionSubmitted = "REVISION_SUBMITTED"
	StatusApproved          = "APPROVED"
	StatusScheduled         = "SCHEDULED"
	StatusPublished         = "PUBLISHED"
	StatusRejected          = "REJECTED"
	StatusCompleted         = "COMPLETED"

	// Legacy values are accepted by compatibility paths while existing clients migrate.
	StatusPending  = "PENDING"
	StatusInReview = "IN_REVIEW"

	ServiceTypeContent = "CONTENT"
	ServiceTypeArticle = "ARTICLE"

	ContentTypeFeed       = "FEED_INSTAGRAM"
	ContentTypeReels      = "REELS_INSTAGRAM"
	ContentTypeInstastory = "INSTASTORY"
	ContentTypeLegacy     = "FEEDS_REELS_LEGACY"
	ContentTypeFeedsReels = "FEEDS_REELS"
	ContentTypeArtikel    = "ARTIKEL"

	ArticleStatusDraft     = "DRAFT"
	ArticleStatusPublished = "PUBLISHED"
)
