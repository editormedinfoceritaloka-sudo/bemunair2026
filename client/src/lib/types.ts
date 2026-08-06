export type Role = 'ADMIN' | 'ADMIN_MEDINFO';
export type SubmissionStatus = 'DRAFT' | 'SUBMITTED' | 'PENDING_REVIEW' | 'REVISION_REQUIRED' | 'REVISION_SUBMITTED' | 'APPROVED' | 'SCHEDULED' | 'PUBLISHED' | 'REJECTED' | 'COMPLETED' | 'PENDING' | 'IN_REVIEW';
export type ArticleStatus = 'DRAFT' | 'PUBLISHED';

export interface Meta { page: number; per_page: number; total: number; total_pages: number }
export interface ApiErrorBody { code?: string; details?: unknown }
export interface ApiEnvelope<T> { status: boolean; success?: boolean; message: string; data: T; meta?: Meta; error?: ApiErrorBody | string }
export interface User { id: number; name: string; email: string; role: Role; ministry_id?: number; ministry?: string; phone?: string }
export type UserSummary = User

export interface ContentSubmission {
  id: number; request_code?: string; service_type: 'CONTENT' | 'ARTICLE'; content_format?: 'FEED_INSTAGRAM' | 'REELS_INSTAGRAM' | 'INSTASTORY'; submitter_name: string; submitter_phone?: string; ministry_id?: number; submitter_id: number; submitter?: UserSummary; ministry: string;
  submission_type: 'FEEDS_REELS' | 'FEED_INSTAGRAM' | 'REELS_INSTAGRAM' | 'INSTASTORY' | 'ARTIKEL'; title: string;
  add_song?: string; caption: string; additional_notes?: string; publish_date?: string;
  publish_time?: string; design_drive_link?: string; canva_link?: string;
  article_drive_link?: string; documentation_drive_link?: string; required_information?: string; deadline?: string; brief_link: string;
  assigned_pj_id?: number; assigned_pj?: UserSummary; status: SubmissionStatus;
  notes?: string; created_at: string; updated_at: string;
}

export interface LetterSubmission {
  id: number; request_code?: string; submitter_name: string; submitter_phone?: string; ministry_id?: number; submitter_id: number; submitter?: UserSummary; ministry: string;
  letter_type: string; subject: string; body: string; deadline: string;
  assigned_pj_id?: number; assigned_pj?: UserSummary; status: SubmissionStatus;
  notes?: string; created_at: string; updated_at: string;
}

export interface QueueItem { id: number; user_id: number; user?: UserSummary; position: number; is_current: boolean; is_busy: boolean; active_task_type?: "CONTENT" | "LETTER"; active_task_id?: number; active_request_code?: string; active_task_title?: string }
export interface LetterTemplate { id: number; name: string; type: string; subject: string; body: string; media_asset_id?: number; file?: { id: number; url: string; name: string; mime_type: string; size_bytes: number }; download_url?: string; is_active: boolean; display_order: number; created_at: string; updated_at: string }
export interface Article {
  id: number; slug: string; title: string; excerpt?: string; body: string;
  cover_image?: string; author_id: number; author?: { id: number; name: string };
  status: ArticleStatus; published_at?: string; created_at: string; updated_at: string;
}
export interface Ministry { id: number; code: string; name: string; is_active: boolean; created_at: string; updated_at: string }
export interface MediaAsset {
  id: number;
  file_id: string;
  url: string;
  thumbnail_url?: string | null;
  name: string;
  alt_text?: string | null;
  caption?: string | null;
  mime_type: string;
  size_bytes: number;
  width?: number | null;
  height?: number | null;
  purpose: string;
  status: string;
}
export type OrganizationPositionType =
  | 'MENKO'
  | 'MINISTER'
  | 'DIRECTOR_GENERAL';

export type OrganizationUnitType =
  | 'MENKO'
  | 'KEMENTERIAN'
  | 'BPII';

export interface OrganizationMember {
  id: number;
  name: string;
  position: string;
  position_type: OrganizationPositionType;
  biography?: string | null;
  quote?: string | null;
  photo?: MediaAsset | null;
  display_order: number;
  is_leader: boolean;
  is_active: boolean;
}

export interface OrganizationUnit {
  id: number;
  cabinet_term_id?: number;
  parent_id?: number | null;
  code: string;
  name: string;
  unit_type: OrganizationUnitType;
  slug: string;
  short_name?: string | null;
  description?: string | null;
  vision?: string | null;
  mission?: string | null;
  logo?: MediaAsset | null;
  cover?: MediaAsset | null;
  display_order: number;
  is_active: boolean;
  is_published: boolean;
  members?: OrganizationMember[];
  programs?: WorkProgram[];
  children?: OrganizationUnit[];
}
export interface Milestone { id: number; title: string; description?: string; start_date?: string; end_date?: string; status: string; display_order: number }
export interface Documentation { id: number; media?: MediaAsset; title?: string; caption?: string; taken_at?: string; display_order: number; is_cover: boolean }
export interface WorkProgram { id: number; ministry_id: number; ministry_name?: string; name: string; slug: string; short_description?: string; description?: string; objectives?: string; target_audience?: string; start_date?: string; end_date?: string; execution_month?: string; status: string; cover?: MediaAsset; display_order: number; is_featured: boolean; is_published: boolean; published_at?: string; milestones?: Milestone[]; documentations?: Documentation[] }
export interface Cabinet { id: number; name: string; slug: string; tagline?: string; description?: string; logo?: MediaAsset; hero?: MediaAsset; period_start?: string; period_end?: string; is_active: boolean; is_published: boolean; meta_title?: string; meta_description?: string; kemenkoan?: OrganizationUnit[] }
export interface MediaSubmissionSetting { ServiceType?: string; service_type?: string; SOPURL?: string; sop_url?: string; MinistryTemplateURL?: string; ministry_template_url?: string; BriefTemplateURL?: string; brief_template_url?: string; CaptionTemplateURL?: string; caption_template_url?: string; PICName?: string; pic_name?: string; PICWhatsApp?: string; pic_whatsapp?: string; MinimumLeadDays?: number; minimum_lead_days?: number; PublishTimeStart?: string; publish_time_start?: string; PublishTimeEnd?: string; publish_time_end?: string; SlotIntervalMinutes?: number; slot_interval_minutes?: number; terms: string[] }
export interface SubmissionHistory { id: number; event_type?: "STATUS_CHANGED" | "PJ_ASSIGNED" | "PJ_REASSIGNED"; actor?: UserSummary; from_status?: string; to_status?: string; from_pj?: UserSummary; to_pj?: UserSummary; note?: string; created_at: string }

export interface LoginResult { token: string; user: User }
