# go-shortener

## Functional Requirements:
- Service should be able to create shortened url/links against a long url
- Click to the short URL should redirect the user to the original long URL
- Shortened link should be as small as possible
- Users can create custom url with maximum character limit of 16
- Service should collect metrics like most clicked links
- Once a shortened link is generated it should stay in system for lifetime
  Non-Functional Requirements:

## Non-Functional Requirements:
- Service should be up and running all the time
- URL redirection should be fast and should not degrade at any point of time (Even during peak loads)
_ Service should expose REST API’s so that it can be integrated with third party applications