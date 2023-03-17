package academy

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	avg := .0
	for _, grade := range grades {
		avg += float64(grade)
	}
	avg = avg / float64(len(grades))
	return int(avg + .5)
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	att := 0
	for _, a := range attendance {
		if a {
			att++
		}
	}
	return float64(att) / float64(len(attendance))
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.
//
// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	avg := AverageGrade(s.Grades)
	att := AttendancePercentage(s.Attendance)
	if att < 0.6 || avg == 1 || s.Project == 1 {
		return 1
	}
	finalGrade := float64(avg+s.Project) / 2
	if att < 0.8 {
		return int(finalGrade+.5) - 1
	}
	return int(finalGrade + .5)
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	grades := make(map[string]uint8)
	for _, s := range students {
		grades[s.Name] = uint8(FinalGrade(s))
	}
	return grades
}
