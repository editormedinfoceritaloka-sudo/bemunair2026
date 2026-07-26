export type Role = 'ADMIN' | 'MENTRI';
export type SubmissionStatus = 'PENDING' | 'IN_REVIEW' | 'APPROVED' | 'REJECTED';
export type ArticleStatus = 'DRAFT' | 'PUBLISHED';

export interface Meta { page: number; per_page: number; total: number; total_pages: number }
export interface ApiErrorBody { code?: string; details?: unknown }
export interface ApiEnvelope<T> { status: boolean; success?: boolean; message: string; data: T; meta?: Meta; error?: ApiErrorBody | string }
export interface User { id: number; name: string; email: string; role: Role; ministry?: string; phone?: string }
export interface UserSummary extends User {}

export interface ContentSubmission {
  id: number; submitter_id: number; submitter?: UserSummary; ministry: string;
  submission_type: 'FEEDS_REELS' | 'INSTASTORY' | 'ARTIKEL'; title: string;
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
export interface LoginResult { token: string; user: User }
