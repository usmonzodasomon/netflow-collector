package models

import "sync"

type TemplateFlowSet struct {
	FlowSetHeader
	Header TemplateHeader
	Fields TemplateFields
}

type TemplateHeader struct {
	TemplateID uint16
	FieldCount uint16
}

type TemplateField struct {
	FieldType   uint16
	FieldLength uint16
}

type TemplateFields []TemplateField

type TemplateSystem struct {
	Templates     map[uint16]TemplateFields
	Templateslock *sync.RWMutex
}

var Templates = CreateTemplateSystem()

func CreateTemplateSystem() *TemplateSystem {
	ts := &TemplateSystem{
		Templates:     make(map[uint16]TemplateFields),
		Templateslock: &sync.RWMutex{},
	}
	return ts
}

func (t *TemplateSystem) AddTemplate(templateID uint16, template TemplateFields) {
	t.Templateslock.Lock()
	t.Templates[templateID] = template
	t.Templateslock.Unlock()
}

func (t *TemplateSystem) GetTemplates(templateID uint16) TemplateFields {
	t.Templateslock.RLock()
	tmp := t.Templates[templateID]
	t.Templateslock.RUnlock()
	return tmp
}
