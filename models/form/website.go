/**
 * @author: hqd
 * @description: website form
 * @file: website
 * @date: 2021-02-07 20:52
 */

package form

type AddWebsiteReq struct {
	WebName        string `json:"web_name" valid:"Required; MaxSize(100)"`
	WebURL         string `json:"web_url"  valid:"Required; MaxSize(200)"`
	WebDescription string `json:"web_desc" valid:"MaxSize(1000)"`
	WebSort        int    `json:"web_sort"`
	WebStatus      uint   `json:"web_status"`
	IsDeleted      uint   `json:"is_deleted"`
}

type UpdateWebsiteReq struct {
	WebID          uint   `json:"web_id"  valid:"Required"`
	WebName        string `json:"web_name" valid:"MaxSize(100)"`
	WebURL         string `json:"web_url"  valid:"MaxSize(200)"`
	WebDescription string `json:"web_desc" valid:"MaxSize(1000)"`
	WebSortNum     int    `json:"web_sort"`
	WebStatus      uint   `json:"web_status"`
}

type Website struct {
	WebID          uint   `json:"web_id"`
	WebName        string `json:"web_name"`
	WebURL         string `json:"web_url" `
	WebDescription string `json:"web_desc"`
	WebSort        int    `json:"web_sort"`
	WebStatus      uint   `json:"web_status"`
	IsDeleted      uint   `json:"is_deleted"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type WebsiteQueryParam struct {
	WebName   string
	WebURL    string
	WebStatus string
}

type WebsiteQueryOptions struct {
	PageParam *PaginationParam
}

type Websites []*Website

func (w Websites) ToFormWebsites() []*Website {
	list := make([]*Website, len(w))
	for i, v := range w {
		list[i] = v.ToWebsite()
	}
	return list
}

func (w *Website) ToWebsite() *Website {
	return &Website{
		WebID:          w.WebID,
		WebName:        w.WebName,
		WebURL:         w.WebURL,
		WebDescription: w.WebDescription,
		WebSort:        w.WebSort,
		WebStatus:      w.WebStatus,
		IsDeleted:      w.IsDeleted,
		CreatedAt:      w.CreatedAt,
		UpdatedAt:      w.UpdatedAt,
	}
}
