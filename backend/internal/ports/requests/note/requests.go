package requests

type CreateNoteRequest struct {
	Content string `json:"content" validate:"required"`
	TopicID uint   `json:"topic_id" validate:"required"`
}

type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

type GetNoteRequest struct {
	NoteID uint `json:"note_id"`
}

type UpdateNoteRequest struct {
	NoteID  uint   `json:"note_id"`
	Content string `json:"content"`
}

type ShareNoteRequest struct {
	NoteID uint   `json:"note_id"`
	Email  string `json:"email"`
}

type UnshareNoteRequest struct {
	NoteID uint   `json:"note_id"`
	Email  string `json:"email"`
}
