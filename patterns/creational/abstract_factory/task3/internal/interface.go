package internal

const (
	VideoLessonType       = "video"
	TextLessonType        = "text"
	InteractiveLessonType = "interactive"
)

type Course struct {
	chapters []Chapter
}

func (c *Course) AddChapter(chapter Chapter) {
	c.chapters = append(c.chapters, chapter)
}

func (c *Course) GetChapters() []Chapter {
	return c.chapters
}

type Chapter struct {
	sections []Section
}

func (c *Chapter) AddSection(section Section) {
	c.sections = append(c.sections, section)
}

func (c *Chapter) GetSections() []Section {
	return c.sections
}

type Section struct {
	lessons []Lesson
}

func (c *Section) AddLesson(lesson Lesson) {
	c.lessons = append(c.lessons, lesson)
}

func (c *Section) GetLessons() []Lesson {
	return c.lessons
}

type Lesson interface {
	Display()
}

type CourseFactory interface {
	CreateCourse() Course
	CreateChapter() Chapter
	CreateSection() Section
	CreateLesson(lessonType string) Lesson
}

type OnlineCourseFactory struct{}

func (c *OnlineCourseFactory) CreateCourse() Course {
	return Course{}
}

func (c *OnlineCourseFactory) CreateChapter() Chapter {
	return Chapter{}
}

func (c *OnlineCourseFactory) CreateSection() Section {
	return Section{}
}

func (c *OnlineCourseFactory) CreateLesson(lessonType string) Lesson {
	switch lessonType {
	case VideoLessonType:
		return &VideoLesson{}
	case TextLessonType:
		return &TextLesson{}
	case InteractiveLessonType:
		return &InteractiveLesson{}
	default:
		return nil
	}
}

type VideoLesson struct{}

func (l *VideoLesson) Display() {
	println("Это видеоурок")
}

type TextLesson struct{}

func (l *TextLesson) Display() {
	println("Это текстовый урок")
}

type InteractiveLesson struct{}

func (l *InteractiveLesson) Display() {
	println("Это интерактивный урок")
}
