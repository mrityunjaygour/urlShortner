package database

//Here we create an interface which has methods for creating
//a short code for url and also for getting actual url with the
//provided short code
type UrlShortnerDbService interface {
	Save(url string) (string, error)
	Get(shortUrl string) (string, error)
}
