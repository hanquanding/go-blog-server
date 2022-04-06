/**
 * @author: hqd
 * @description: form common
 * @file: common
 * @date: 2021-02-11 19:24
 */

package form

type PaginationParam struct {
	PageNum  int
	PageSize int
}

type PaginationResult struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
