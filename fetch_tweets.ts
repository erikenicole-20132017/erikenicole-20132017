const token = process.env["TWITTER_BEARER_TOKEN"]
const fetchTweetsFromUser = (userName: string) => {
  const url = `https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=${userName}&count=20`
  return fetch(url, {
    headers: {
      "Authorization": `Bearer ${token}`
    }
  })
    .then(res => res.json())
    .then(tweets => tweets.map(tweet => ({
      id: tweet.id,
      text: tweet.text,
      created_at: tweet.created_at,
      user: {
        id: tweet.user.id,
        name: tweet.user.name,
        screen_name: tweet.user.screen_name,
        profile_image_url: tweet.user.profile_image_url
      }
    })))
}
