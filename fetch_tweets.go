package main
var apiKey = os.Getenv("TWITTER_BEARER_TOKEN")
type Tweet = struct{ Text string }
func fetchTweetsFromUser(user string) ([]Tweet, error) {
    url := "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=" + user + "&count=200"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+apiKey)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
    }
    var tweets []Tweet
    if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
        return nil, err
    }
    return tweets, nil
}
