/**
 * @author: hqd
 * @description: tag form
 * @file: tag
 * @date: 2021-02-12 18:42
 */

package form

type TagQueryParam struct {
	TagName   string
	TagStatus string
}

type TagQueryOptions struct {
	PageParam *PaginationParam
}

type Tag struct {
	ID        uint   `json:"id"`
	TagName   string `json:"tag_name"`
	TagStatus uint   `json:"tag_status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddTag struct {
	TagName   string `json:"tag_name" valid:"Required; MaxSize(100)"`
	TagStatus uint   `json:"tag_status"`
}

type UpdateTag struct {
	ID        uint   `json:"id" valid:"Required"`
	TagName   string `json:"tag_name" valid:"Required; MaxSize(100)"`
	TagStatus uint   `json:"tag_status"`
}

type DeleteTag struct {
	ID string `json:"ids" valid:"Required"`
}

type Tags []*Tag

func (t Tags) ToFormTags() []*Tag {
	list := make([]*Tag, len(t))
	for i, v := range t {
		list[i] = v.ToTag()
	}
	return list
}

func (t Tag) ToTag() *Tag {
	return &Tag{
		ID:        t.ID,
		TagName:   t.TagName,
		TagStatus: t.TagStatus,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
