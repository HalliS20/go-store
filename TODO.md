# Main stack

- DB + auth : pocketbase
- Backend : Go/Gin
- Frontend : Alpine/Htmx
- CSS Library : Daisy UI
- DNS/CDN/Proxy : Cloudflare

## Additional software

### Main software

- Payment : Paypal
- Mail : Send-Grid
- hosting : Fly.io

### Efficient Deployment

- form validation : go-playground/validator
- caching : go-cache , dragonfly, redis
- SEO : go-meta-tags

### Full Feature Deployment

- Analytics : Google Analytics or Plausible
- Request Limiting : ulule/limiter
- CSRF Security : unrolled/secure
- image optimization : go-imaging
- Multi language : go-i18n.

### Large scale deployment

- Search : Algolia

## Pocketbase

- Stored on Fly.io volumes
- handles auth, cart, users, pictures/products.
