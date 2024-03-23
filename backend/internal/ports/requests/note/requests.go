package requests

type CreateNoteRequest struct {
	Content  string `json:"content" validate:"required"`
	TopicIDs []uint `json:"topic_ids" validate:"required"`
}

type DeleteNoteRequest struct {
	NoteID uint `json:"note_id"`
}

type GetNoteRequest struct {
	NoteID uint `json:"note_id"`
}

type UpdateNoteRequest struct {
	NoteID   uint   `json:"note_id"`
	Content  string `json:"content" validate:"required"`
	TopicIDs []uint `json:"topic_ids" validate:"required"`
}

type ShareNoteRequest struct {
	NoteID uint   `json:"note_id"`
	Email  string `json:"email" validate:"required"`
}

type UnshareNoteRequest struct {
	NoteID uint   `json:"note_id"`
	Email  string `json:"email" validate:"required"`
}
