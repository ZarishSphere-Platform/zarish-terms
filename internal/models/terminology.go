package models

// CodeSystem represents a terminology code system (SNOMED CT, ICD-11, LOINC, etc.)
type CodeSystem struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	URL         string `gorm:"size:500;uniqueIndex" json:"url"`
	Version     string `gorm:"size:50" json:"version"`
	Status      string `gorm:"size:50;default:'active'" json:"status"` // draft, active, retired
	Publisher   string `gorm:"size:255" json:"publisher"`
	Description string `gorm:"type:text" json:"description"`
	Content     string `gorm:"size:50" json:"content"` // not-present, example, fragment, complete
}

// Concept represents a single code within a code system
type Concept struct {
	BaseModel

	CodeSystemID uint   `gorm:"index;not null" json:"code_system_id"`
	Code         string `gorm:"size:100;not null;index" json:"code"`
	Display      string `gorm:"size:500" json:"display"`
	Definition   string `gorm:"type:text" json:"definition"`
	ParentID     *uint  `gorm:"index" json:"parent_id,omitempty"`
}

// ValueSet represents a set of codes from one or more code systems
type ValueSet struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	URL         string `gorm:"size:500;uniqueIndex" json:"url"`
	Version     string `gorm:"size:50" json:"version"`
	Status      string `gorm:"size:50;default:'active'" json:"status"`
	Description string `gorm:"type:text" json:"description"`
}

// TableName overrides
func (CodeSystem) TableName() string {
	return "code_systems"
}

func (Concept) TableName() string {
	return "concepts"
}

func (ValueSet) TableName() string {
	return "value_sets"
}
