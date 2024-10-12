package repository

type Repo struct{

}

type User struct{
	Username string
	Password string
}

func New() Repo {
	return Repo{}
}
