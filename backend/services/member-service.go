package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type MemberService interface {
	CreateMember(req dto.CreateMemberRequest) (models.Member, error)
	FindMember(phone string) (models.Member, error)
	FindAllMembers(filter models.MemberFilter, page, limt int) ([]models.MemberResponse, error)
}

type memberService struct {
	repo repository.MemberRepo
}

func NewMemberService(repo repository.MemberRepo) MemberService {
	return &memberService{repo}
}

func (s *memberService) CreateMember(req dto.CreateMemberRequest) (models.Member, error) {
	data, err := s.repo.Create(models.Member{PhoneNumber: req.PhoneNumber, FullName: req.FullName, Provider: "line"})
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *memberService) FindMember(phone string) (models.Member, error) {
	data, err := s.repo.FindOne(phone)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *memberService) FindAllMembers(filter models.MemberFilter, page, limit int) ([]models.MemberResponse, error) {
	var result []models.MemberResponse

	members, err := s.repo.FindAllIncluded(filter, page, limit)
	if err != nil {
		return result, err
	}

	for _, m := range members {
		result = append(result, models.MemberResponse{
			ID:          m.ID,
			FullName:    m.FullName,
			Provider:    m.Provider,
			PhoneNumber: maskPhone(m.PhoneNumber),
			CreatedAt:   m.CreatedAt,
			MemberPoint: m.MemberPoint,
		})
	}

	return result, err
}

func maskPhone(p string) string {
	if len(p) < 7 {
		return p
	}

	return p[:3] + "****" + p[len(p)-3:]
}
