package academy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	name           string
	mockRepository func(t *testing.T) *MockRepository
}

func Test_GradeYear(t *testing.T) {
	tests := []test{
		{
			name: "No students",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				var year uint8 = 2
				mockRepository.EXPECT().List(year).Return(nil, nil)
				return mockRepository
			},
		},
		{
			name: "Two students",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				var year uint8 = 2
				mockRepository.EXPECT().List(year).Return([]string{"Jan Kowalski", "Adam Nowak"}, nil)
				mockRepository.EXPECT().Get("Jan Kowalski").Return(&Sophomore{
					name: "Jan Kowalski",
					grades: []int{
						1, 1, 1, 1, 1,
					},
					project: 1,
					attendance: []bool{
						true, true, true, true, true,
					},
				}, nil)
				mockRepository.EXPECT().Get("Adam Nowak").Return(&Sophomore{
					name: "Adam Nowak",
					grades: []int{
						5, 5, 3, 4, 3,
					},
					project: 4,
					attendance: []bool{
						true, true, true, true, true,
					},
				}, nil)
				mockRepository.EXPECT().Save("Jan Kowalski", year).Return(nil)
				year = 3
				mockRepository.EXPECT().Save("Adam Nowak", year).Return(nil)
				return mockRepository
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := tt.mockRepository(t)
			var year uint8 = 2
			err := GradeYear(mockRepository, year)
			assert.NoError(t, err)
		})
	}
}

func Test_GradeStudent(t *testing.T) {
	tests := []test{
		{
			name: "Student does not exist",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Jan Kowalski").Return(nil, ErrStudentNotFound)
				return mockRepository
			},
		},
		{
			name: "Final grade equals 1",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Jan Kowalski").Return(&Sophomore{
					name: "Jan Kowalski",
					grades: []int{
						1, 1, 1, 1, 1,
					},
					project: 1,
					attendance: []bool{
						true, true, true, true, true,
					},
				}, nil)
				var year uint8 = 2
				mockRepository.EXPECT().Save("Jan Kowalski", year).Return(nil)
				return mockRepository
			},
		},
		{
			name: "Final grade OK",
			mockRepository: func(t *testing.T) *MockRepository {
				mockRepository := NewMockRepository(t)
				mockRepository.EXPECT().Get("Jan Kowalski").Return(&Sophomore{
					name: "Jan Kowalski",
					grades: []int{
						5, 5, 3, 4, 3,
					},
					project: 4,
					attendance: []bool{
						true, true, true, true, true,
					},
				}, nil)
				var year uint8 = 3
				mockRepository.EXPECT().Save("Jan Kowalski", year).Return(nil)
				return mockRepository
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := tt.mockRepository(t)
			err := GradeStudent(mockRepository, "Jan Kowalski")
			assert.NoError(t, err)
		})
	}
}
