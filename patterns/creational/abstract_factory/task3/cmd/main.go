package main

import "task3/internal"

func main() {
	factory := internal.OnlineCourseFactory{}

	// Создание курса
	course := factory.CreateCourse()

	// Создание главы
	chapter := factory.CreateChapter()

	// Создание разделов
	section1 := factory.CreateSection()
	section2 := factory.CreateSection()

	// Создание уроков
	lesson1 := factory.CreateLesson(internal.VideoLessonType)
	lesson2 := factory.CreateLesson(internal.TextLessonType)
	lesson3 := factory.CreateLesson(internal.InteractiveLessonType)

	section1.AddLesson(lesson1)
	section1.AddLesson(lesson2)

	section2.AddLesson(lesson1)
	section2.AddLesson(lesson2)
	section2.AddLesson(lesson3)

	chapter.AddSection(section1)
	chapter.AddSection(section2)

	course.AddChapter(chapter)

	chapters := course.GetChapters()

	for _, ch := range chapters {
		for _, sec := range ch.GetSections() {
			for _, les := range sec.GetLessons() {
				les.Display()
			}
		}
	}
}
