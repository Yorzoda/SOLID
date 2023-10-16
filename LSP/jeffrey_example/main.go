package main

import "gorm.io/gorm"

type LessonRepository interface {
	// getAll should only return an array of something
	getAll() any
}

type FileLessonRepository struct {
}

type DbLessonRepository struct {
}

func (f FileLessonRepository) getAll() any {
	var lessonRepository []FileLessonRepository
	return lessonRepository
}

func (d DbLessonRepository) getAll() any {
	var db gorm.DB
	var Lessons []DbLessonRepository
	if err := db.Table("lessons").Find(&Lessons); err != nil {
		return nil
	}

	return Lessons
}

func main() {
	var repo DbLessonRepository
	repo.getAll()
}
