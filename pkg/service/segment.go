package service

import (
	app "rest-api-golang"
	"rest-api-golang/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment app.Segment) (int, error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) GetAll() ([]app.Segment, error) {
	return s.repo.GetAll()
}

func (s *SegmentService) GetById(segmentId int) (app.Segment, error) {
	return s.repo.GetById(segmentId)
}

func (s *SegmentService) Delete(segmentId int) error {
	return s.repo.Delete(segmentId)
}

func (s *SegmentService) Update(segmentId int, input app.UpdateSegmentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(segmentId, input)
}
