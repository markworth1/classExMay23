package repo
import (
	"fmt"
	"github.com/gSchool/classExMay23/users/pkg/repo"
	"testing"
)

func TestInsert(t *testing.T) {
	r := repo.Repo{}

	r.Insert("unit@test.com")
}

func TestGetAllData(t *testing.T) {
	var tests = []struct {
		insertData     string
		lenAfterInsert int
	}{
		{"test1", 1},
		{"user@user.com", 2},
		{"j", 3},
	}

	r := repo.Repo{}

	for _, tt := range tests {
		//insert each
		r.Insert(tt.insertData)

		repoData := r.GetAllData()
		fmt.Printf("%v",repoData)
		if tt.lenAfterInsert != len(repoData) {
			t.Error("Unexpected len from GetAllData()")
		}

		if !isStringInSlice(tt.insertData, repoData) {
			t.Error("Unexpected data from GetAllData()")
		}
	}
}

func isStringInSlice(check string, sl []string) bool {
	for _, s := range sl {
		if check == s {
			return true
		}
	}

	return false
}