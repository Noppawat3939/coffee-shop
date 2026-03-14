package services

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/repository"
)

type MemberService interface {
	CreateMember(req dto.CreateMemberRequest) (model.Member, error)
	FindMember(phone string) (model.Member, error)
	FindAllMembers(filter model.MemberFilter, page, limt int) ([]model.MemberResponse, error)
}

type memberService struct {
	repo repository.MemberRepo
}

func NewMemberService(repo repository.MemberRepo) MemberService {
	return &memberService{repo}
}

func (s *memberService) CreateMember(req dto.CreateMemberRequest) (model.Member, error) {
	data, err := s.repo.Create(model.Member{PhoneNumber: req.PhoneNumber, FullName: req.FullName, Provider: "line"})
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *memberService) FindMember(phone string) (model.Member, error) {
	data, err := s.repo.FindOne(phone)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *memberService) FindAllMembers(filter model.MemberFilter, page, limit int) ([]model.MemberResponse, error) {
	var result []model.MemberResponse

	members, err := s.repo.FindAllIncluded(filter, page, limit)
	if err != nil {
		return result, err
	}

	for _, m := range members {
		result = append(result, model.MemberResponse{
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
