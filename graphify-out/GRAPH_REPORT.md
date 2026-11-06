# Graph Report - server  (2026-08-07)

## Corpus Check
- Corpus is ~47,686 words - fits in a single context window. You may not need a graph.

## Summary
- 780 nodes · 1876 edges · 24 communities (23 shown, 1 thin omitted)
- Extraction: 84% EXTRACTED · 16% INFERRED · 0% AMBIGUOUS · INFERRED: 293 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Controller Layer (ArticlesAuth)|Controller Layer (Articles/Auth)]]
- [[_COMMUNITY_RequestResponse DTO Patterns|Request/Response DTO Patterns]]
- [[_COMMUNITY_Content Submission Module|Content Submission Module]]
- [[_COMMUNITY_Cabinet Content Controllers|Cabinet Content Controllers]]
- [[_COMMUNITY_Domain Entities & Common Types|Domain Entities & Common Types]]
- [[_COMMUNITY_Auth & User Module|Auth & User Module]]
- [[_COMMUNITY_Article Domain (entityDTOcontroller)|Article Domain (entity/DTO/controller)]]
- [[_COMMUNITY_Bootstrap & Config|Bootstrap & Config]]
- [[_COMMUNITY_Letter Submission Module|Letter Submission Module]]
- [[_COMMUNITY_Cron Reminder & Notifications|Cron Reminder & Notifications]]
- [[_COMMUNITY_Medinfo PJ Queue Module|Medinfo PJ Queue Module]]
- [[_COMMUNITY_Letter Template Module|Letter Template Module]]
- [[_COMMUNITY_User Service & Validation Tests|User Service & Validation Tests]]
- [[_COMMUNITY_Validation Rules (SubmissionsArticles)|Validation Rules (Submissions/Articles)]]
- [[_COMMUNITY_Migration Manager|Migration Manager]]
- [[_COMMUNITY_Docs Controller & Service|Docs Controller & Service]]
- [[_COMMUNITY_Controller Integration Tests|Controller Integration Tests]]
- [[_COMMUNITY_WA Client Integration|WA Client Integration]]
- [[_COMMUNITY_Repository Tests|Repository Tests]]
- [[_COMMUNITY_Main package|Main package]]

## God Nodes (most connected - your core abstractions)
1. `BuildResponseSuccess()` - 54 edges
2. `BuildResponseFailed()` - 46 edges
3. `User` - 36 edges
4. `CurrentClaims()` - 32 edges
5. `cabinetService` - 30 edges
6. `Controller` - 27 edges
7. `cabinetRepository` - 26 edges
8. `New()` - 25 edges
9. `OK()` - 24 edges
10. `UserRepository` - 22 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `NewArticleRepository()`  [INFERRED]
  cmd/main.go → modules/article/repository/repository.go
- `main()` --calls--> `NewContentSubmissionRepository()`  [INFERRED]
  cmd/main.go → modules/content_submission/repository/repository.go
- `main()` --calls--> `StartDailyCron()`  [INFERRED]
  cmd/main.go → modules/cron/daily_reminder.go
- `main()` --calls--> `NewLetterSubmissionRepository()`  [INFERRED]
  cmd/main.go → modules/letter_submission/repository/repository.go
- `main()` --calls--> `NewLetterTemplateRepository()`  [INFERRED]
  cmd/main.go → modules/letter_template/repository/repository.go

## Import Cycles
- None detected.

## Communities (24 total, 1 thin omitted)

### Community 0 - "Controller Layer (Articles/Auth)"
Cohesion: 0.06
Nodes (36): ArticleController, AuthController, ContentSubmissionController, LetterSubmissionController, LetterTemplateController, MedinfoPJController, UserController, Context (+28 more)

### Community 1 - "Request/Response DTO Patterns"
Cohesion: 0.08
Nodes (50): CabinetRequest, CabinetResponse, DocumentationRequest, DocumentationResponse, MediaRequest, MediaResponse, MemberRequest, MemberResponse (+42 more)

### Community 2 - "Content Submission Module"
Cohesion: 0.07
Nodes (41): Time, AttachmentResponse, ContentSubmissionResponse, ContentSubmission, ContentSubmissionAssignmentHistory, ContentSubmissionAttachment, ContentSubmissionStatusHistory, AssignPJRequest (+33 more)

### Community 3 - "Cabinet Content Controllers"
Cohesion: 0.09
Nodes (24): Controller, MediaSubmissionSetting, CurrentClaims(), Context, bind(), Context, New(), paramID() (+16 more)

### Community 4 - "Domain Entities & Common Types"
Cohesion: 0.08
Nodes (25): Time, Time, Time, Time, CabinetTerm, MediaAsset, Ministry, OrganizationMember (+17 more)

### Community 5 - "Auth & User Module"
Cohesion: 0.06
Nodes (27): LoginRequest, LoginResponse, RegisterRequest, User, NewAuthController(), UserResponse, NewUserResponse(), RouterGroup (+19 more)

### Community 6 - "Article Domain (entity/DTO/controller)"
Cohesion: 0.07
Nodes (29): Time, ArticleListItem, ArticleResponse, AuthorSummary, PublishRequest, Article, NewArticleController(), CreateRequest (+21 more)

### Community 7 - "Bootstrap & Config"
Cohesion: 0.07
Nodes (38): main(), runDatabaseCommand(), Config, env(), envInt(), Load(), Connect(), DB (+30 more)

### Community 8 - "Letter Submission Module"
Cohesion: 0.09
Nodes (27): Time, LetterSubmissionResponse, LetterSubmission, LetterSubmissionAssignmentHistory, LetterSubmissionStatusHistory, AssignPJRequest, CreateRequest, StatusHistoryResponse (+19 more)

### Community 9 - "Cron Reminder & Notifications"
Cohesion: 0.13
Nodes (36): contentRow, letterRow, Duration, DB, ReminderItemsFromSubmissions(), StartDailyCron(), stringValue(), AdminDetailURL() (+28 more)

### Community 10 - "Medinfo PJ Queue Module"
Cohesion: 0.09
Nodes (25): QueueResponse, ReorderRequest, MedinfoPJQueue, CreateRequest, UserSummary, NewAvailabilityResponses(), NewQueueResponse(), NewQueueResponses() (+17 more)

### Community 11 - "Letter Template Module"
Cohesion: 0.11
Nodes (15): LetterTemplateResponse, TemplateFileResponse, LetterTemplate, CreateRequest, UpdateRequest, Time, NewLetterTemplateResponse(), NewLetterTemplateResponses() (+7 more)

### Community 12 - "User Service & Validation Tests"
Cohesion: 0.12
Nodes (16): UserCreateRequest, UserUpdateRequest, UserResponse, NewUserResponse(), NewUserResponses(), UserResponse, NewUserService(), T (+8 more)

### Community 13 - "Validation Rules (Submissions/Articles)"
Cohesion: 0.23
Nodes (16): CreateRequest, Time, hasHost(), isBlank(), isGoogleDrive(), isValidType(), NewContentSubmissionValidation(), startOfDay() (+8 more)

### Community 14 - "Migration Manager"
Cohesion: 0.25
Nodes (8): DB, Time, migrationNames(), NewMigrationManager(), splitSQLStatements(), MigrationManager, migrationRecord, MigrationStatus

### Community 15 - "Docs Controller & Service"
Cohesion: 0.20
Nodes (7): DocsController, Context, NewDocsController(), RouterGroup, RegisterRoutes(), NewDocsService(), DocsService

### Community 16 - "Controller Integration Tests"
Cohesion: 0.38
Nodes (12): Engine, doRequest(), fixTimestampColumns(), DB, T, newTestDB(), newTestRouter(), seedPublishedProgram() (+4 more)

### Community 17 - "WA Client Integration"
Cohesion: 0.26
Nodes (7): Client, roundTripFunc, NewWAClient(), previewWAMessage(), T, TestWAClientSendsBearerRequest(), WAClient

### Community 18 - "Repository Tests"
Cohesion: 0.56
Nodes (9): New(), fixTimestampColumns(), DB, T, newTestDB(), seedPublishedProgram(), TestProgramBySlugOnlyFound(), TestProgramBySlugOnlyNotFound() (+1 more)

## Knowledge Gaps
- **14 isolated node(s):** `bemunair2026/server`, `CreateRequest`, `UpdateRequest`, `PublishRequest`, `UpdateStatusRequest` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `User` connect `Auth & User Module` to `Content Submission Module`, `Domain Entities & Common Types`, `Article Domain (entity/DTO/controller)`, `Letter Submission Module`, `Cron Reminder & Notifications`, `Medinfo PJ Queue Module`, `User Service & Validation Tests`?**
  _High betweenness centrality (0.180) - this node is a cross-community bridge._
- **Why does `main()` connect `Bootstrap & Config` to `Content Submission Module`, `Cabinet Content Controllers`, `Auth & User Module`, `Article Domain (entity/DTO/controller)`, `Letter Submission Module`, `Cron Reminder & Notifications`, `Medinfo PJ Queue Module`, `Letter Template Module`, `WA Client Integration`?**
  _High betweenness centrality (0.150) - this node is a cross-community bridge._
- **Why does `BuildResponseSuccess()` connect `Controller Layer (Articles/Auth)` to `Cabinet Content Controllers`?**
  _High betweenness centrality (0.109) - this node is a cross-community bridge._
- **Are the 49 inferred relationships involving `BuildResponseSuccess()` (e.g. with `.Create()` and `.Delete()`) actually correct?**
  _`BuildResponseSuccess()` has 49 INFERRED edges - model-reasoned connections that need verification._
- **Are the 43 inferred relationships involving `BuildResponseFailed()` (e.g. with `.Create()` and `.Delete()`) actually correct?**
  _`BuildResponseFailed()` has 43 INFERRED edges - model-reasoned connections that need verification._
- **Are the 29 inferred relationships involving `CurrentClaims()` (e.g. with `.Create()` and `.Me()`) actually correct?**
  _`CurrentClaims()` has 29 INFERRED edges - model-reasoned connections that need verification._
- **What connects `bemunair2026/server`, `CreateRequest`, `UpdateRequest` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._