require 'goodreads'
def get_rating_of_book_by_title(title)
  client = Goodreads::Client.new(
    api_key: "YOUR_API_KEY",
    api_secret: "YOUR_API_SECRET"
  )
  book = client.book(title)
  book.rating
end
