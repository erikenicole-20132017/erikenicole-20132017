package main
import ("log"; "net/http"; "net/url"; "os"; "strconv"; "github.com/beevik/etree")
var (
    apiKey = os.Getenv("GOODREADS_API_KEY")
)
// GetRating gets the average rating of a book using the Goodreads API.
func GetRating(title string) (float64, error) {
    u := url.URL{
        Scheme: "https",
        Host:   "www.goodreads.com",
        Path:   "/book/title.xml",
        RawQuery: url.Values{
            "key":   {apiKey},
            "title": {title},
    }.Encode(),
  }
  resp, err := http.Get(u.String())
  if err != nil {
      return 0, err
  }
  defer resp.Body.Close()
  doc := etree.NewDocument()
  if _, err := doc.ReadFrom(resp.Body); err != nil {
      return 0, err
  }
  ratings := doc.FindElements("//average_rating")
  if len(ratings) == 0 {
      return 0, nil
  }
  rating, err := strconv.ParseFloat(ratings[0].Text(), 64)
  if err != nil {
      return 0, err
  }
  return rating, nil
}
