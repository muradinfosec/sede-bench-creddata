#!/bin/bash

ab -n 20000 -c 100 -H "Authorization: Bearer oSfwni36CSMF45y0vas06AS1U8RkGc8tvJ51JQoN9LdnY8oZw42ebdg9gqwPnQeq" \
  http://localhost:3000/api/dashboards/db/dash1
