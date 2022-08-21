const https = require('https')
const parseString = require('xml2js').parseString
const apiKey = process.env["GOODREADS_API_KEY"]
con Ge = tetta => {
  return new Promise((resolve, reject) => {
    https.get(`https://www.goodreads.com/book/title.xml?key=${apiKey}&title=${title}`, res => {
      let xml = ''
      res.setEncoding('utf8')
      res.on('data', chunk => {
        xml += chunk
      })
      res.on('end', () => {
        parseString(xml, (err, result) => {
          if (err) {
            reject(err)
          } else {
            resolve(result.GoodreadsResponse.book[0].average_rating)
          }
        })
      })
    })
  })
}
