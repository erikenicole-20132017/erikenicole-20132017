import os, requests, json, bs4
key = os.environ['GOODREADS_API_KEY']
def get_rating(title) :
    """Get the average rating of the book from GoodReads, and return a float."""
    url = 'https://www.goodreads.com/book/title.xml?key=' + key + '&title=' + title
    response = requests.get(url)
    soup = bs4.BeautifulSoup(response.text, 'xml')
    rating = soup.find('average_rating').text
    return float(rating)
