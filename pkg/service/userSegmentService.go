package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Bionic2113/avito/internal/models"
	"github.com/Bionic2113/avito/internal/repository"
)

// сначала найди пользователя и чекни есть ли какие-то
// нужные записи для удаления, чтобы не получить ошибки,
// если нет таких данных
// также нужно проверить нет ли уже каких-то, что надо
// добавить
// можно даже отчет в конце намутить какие лишние были
// в списках на удаление и добавление
type Service struct {
	userRepo        repository.UserRepository
	userSegmentRepo repository.UserSegmentRepository
}

type elem struct {
	name   string
	status bool
}

type Record struct {
	Id           uint64
	Segment_name string
	Operation    Operations
	Time         time.Time
}

type Operations string

const (
	CREATE Operations = "CREATE"
	DELETE Operations = "DELETE"
)

func (service *Service) New(ur repository.UserRepository, usr repository.UserSegmentRepository) *Service {
	return &Service{userRepo: ur, userSegmentRepo: usr}
}

// Вопросы:
// 1) Если сегмент удален, то надо ли удалять у пользователя сегмент?
// Под удален имею ввиду, что active = false
// Решил, что в этом нет необходимости, тк может даже удаление
// сегмента вообще не нужно, ведь в задании об этом не упоминалось.
func (service *Service) UpdateUserSegments(id int, add []string, del []string, dur []uint64) ([]models.UserSegment, error) {
	result, deleted, err := service.checkData(id, add, del)
	if err != nil {
		return nil, err
	}
	create_chan := make(chan []models.UserSegment)
	go func() {
		create_chan <- service.create(id, add, dur)
	}()
	result = service.delete(result, del)
	new_data := <-create_chan
	result = append(result, deleted...)
	result = append(result, new_data...)
	return result, nil
}

// для создания же по сути тебе надо только айди чела и название сегмента
// поэтому вроде проблем в данных нет
func (service *Service) create(id int, add []string, dur []uint64) []models.UserSegment {
	for len(add) > len(dur){
	  dur = append(dur, 0)
	}
	create_array := make([]models.UserSegment, len(add))
	for i, v := range add {
		userSegment := &models.UserSegment{
			User_id:      uint64(id),
			Segment_name: v,
			Duration: 0,
		}
		
		userSegment, err := service.userSegmentRepo.Create(userSegment)
		if err != nil {
			userSegment = &models.UserSegment{Segment_name: "Error in the creation method"}
		}
		if dur[i] != 0{
		    userSegment.Duration = uint64(time.Now().AddDate(0, 0, int(dur[i])).Unix())
			userS, err := service.userSegmentRepo.Update(userSegment)
			if err == nil {userSegment = userS}
		}

		create_array[i] = *userSegment
	}
	return create_array
}

// но тут тебе надо список тех, у кого надо бы изменить данные
// получается что сюда надо передать массив из service.userSegmentRepo.FindAllById(id, true)
// затем ты его тут отредачишь, потом вернешь и к нему добавишь массив из create
func (service *Service) delete(data []models.UserSegment, del []string) []models.UserSegment {
	for i, userSegment := range data {
		for _, str := range del {
			if userSegment.Segment_name == str {
				userSegment.Active = false
				userSegment.DeletionTime = uint64(time.Now().Unix())
				userSegment, err := service.userSegmentRepo.Update(&userSegment)
				if err != nil {
					userSegment = &models.UserSegment{Segment_name: "Error in the deletion method"}
				}
				data[i] = *userSegment
			}
		}
	}
	return data
}

func (service *Service) checkData(id int, add []string, del []string) ([]models.UserSegment, []models.UserSegment, error) {
	if user, err := service.userRepo.FindById(id); err != nil || !user.Active {
		return nil, nil, fmt.Errorf("user is not found or have status deleted. user status is %v", user.Active)
	}
	active, err := service.userSegmentRepo.FindAllById(id, true)
	if err != nil {
		return nil, nil, errors.New("error when searching for active")
	}
	activeChan := make(chan elem)
	go func() {
		var e elem
		for _, s := range add {
			e = contains(s, active)
			if e.status {
				break
			}
		}
		activeChan <- e
	}()
	deleted, err := service.userSegmentRepo.FindAllById(id, false)
	if err != nil {
		return nil, nil, errors.New("error when searching for deleted")
	}
	for _, s := range del {
		if deleted_elem := contains(s, deleted); deleted_elem.status {
			return nil, nil, fmt.Errorf("user already have %s in deleted array", deleted_elem.name)
		}
	}
	active_elem := <-activeChan
	if active_elem.status {
		return nil, nil, fmt.Errorf("user already have %s in active array", active_elem.name)
	}
	return active, deleted, nil
}

func contains(str string, arr []models.UserSegment) elem {
	for _, s := range arr {
		if str == s.Segment_name {
			return elem{s.Segment_name, true}
		}
	}
	return elem{status: false}
}

func (service *Service) filterByMonthYear(year int, month int) []Record {
	data, err := service.userSegmentRepo.FindAll()
	if err != nil {
		log.Println("filterByMonthYear() Error in findAll", err)
		return nil
	}
	filtered := make([]Record, 0)
	for _, userSegment := range data {
		t_create := time.Unix(int64(userSegment.CreationTime), 0)
		t_delete := time.Unix(int64(userSegment.DeletionTime), 0)
		if t_create.Year() == year && t_create.Month() == time.Month(month) {
			filtered = append(filtered, Record{
				Id:           userSegment.Id,
				Segment_name: userSegment.Segment_name,
				Operation:    CREATE,
				Time:         t_create,
			})
		}
		if t_delete.Year() == year && t_delete.Month() == time.Month(month) {
			filtered = append(filtered, Record{
				Id:           userSegment.Id,
				Segment_name: userSegment.Segment_name,
				Operation:    DELETE,
				Time:         t_delete,
			})
		}

	}
	if len(filtered) == 0 {
		return nil
	}
	return filtered
}

// Вопрос: если нет элементов, то возвращать
// ссылку на пустой пустой файл или же ошибку
// Выбрал строку с ошибкой
func (service *Service) CreateCSV(year, month int) string {
	data := service.filterByMonthYear(year, month)
	if data == nil {
		return "Not found elements"
	}
	filename := fmt.Sprintf("%d-%d_%s.csv", year, month, time.Now().Local().Format("2006-01-02_15:04:05"))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "Not found elements"
	}
	defer file.Close()
	for _, record := range data {
		file.WriteString(fmt.Sprintf(
			"id: %d;segment: %s;operation: %s;date: %s\n",
			record.Id,
			record.Segment_name,
			record.Operation,
			record.Time.Local().String(),
		))
	}

	return filename
}

func (service *Service) AddPercent()([]models.UserSegment, error){
  return nil,nil
}

func (service *Service) CleaningOld() {
	for {
		service.userSegmentRepo.Cleaning()
		time.Sleep(1 * time.Minute)
	}
}
