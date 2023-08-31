package server3_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Bionic2113/avito/internal/models"
	"github.com/stretchr/testify/require"
)

func TestTrackingForDeletion(t *testing.T) {
	testTable := []struct {
		LastName     string
		FirstName    string
		Segment_name string
	}{
		{"Parker", "Piter", "DELETE_ME"},
		{"Wayne", "Bruce", "I_AM_BATMAN"},
	}
	for _, v := range testTable {
		user := &models.User{LastName: v.LastName, FirstName: v.FirstName}
		user, err := userRepo.Create(user)
		if err != nil {
			t.Fatal("Error in create user method: ", err)
		}

		user_segment := &models.UserSegment{User_id: user.Id, Segment_name: v.Segment_name}
		// метод должен сам создать сегмент
		user_segment, err = userSegmentRepo.Create(user_segment)
		if err != nil {
			t.Fatal("Error in create user_segment method: ", err)
		}
		require.Equal(t, true, user_segment.Active)
		user_segment.Duration = uint64(time.Now().Unix())
		_, err = userSegmentRepo.Update(user_segment)
		if err != nil {
			t.Fatal("Error in update user_segment method: ", err)
		}
		time.Sleep(63 * time.Second)
		u_arr, err := userSegmentRepo.FindAll()
		if err != nil {
			t.Log("Никого не нашел: ", err)
		}
		for i, v := range u_arr {
			fmt.Printf(
				"iter: %d, id=%d\nu_id=%d\nseg_n=%s\nct=%s\ncd=%s\ndur=%s\nactive=%v",
				i, v.Id, v.User_id, v.Segment_name,
				time.Unix(int64(v.CreationTime), 0).String(),
				time.Unix(int64(v.DeletionTime), 0).String(),
				time.Unix(int64(v.Duration), 0).String(),
				v.Active,
			)
		}
		user_segment, err = userSegmentRepo.FindById(int(user_segment.Id))
		if err != nil {
			t.Fatal("Error in FindById user_segment method: ", err)
		}
		require.Equal(t, false, user_segment.Active)

	}
}
