package todoist

type CompletedTask struct {
	Content     string `json:"content"`
	MetaData    string `json:"meta_data"`
	UserID      string `json:"user_id"`
	TaskID      string `json:"task_id"`
	NoteCount   int    `json:"note_count"`
	ProjectID   string `json:"project_id"`
	SectionID   string `json:"section_id"`
	CompletedAt string `json:"completed_at"`
	ID          string `json:"id"`
}

type Project struct {
	Color      string `json:"color"`
	Collapsed  bool   `json:"collapsed"`
	ParentID   string `json:"parent_id"`
	IsDeleted  bool   `json:"is_deleted"`
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	ChildOrder int    `json:"child_order"`
	IsArchived bool   `json:"is_archived"`
	ViewStyle  string `json:"view_style"`
}

type Section struct {
	Collapsed    bool   `json:"collapsed"`
	AddedAt      string `json:"added_at"`
	ArchivedAt   string `json:"archived_at"`
	ID           string `json:"id"`
	IsArchived   bool   `json:"is_archived"`
	IsDeleted    bool   `json:"is_deleted"`
	Name         string `json:"name"`
	ProjectID    string `json:"project_id"`
	SectionOrder int    `json:"section_order"`
	SyncID       string `json:"sync_id"`
	UserID       string `json:"user_id"`
}

type CompletedTasksResponse struct {
	Items    []CompletedTask    `json:"items"`
	Projects map[string]Project `json:"projects"`
	Sections map[string]Section `json:"sections"`
}

type Task struct {
	Content         string    `json:"content"`                     // Required
	Description     string    `json:"description,omitempty"`       // Optional
	ProjectID       string    `json:"project_id,omitempty"`        // Optional
	Due             *DueDate  `json:"due,omitempty"`               // Optional
	Priority        int       `json:"priority,omitempty"`          // Optional
	ParentID        string    `json:"parent_id,omitempty"`         // Optional
	ChildOrder      int       `json:"child_order,omitempty"`       // Optional
	SectionID       string    `json:"section_id,omitempty"`        // Optional
	DayOrder        int       `json:"day_order,omitempty"`         // Optional
	Collapsed       bool      `json:"collapsed,omitempty"`         // Optional
	Labels          []string  `json:"labels,omitempty"`            // Optional
	AssignedByUID   string    `json:"assigned_by_uid,omitempty"`   // Optional
	ResponsibleUID  string    `json:"responsible_uid,omitempty"`   // Optional
	AutoReminder    bool      `json:"auto_reminder,omitempty"`     // Optional
	AutoParseLabels bool      `json:"auto_parse_labels,omitempty"` // Optional
	Duration        *Duration `json:"duration,omitempty"`          // Optional
}

type DueDate struct {
	Date        string `json:"date"`
	Timezone    string `json:"timezone,omitempty"`
	IsRecurring bool   `json:"is_recurring"`
}

type Duration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}
type CreateTaskResponse struct {
	SyncStatus    map[string]string `json:"sync_status"`
	TempIDMapping map[string]string `json:"temp_id_mapping"`
}

type Command struct {
	Type   string `json:"type"`
	TempID string `json:"temp_id"`
	UUID   string `json:"uuid"`
	Args   Task   `json:"args"`
}

type Request struct {
	Commands []Command `json:"commands"`
}
