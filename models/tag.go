package models

type Tag struct {
	ID   string
	Name string
}

// Tags is a collection of Tag elements
type Tags []*Tag

func (z Tags) GetByID(id string) *Tag {
	for _, tag := range z {
		if tag.ID == id {
			return tag
		}
	}
	return nil
}

func (z Tags) GetByName(name string) *Tag {
	for _, tag := range z {
		if tag.Name == name {
			return tag
		}
	}
	return nil
}
