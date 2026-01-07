package dto

type Word struct {
	English string
	Russian string
}

type Homepage struct {
	Words []Word
	Year  int
}
