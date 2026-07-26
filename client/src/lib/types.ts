export type Role = 'ADMIN' | 'ADMIN_MEDINFO';
export type SubmissionStatus = 'DRAFT' | 'SUBMITTED' | 'PENDING_REVIEW' | 'REVISION_REQUIRED' | 'REVISION_SUBMITTED' | 'APPROVED' | 'SCHEDULED' | 'PUBLISHED' | 'REJECTED' | 'COMPLETED' | 'PENDING' | 'IN_REVIEW';
export type ArticleStatus = 'DRAFT' | 'PUBLISHED';

export interface Meta { page: number; per_page: number; total: number; total_pages: number }
export interface ApiErrorBody { code?: string; details?: unknown }
export interface ApiEnvelope<T> { status: boolean; success?: boolean; message: string; data: T; meta?: Meta; error?: ApiErrorBody | string }
export interface User { id: number; name: string; email: string; role: Role; ministry_id?: number; ministry?: string; phone?: string }
export interface UserSummary extends User {}

export interface ContentSubmission {
  id: number; request_code?: string; service_type: 'CONTENT' | 'ARTICLE'; content_format?: 'FEED_INSTAGRAM' | 'REELS_INSTAGRAM' | 'INSTASTORY'; submitter_name: string; submitter_phone?: string; ministry_id?: number; submitter_id: number; submitter?: UserSummary; ministry: string;
  submission_type: 'FEEDS_REELS' | 'FEED_INSTAGRAM' | 'REELS_INSTAGRAM' | 'INSTASTORY' | 'ARTIKEL'; title: string;
  add_song?: string; caption: string; additional_notes?: string; publish_date?: string;
  publish_time?: string; design_drive_link?: string; canva_link?: string;
  article_drive_link?: string; deadline?: string; brief_link: string;
  assigned_pj_id?: number; assigned_pj?: UserSummary; status: SubmissionStatus;
  notes?: string; created_at: string; updated_at: string;
}

export interface LetterSubmission {
  id: number; submitter_id: number; submitter?: UserSummary; ministry: string;
  letter_type: string; subject: string; body: string; deadline: string;
  assigned_pj_id?: number; assigned_pj?: UserSummary; status: SubmissionStatus;
  notes?: string; created_at: string; updated_at: string;
}

export interface QueueItem { id: number; user_id: number; user?: UserSummary; position: number; is_current: boolean }
export interface LetterTemplate { id: number; name: string; type: string; subject: string; body: string; created_at: string; updated_at: string }
export interface Article {
  id: number; slug: string; title: string; excerpt?: string; body: string;
  cover_image?: string; author_id: number; author?: { id: number; name: string };
  status: ArticleStatus; published_at?: string; created_at: string; updated_at: string;
}
export interface Ministry { id: number; code: string; name: string; is_active: boolean; created_at: string; updated_at: string }
export interface MediaSubmissionSetting { ServiceType?: string; service_type?: string; SOPURL?: string; sop_url?: string; MinistryTemplateURL?: string; ministry_template_url?: string; BriefTemplateURL?: string; brief_template_url?: string; CaptionTemplateURL?: string; caption_template_url?: string; PICName?: string; pic_name?: string; PICWhatsApp?: string; pic_whatsapp?: string; MinimumLeadDays?: number; minimum_lead_days?: number; PublishTimeStart?: string; publish_time_start?: string; PublishTimeEnd?: string; publish_time_end?: string; SlotIntervalMinutes?: number; slot_interval_minutes?: number; terms: string[] }
export interface SubmissionHistory { id: number; actor?: UserSummary; from_status?: string; to_status: string; note?: string; created_at: string }

export interface LoginResult { token: string; user: User }
