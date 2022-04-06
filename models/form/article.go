/**
 * @author: hqd
 * @description: article form
 * @file: article
 * @date: 2021-02-12 21:01
 */

package form

type Article struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description" `
	Content     string `json:"content" `
	CoverImgURL string `json:"cover_img_url"`
	Sort        int    `json:"sort"`
	Status      uint   `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AddArticle struct {
	Title       string `json:"title" valid:"Required; MaxSize(150)"`
	Description string `json:"description" valid:"MaxSize(1000)"`
	Content     string `json:"content" valid:"Required"`
	CoverImgURL string `json:"cover_img_url"`
	Sort        int    `json:"sort"`
	Status      uint   `json:"status"`
}

type UpdateArticle struct {
	ID          uint   `json:"id" valid:"Required"`
	Title       string `json:"title" valid:"Required; MaxSize(150)"`
	Description string `json:"description" valid:"MaxSize(1000)"`
	Content     string `json:"content" valid:"Required"`
	CoverImgURL string `json:"cover_img_url"`
	Sort        int    `json:"sort"`
	Status      uint   `json:"status"`
}

type DeleteArticle struct {
	ID string `json:"ids" valid:"Required"`
}

type ArticleQueryParam struct {
	Title       string
	Description string
	Status      string
}

type ArticleQueryOptions struct {
	PageParam *PaginationParam
}

type Articles []*Article

func (a Articles) ToFormArticles() []*Article {
	list := make([]*Article, len(a))
	for i, v := range a {
		list[i] = v.ToArticle()
	}
	return list
}

func (a Article) ToArticle() *Article {
	return &Article{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		Content:     a.Content,
		CoverImgURL: a.CoverImgURL,
		Sort:        a.Sort,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
