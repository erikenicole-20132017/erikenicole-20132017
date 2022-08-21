const apiKey = process.env["GOODREADS_API_KEY"]
const getRating = (title: string) => {
  return fetch(`https://www.goodreads.com/book/title.xml?key=${apiKey}&title=${title}`)
    .then(res => res.text())
    .then(res => {
      const parser = new DOMParser();
      const xml = parser.parseFromString(res, "text/xml");
      const rating = xml.getElementsByTagName("average_rating")[0].textContent;
      return rating;
    })
    .catch(error => console.log(error));
}
