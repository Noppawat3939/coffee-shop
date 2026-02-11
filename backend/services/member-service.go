package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
)

type MemberService interface {
	CreateMember(req dto.CreateMemberRequest) (models.Member, error)
	FindMember(phone string) (models.Member, error)
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
