package domain

type Repository interface {
	SavePostData(postData PostData) error
}
