package service

import (
	app "rest-api-golang"
	"rest-api-golang/pkg/repository"
	//"errors"
	//"github.com/lib/pq"
	//"slices"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

//func checkSegmentsExists(segments pq.Int64Array) error {
//	segmentsFromDb, err := SegmentService{repo: NewSegmentService(Segment)}.repo.GetAll()
//	if err != nil {
//		return err
//	}
//	var segmentsIds []int64
//	for _, segment := range segmentsFromDb {
//		segmentsIds = append(segmentsIds, int64(segment.Id))
//	}
//
//	for _, id := range segments {
//		if slices.Contains(segmentsIds, id) == false {
//			err = errors.New("segment id not exists")
//			return err
//		}
//	}
//	return err
//}

func (s *UserService) Create(user app.CreateUserInput) (int, error) {
	//if err := checkSegmentsExists(*user.Segments); err != nil {
	//	return -1, err
	//}
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]app.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId int) (app.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId int, input app.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, input)
}

func (s *UserService) AddSegments(userId int, input app.AddUserSegmentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.AddSegments(userId, input)
}

func (s *UserService) GetSegments(userId int) ([]app.Segment, error) {
	return s.repo.GetSegments(userId)
}

func (s *UserService) DeleteSegments(userId int) error {
	return s.repo.DeleteSegments(userId)
}
