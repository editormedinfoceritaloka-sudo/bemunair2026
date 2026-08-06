package service

import (
	"errors"
	"strings"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/cabinet/dto"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/pkg/utils"
	"gorm.io/gorm"
)

const PublicProgramPageSize = 8

var validStatuses = map[string]bool{"DRAFT": true, "PLANNED": true, "ONGOING": true, "COMPLETED": true, "POSTPONED": true, "CANCELLED": true, "ARCHIVED": true}

type Service interface {
	PublicCabinet() (*dto.CabinetResponse, error)
	PublicCabinetBySlug(slug string) (*dto.CabinetResponse, error)
	PublicUnit(slug string) (*dto.UnitResponse, error)
	PublicPrograms(slug string, page int) (repository.ListResult[dto.ProgramResponse], error)
	PublicProgram(unitSlug, programSlug string) (*dto.ProgramResponse, error)
	Cabinets(page, perPage int) (repository.ListResult[dto.CabinetResponse], error)
	CreateCabinet(req dto.CabinetRequest) (*dto.CabinetResponse, error)
	UpdateCabinet(id uint64, req dto.CabinetRequest) (*dto.CabinetResponse, error)
	Units(cabinetID uint64, unitType string, public bool) ([]dto.UnitResponse, error)
	CreateUnit(req dto.UnitRequest, actor *middlewares.Claims) (*dto.UnitResponse, error)
	UpdateUnit(id uint64, req dto.UnitRequest, actor *middlewares.Claims) (*dto.UnitResponse, error)
	Members(unitID uint64, public bool) ([]dto.MemberResponse, error)
	CreateMember(req dto.MemberRequest, actor *middlewares.Claims) (*dto.MemberResponse, error)
	UpdateMember(id uint64, req dto.MemberRequest, actor *middlewares.Claims) (*dto.MemberResponse, error)
	Programs(unitID uint64, page, perPage int, public bool) (repository.ListResult[dto.ProgramResponse], error)
	Program(id uint64) (*dto.ProgramResponse, error)
	CreateProgram(req dto.ProgramRequest, actor *middlewares.Claims) (*dto.ProgramResponse, error)
	UpdateProgram(id uint64, req dto.ProgramRequest, actor *middlewares.Claims) (*dto.ProgramResponse, error)
	PublishProgram(id uint64, published bool, actor *middlewares.Claims) (*dto.ProgramResponse, error)
	CreateMilestone(req dto.MilestoneRequest, actor *middlewares.Claims) (*dto.MilestoneResponse, error)
	CreateDocumentation(programID uint64, req dto.DocumentationRequest, actor *middlewares.Claims) (*dto.DocumentationResponse, error)
	ReorderDocumentation(programID uint64, ids []uint64, actor *middlewares.Claims) error
	CreateMedia(req dto.MediaRequest, actor *middlewares.Claims) (*dto.MediaResponse, error)
}

type cabinetService struct{ repo repository.Repository }

var _ Service = (*cabinetService)(nil)

func New(repo repository.Repository) Service { return &cabinetService{repo: repo} }

func (s *cabinetService) PublicCabinet() (*dto.CabinetResponse, error) {
	cabinet, err := s.repo.ActiveCabinet()
	if err != nil {
		return nil, err
	}
	if cabinet == nil {
		return nil, nil
	}
	response := dto.NewCabinetResponse(*cabinet)
	return &response, nil
}

func (s *cabinetService) PublicCabinetBySlug(slug string) (*dto.CabinetResponse, error) {
	cabinet, err := s.repo.CabinetBySlug(slug, true)
	if err != nil || cabinet == nil {
		return nil, err
	}
	return s.cabinetResponse(cabinet, true)
}

func (s *cabinetService) cabinetResponse(cabinet *entities.CabinetTerm, public bool) (*dto.CabinetResponse, error) {
	units, err := s.repo.Units(cabinet.ID, "", public)
	if err != nil {
		return nil, err
	}
	result := dto.CabinetResponse{ID: cabinet.ID, Name: cabinet.Name, Slug: cabinet.Slug, Tagline: cabinet.Tagline, Description: cabinet.Description, PeriodStart: cabinet.PeriodStart, PeriodEnd: cabinet.PeriodEnd, IsActive: cabinet.IsActive, IsPublished: cabinet.IsPublished, MetaTitle: cabinet.MetaTitle, MetaDescription: cabinet.MetaDescription}
	if cabinet.LogoMedia != nil {
		result.Logo = mediaToDTO(cabinet.LogoMedia)
	}
	if cabinet.HeroMedia != nil {
		result.Hero = mediaToDTO(cabinet.HeroMedia)
	}
	for _, unit := range units {
		if unit.ParentID == nil {
			result.Kemenkoan = append(result.Kemenkoan, unitToDTO(unit))
		}
	}
	return &result, nil
}

func (s *cabinetService) PublicUnit(slug string) (*dto.UnitResponse, error) {
	unit, err := s.repo.UnitBySlug(slug, true)
	if err != nil || unit == nil {
		return nil, err
	}
	result := unitToDTO(*unit)
	return &result, nil
}

func (s *cabinetService) PublicPrograms(slug string, page int) (repository.ListResult[dto.ProgramResponse], error) {
	unit, err := s.repo.UnitBySlug(slug, true)
	if err != nil || unit == nil {
		return repository.ListResult[dto.ProgramResponse]{}, err
	}
	return s.Programs(unit.ID, page, PublicProgramPageSize, true)
}

func (s *cabinetService) PublicProgram(unitSlug, programSlug string) (*dto.ProgramResponse, error) {
	program, err := s.repo.ProgramBySlug(unitSlug, programSlug, true)
	if err != nil || program == nil {
		return nil, err
	}
	result := programToDTO(*program)
	return &result, nil
}

func (s *cabinetService) Cabinets(page, perPage int) (repository.ListResult[dto.CabinetResponse], error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 50 {
		perPage = 20
	}
	values, err := s.repo.Cabinets(page, perPage)
	if err != nil {
		return repository.ListResult[dto.CabinetResponse]{}, err
	}
	result := repository.ListResult[dto.CabinetResponse]{Items: make([]dto.CabinetResponse, 0, len(values.Items)), Total: values.Total, Page: values.Page, PerPage: values.PerPage, TotalPages: values.TotalPages}
	for _, value := range values.Items {
		converted, convertErr := s.cabinetResponse(&value, false)
		if convertErr != nil {
			return result, convertErr
		}
		result.Items = append(result.Items, *converted)
	}
	return result, nil
}

func (s *cabinetService) CreateCabinet(req dto.CabinetRequest) (*dto.CabinetResponse, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("nama kabinet wajib diisi")
	}
	slug := strings.TrimSpace(req.Slug)
	if slug == "" {
		slug = utils.Slugify(req.Name)
	}
	value := &entities.CabinetTerm{Name: strings.TrimSpace(req.Name), Slug: slug, Tagline: req.Tagline, Description: req.Description, LogoMediaID: req.LogoMediaID, HeroMediaID: req.HeroMediaID, IsActive: req.IsActive, IsPublished: req.IsPublished, MetaTitle: req.MetaTitle, MetaDescription: req.MetaDescription}
	value.PeriodStart = parseDate(req.PeriodStart)
	value.PeriodEnd = parseDate(req.PeriodEnd)
	if value.IsActive {
		if err := s.repo.DB().Model(&entities.CabinetTerm{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
			return nil, err
		}
	}
	if err := s.repo.CreateCabinet(value); err != nil {
		return nil, err
	}
	return s.cabinetResponse(value, false)
}

func (s *cabinetService) UpdateCabinet(id uint64, req dto.CabinetRequest) (*dto.CabinetResponse, error) {
	var value entities.CabinetTerm
	if err := s.repo.DB().First(&value, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if strings.TrimSpace(req.Name) != "" {
		value.Name = strings.TrimSpace(req.Name)
	}
	if strings.TrimSpace(req.Slug) != "" {
		value.Slug = strings.TrimSpace(req.Slug)
	}
	value.Tagline = req.Tagline
	value.Description = req.Description
	value.LogoMediaID = req.LogoMediaID
	value.HeroMediaID = req.HeroMediaID
	value.IsActive = req.IsActive
	value.IsPublished = req.IsPublished
	value.MetaTitle = req.MetaTitle
	value.MetaDescription = req.MetaDescription
	value.PeriodStart = parseDate(req.PeriodStart)
	value.PeriodEnd = parseDate(req.PeriodEnd)
	if value.IsActive {
		if err := s.repo.DB().Model(&entities.CabinetTerm{}).Where("id <> ? AND is_active = ?", id, true).Update("is_active", false).Error; err != nil {
			return nil, err
		}
	}
	if err := s.repo.UpdateCabinet(&value); err != nil {
		return nil, err
	}
	return s.cabinetResponse(&value, false)
}

func (s *cabinetService) Units(cabinetID uint64, unitType string, public bool) ([]dto.UnitResponse, error) {
	values, err := s.repo.Units(cabinetID, unitType, public)
	if err != nil {
		return nil, err
	}
	result := make([]dto.UnitResponse, 0, len(values))
	for _, value := range values {
		result = append(result, unitToDTO(value))
	}
	return result, nil
}

func (s *cabinetService) CreateUnit(req dto.UnitRequest, actor *middlewares.Claims) (*dto.UnitResponse, error) {
	if err := validateUnit(req); err != nil {
		return nil, err
	}
	if req.UnitType == "KEMENKOAN" && actor.CanonicalRole() != "ADMIN_MEDINFO" {
		return nil, errors.New("hanya ADMIN_MEDINFO yang dapat mengelola kemenkoan")
	}
	if err := s.validateHierarchy(req.CabinetTermID, req.UnitType, req.ParentID, 0); err != nil {
		return nil, err
	}
	if req.ParentID != nil && !s.canManage(actor, *req.ParentID) {
		return nil, errors.New("akses unit ditolak")
	}
	value := &entities.Ministry{Code: strings.ToUpper(strings.TrimSpace(req.Code)), Name: strings.TrimSpace(req.Name), CabinetTermID: req.CabinetTermID, ParentID: req.ParentID, UnitType: req.UnitType, Slug: strings.TrimSpace(req.Slug), ShortName: req.ShortName, Description: req.Description, Vision: req.Vision, Mission: req.Mission, LogoMediaID: req.LogoMediaID, CoverMediaID: req.CoverMediaID, DisplayOrder: req.DisplayOrder, IsActive: req.IsActive, IsPublished: req.IsPublished}
	if value.Slug == "" {
		value.Slug = utils.Slugify(value.Name)
	}
	if err := s.repo.CreateUnit(value); err != nil {
		return nil, err
	}
	result := unitToDTO(*value)
	return &result, nil
}

func (s *cabinetService) UpdateUnit(id uint64, req dto.UnitRequest, actor *middlewares.Claims) (*dto.UnitResponse, error) {
	var value entities.Ministry
	if err := s.repo.DB().First(&value, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if !s.canManage(actor, id) {
		return nil, errors.New("akses unit ditolak")
	}
	if err := validateUnit(req); err != nil {
		return nil, err
	}
	if err := s.validateHierarchy(req.CabinetTermID, req.UnitType, req.ParentID, id); err != nil {
		return nil, err
	}
	if req.ParentID != nil && !s.canManage(actor, *req.ParentID) {
		return nil, errors.New("akses unit ditolak")
	}
	value.Code = strings.ToUpper(strings.TrimSpace(req.Code))
	value.Name = strings.TrimSpace(req.Name)
	value.CabinetTermID = req.CabinetTermID
	value.ParentID = req.ParentID
	value.UnitType = req.UnitType
	value.Slug = strings.TrimSpace(req.Slug)
	value.ShortName = req.ShortName
	value.Description = req.Description
	value.Vision = req.Vision
	value.Mission = req.Mission
	value.LogoMediaID = req.LogoMediaID
	value.CoverMediaID = req.CoverMediaID
	value.DisplayOrder = req.DisplayOrder
	value.IsActive = req.IsActive
	value.IsPublished = req.IsPublished
	if value.Slug == "" {
		value.Slug = utils.Slugify(value.Name)
	}
	if err := s.repo.UpdateUnit(&value); err != nil {
		return nil, err
	}
	result := unitToDTO(value)
	return &result, nil
}

func (s *cabinetService) Members(unitID uint64, public bool) ([]dto.MemberResponse, error) {
	values, err := s.repo.Members(unitID, public)
	if err != nil {
		return nil, err
	}
	result := make([]dto.MemberResponse, 0, len(values))
	for _, value := range values {
		result = append(result, memberToDTO(value))
	}
	return result, nil
}

func (s *cabinetService) CreateMember(req dto.MemberRequest, actor *middlewares.Claims) (*dto.MemberResponse, error) {
	if !s.canManage(actor, req.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	if err := validateMember(req); err != nil {
		return nil, err
	}
	value := &entities.OrganizationMember{MinistryID: req.MinistryID, Name: strings.TrimSpace(req.Name), Position: strings.TrimSpace(req.Position), PositionType: strings.ToUpper(strings.TrimSpace(req.PositionType)), Biography: req.Biography, Quote: req.Quote, PhotoMediaID: req.PhotoMediaID, DisplayOrder: req.DisplayOrder, IsLeader: req.IsLeader, IsActive: req.IsActive}
	if err := s.repo.CreateMember(value); err != nil {
		return nil, err
	}
	result := memberToDTO(*value)
	return &result, nil
}

func (s *cabinetService) UpdateMember(id uint64, req dto.MemberRequest, actor *middlewares.Claims) (*dto.MemberResponse, error) {
	var value entities.OrganizationMember
	if err := s.repo.DB().First(&value, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if !s.canManage(actor, value.MinistryID) || !s.canManage(actor, req.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	if err := validateMember(req); err != nil {
		return nil, err
	}
	value.MinistryID = req.MinistryID
	value.Name = strings.TrimSpace(req.Name)
	value.Position = strings.TrimSpace(req.Position)
	value.PositionType = strings.ToUpper(strings.TrimSpace(req.PositionType))
	value.Biography = req.Biography
	value.Quote = req.Quote
	value.PhotoMediaID = req.PhotoMediaID
	value.DisplayOrder = req.DisplayOrder
	value.IsLeader = req.IsLeader
	value.IsActive = req.IsActive
	if err := s.repo.UpdateMember(&value); err != nil {
		return nil, err
	}
	result := memberToDTO(value)
	return &result, nil
}

func (s *cabinetService) Programs(unitID uint64, page, perPage int, public bool) (repository.ListResult[dto.ProgramResponse], error) {
	if page < 1 {
		page = 1
	}
	values, err := s.repo.Programs(unitID, page, perPage, public)
	if err != nil {
		return repository.ListResult[dto.ProgramResponse]{}, err
	}
	result := repository.ListResult[dto.ProgramResponse]{Items: make([]dto.ProgramResponse, 0, len(values.Items)), Total: values.Total, Page: values.Page, PerPage: values.PerPage, TotalPages: values.TotalPages}
	for _, value := range values.Items {
		result.Items = append(result.Items, programToDTO(value))
	}
	return result, nil
}

func (s *cabinetService) Program(id uint64) (*dto.ProgramResponse, error) {
	var value entities.WorkProgram
	err := s.repo.DB().Preload("Ministry").Preload("CoverMedia").Preload("Milestones", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations.MediaAsset").First(&value, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	result := programToDTO(value)
	return &result, nil
}

func (s *cabinetService) CreateProgram(req dto.ProgramRequest, actor *middlewares.Claims) (*dto.ProgramResponse, error) {
	if !s.canManage(actor, req.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	if err := validateProgram(req); err != nil {
		return nil, err
	}
	value := programFromRequest(req)
	if err := s.repo.CreateProgram(value); err != nil {
		return nil, err
	}
	result := programToDTO(*value)
	return &result, nil
}

func (s *cabinetService) UpdateProgram(id uint64, req dto.ProgramRequest, actor *middlewares.Claims) (*dto.ProgramResponse, error) {
	var value entities.WorkProgram
	if err := s.repo.DB().First(&value, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if !s.canManage(actor, value.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	if err := validateProgram(req); err != nil {
		return nil, err
	}
	updated := programFromRequest(req)
	updated.ID = id
	updated.CreatedAt = value.CreatedAt
	updated.CreatedBy = value.CreatedBy
	updated.PublishedAt = value.PublishedAt
	if err := s.repo.UpdateProgram(updated); err != nil {
		return nil, err
	}
	result := programToDTO(*updated)
	return &result, nil
}

func (s *cabinetService) PublishProgram(id uint64, published bool, actor *middlewares.Claims) (*dto.ProgramResponse, error) {
	var value entities.WorkProgram
	if err := s.repo.DB().First(&value, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if !s.canManage(actor, value.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	value.IsPublished = published
	if published {
		now := time.Now()
		value.PublishedAt = &now
	} else {
		value.PublishedAt = nil
	}
	if err := s.repo.UpdateProgram(&value); err != nil {
		return nil, err
	}
	result := programToDTO(value)
	return &result, nil
}

func (s *cabinetService) CreateMilestone(req dto.MilestoneRequest, actor *middlewares.Claims) (*dto.MilestoneResponse, error) {
	var program entities.WorkProgram
	if err := s.repo.DB().First(&program, req.WorkProgramID).Error; err != nil {
		return nil, err
	}
	if !s.canManage(actor, program.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("judul milestone wajib diisi")
	}
	value := &entities.WorkProgramMilestone{WorkProgramID: req.WorkProgramID, Title: strings.TrimSpace(req.Title), Description: req.Description, StartDate: parseDate(req.StartDate), EndDate: parseDate(req.EndDate), Status: strings.ToUpper(strings.TrimSpace(req.Status)), DisplayOrder: req.DisplayOrder}
	if value.Status == "" {
		value.Status = "PLANNED"
	}
	if err := s.repo.CreateMilestone(value); err != nil {
		return nil, err
	}
	result := milestoneToDTO(*value)
	return result, nil
}

func (s *cabinetService) CreateDocumentation(programID uint64, req dto.DocumentationRequest, actor *middlewares.Claims) (*dto.DocumentationResponse, error) {
	var program entities.WorkProgram
	if err := s.repo.DB().First(&program, programID).Error; err != nil {
		return nil, err
	}
	if !s.canManage(actor, program.MinistryID) {
		return nil, errors.New("akses unit ditolak")
	}
	media, err := s.repo.MediaByID(req.MediaAssetID)
	if err != nil {
		return nil, err
	}
	if media == nil || media.Status != "ACTIVE" {
		return nil, errors.New("media asset tidak tersedia")
	}
	value := &entities.WorkProgramDocumentation{WorkProgramID: programID, MediaAssetID: req.MediaAssetID, Title: req.Title, Caption: req.Caption, DisplayOrder: req.DisplayOrder, IsCover: req.IsCover}
	if err := s.repo.CreateDocumentation(value); err != nil {
		return nil, err
	}
	if err := s.repo.DB().Preload("MediaAsset").First(value, value.ID).Error; err != nil {
		return nil, err
	}
	result := documentationToDTO(*value)
	return result, nil
}

func (s *cabinetService) ReorderDocumentation(programID uint64, ids []uint64, actor *middlewares.Claims) error {
	var program entities.WorkProgram
	if err := s.repo.DB().First(&program, programID).Error; err != nil {
		return err
	}
	if !s.canManage(actor, program.MinistryID) {
		return errors.New("akses unit ditolak")
	}
	var count int64
	if err := s.repo.DB().Model(&entities.WorkProgramDocumentation{}).Where("work_program_id = ?", programID).Count(&count).Error; err != nil {
		return err
	}
	seen := make(map[uint64]struct{}, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			return errors.New("urutan dokumentasi tidak valid")
		}
		seen[id] = struct{}{}
	}
	if int64(len(ids)) != count {
		return errors.New("semua dokumentasi harus disertakan saat mengurutkan")
	}
	return s.repo.ReorderDocumentations(programID, ids)
}

func (s *cabinetService) CreateMedia(req dto.MediaRequest, actor *middlewares.Claims) (*dto.MediaResponse, error) {
	if actor == nil || (actor.CanonicalRole() != "ADMIN" && actor.CanonicalRole() != "ADMIN_MEDINFO") {
		return nil, errors.New("akses upload ditolak")
	}
	if strings.TrimSpace(req.ImageKitFileID) == "" || strings.TrimSpace(req.URL) == "" || strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.AltText) == "" {
		return nil, errors.New("metadata media wajib lengkap")
	}
	value := &entities.MediaAsset{UploadedBy: &actor.UserID, ImageKitFileID: req.ImageKitFileID, FilePath: req.FilePath, URL: req.URL, ThumbnailURL: req.ThumbnailURL, Name: req.Name, AltText: req.AltText, Caption: req.Caption, MimeType: req.MimeType, SizeBytes: req.SizeBytes, Width: req.Width, Height: req.Height, Purpose: req.Purpose, Status: "ACTIVE"}
	if err := s.repo.CreateMedia(value); err != nil {
		return nil, err
	}
	result := mediaToDTO(value)
	return result, nil
}

func (s *cabinetService) canManage(actor *middlewares.Claims, unitID uint64) bool {
	if actor == nil {
		return false
	}
	if actor.CanonicalRole() == "ADMIN_MEDINFO" {
		return true
	}
	if actor.CanonicalRole() != "ADMIN" {
		return false
	}
	var count int64
	s.repo.DB().Table("user_organization_roles").Where("user_id = ? AND ministry_id = ? AND permission = ? AND is_active = ?", actor.UserID, unitID, "MANAGE_CONTENT", true).Count(&count)
	if count > 0 {
		return true
	}
	s.repo.DB().Table("users").Where("id = ? AND ministry_id = ?", actor.UserID, unitID).Count(&count)
	return count > 0
}

func (s *cabinetService) validateHierarchy(cabinetID *uint64, unitType string, parentID *uint64, currentID uint64) error {
	if cabinetID == nil || *cabinetID == 0 {
		return errors.New("kabinet wajib dipilih")
	}
	if unitType == "KEMENTERIAN" && parentID == nil {
		return errors.New("kementerian wajib memiliki kemenkoan")
	}
	if unitType == "KEMENKOAN" && parentID != nil {
		return errors.New("kemenkoan tidak dapat memiliki parent")
	}
	if parentID == nil {
		return nil
	}
	var parent entities.Ministry
	if err := s.repo.DB().First(&parent, *parentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("parent kemenkoan tidak ditemukan")
		}
		return err
	}
	if parent.UnitType != "KEMENKOAN" || parent.ParentID != nil {
		return errors.New("parent harus berupa kemenkoan")
	}
	if parent.CabinetTermID == nil || *parent.CabinetTermID != *cabinetID {
		return errors.New("parent harus berada dalam kabinet yang sama")
	}
	if currentID == 0 {
		return nil
	}
	visited := map[uint64]struct{}{currentID: {}}
	ancestorID := parent.ID
	for ancestorID != 0 {
		if _, exists := visited[ancestorID]; exists {
			return errors.New("hierarki organisasi membentuk cycle")
		}
		visited[ancestorID] = struct{}{}
		var ancestor entities.Ministry
		if err := s.repo.DB().Select("id, parent_id").First(&ancestor, ancestorID).Error; err != nil {
			return err
		}
		if ancestor.ParentID == nil {
			break
		}
		ancestorID = *ancestor.ParentID
	}
	return nil
}

func validateUnit(req dto.UnitRequest) error {
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Code) == "" {
		return errors.New("kode dan nama unit wajib diisi")
	}
	if req.UnitType != "KEMENKOAN" && req.UnitType != "KEMENTERIAN" {
		return errors.New("tipe unit tidak valid")
	}
	return nil
}
func validateMember(req dto.MemberRequest) error {
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Position) == "" {
		return errors.New("nama dan jabatan wajib diisi")
	}
	position := strings.ToUpper(strings.TrimSpace(req.PositionType))
	if position != "MINISTER" && position != "DIRECTOR_GENERAL" {
		return errors.New("posisi harus Menteri atau Dirjen")
	}
	return nil
}
func validateProgram(req dto.ProgramRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("nama program wajib diisi")
	}
	if !validStatuses[strings.ToUpper(strings.TrimSpace(req.LifecycleStatus))] {
		return errors.New("status program tidak valid")
	}
	return nil
}
func programFromRequest(req dto.ProgramRequest) *entities.WorkProgram {
	status := strings.ToUpper(strings.TrimSpace(req.LifecycleStatus))
	if status == "" {
		status = "DRAFT"
	}
	slug := strings.TrimSpace(req.Slug)
	if slug == "" {
		slug = utils.Slugify(req.Name)
	}
	return &entities.WorkProgram{MinistryID: req.MinistryID, Name: strings.TrimSpace(req.Name), Slug: slug, ShortDescription: req.ShortDescription, Description: req.Description, Objectives: req.Objectives, TargetAudience: req.TargetAudience, ExecutionMonth: req.ExecutionMonth, LifecycleStatus: status, CoverMediaID: req.CoverMediaID, DisplayOrder: req.DisplayOrder, IsFeatured: req.IsFeatured, IsPublished: req.IsPublished}
}
func parseDate(value *string) *time.Time {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil
	}
	parsed, err := time.Parse("2006-01-02", strings.TrimSpace(*value))
	if err != nil {
		return nil
	}
	return &parsed
}
func mediaToDTO(value *entities.MediaAsset) *dto.MediaResponse {
	if value == nil {
		return nil
	}
	return &dto.MediaResponse{ID: value.ID, FileID: value.ImageKitFileID, URL: value.URL, ThumbnailURL: value.ThumbnailURL, Name: value.Name, AltText: value.AltText, Caption: value.Caption, MimeType: value.MimeType, SizeBytes: value.SizeBytes, Width: value.Width, Height: value.Height, Purpose: value.Purpose, Status: value.Status}
}
func memberToDTO(value entities.OrganizationMember) dto.MemberResponse {
	return dto.MemberResponse{ID: value.ID, Name: value.Name, Position: value.Position, PositionType: value.PositionType, Biography: value.Biography, Quote: value.Quote, Photo: mediaToDTO(value.PhotoMedia), DisplayOrder: value.DisplayOrder, IsLeader: value.IsLeader}
}
func unitToDTO(value entities.Ministry) dto.UnitResponse {
	result := dto.UnitResponse{ID: value.ID, CabinetTermID: value.CabinetTermID, ParentID: value.ParentID, Code: value.Code, Name: value.Name, UnitType: value.UnitType, Slug: value.Slug, ShortName: value.ShortName, Description: value.Description, Vision: value.Vision, Mission: value.Mission, Logo: mediaToDTO(value.LogoMedia), Cover: mediaToDTO(value.CoverMedia), DisplayOrder: value.DisplayOrder, IsActive: value.IsActive, IsPublished: value.IsPublished, Members: make([]dto.MemberResponse, 0, len(value.Members)), Programs: make([]dto.ProgramResponse, 0, len(value.Programs)), Children: make([]dto.UnitResponse, 0, len(value.Children))}
	for _, member := range value.Members {
		result.Members = append(result.Members, memberToDTO(member))
	}
	for _, program := range value.Programs {
		response := programToDTO(program)
		if response.MinistryName == "" {
			response.MinistryName = value.Name
		}
		result.Programs = append(result.Programs, response)
	}
	for _, child := range value.Children {
		result.Children = append(result.Children, unitToDTO(child))
	}
	return result
}
func programToDTO(value entities.WorkProgram) dto.ProgramResponse {
	result := dto.ProgramResponse{ID: value.ID, MinistryID: value.MinistryID, Name: value.Name, Slug: value.Slug, ShortDescription: value.ShortDescription, Description: value.Description, Objectives: value.Objectives, TargetAudience: value.TargetAudience, StartDate: value.StartDate, EndDate: value.EndDate, ExecutionMonth: value.ExecutionMonth, Status: value.LifecycleStatus, Cover: mediaToDTO(value.CoverMedia), DisplayOrder: value.DisplayOrder, IsFeatured: value.IsFeatured, IsPublished: value.IsPublished, PublishedAt: value.PublishedAt, Milestones: make([]dto.MilestoneResponse, 0, len(value.Milestones)), Documentations: make([]dto.DocumentationResponse, 0, len(value.Documentations))}
	if value.Ministry != nil {
		result.MinistryName = value.Ministry.Name
	}
	for _, milestone := range value.Milestones {
		result.Milestones = append(result.Milestones, *milestoneToDTO(milestone))
	}
	for _, documentation := range value.Documentations {
		result.Documentations = append(result.Documentations, *documentationToDTO(documentation))
	}
	return result
}
func milestoneToDTO(value entities.WorkProgramMilestone) *dto.MilestoneResponse {
	return &dto.MilestoneResponse{ID: value.ID, Title: value.Title, Description: value.Description, StartDate: value.StartDate, EndDate: value.EndDate, Status: value.Status, DisplayOrder: value.DisplayOrder}
}
func documentationToDTO(value entities.WorkProgramDocumentation) *dto.DocumentationResponse {
	return &dto.DocumentationResponse{ID: value.ID, Media: mediaToDTO(value.MediaAsset), Title: value.Title, Caption: value.Caption, TakenAt: value.TakenAt, DisplayOrder: value.DisplayOrder, IsCover: value.IsCover}
}
