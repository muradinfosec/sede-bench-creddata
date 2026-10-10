require(httr)

headers = c(
  `Accept` = 'application/json',
  `user-token` = '03b8fd4620a7c6520998ee99b5b9d029'
)

res <- httr::GET(url = 'http://localhost:8080/api/retail/books/list', httr::add_headers(.headers=headers))
