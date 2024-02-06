# `mentos1386/golang-rest-example`

### Development

Tools required:
 * [Just](https://github.com/casey/just)
 * [watchexec](https://github.com/watchexec/watchexec) for `-watch` commands
 * Docker and Docker Compose

```sh
# Run development version
just run
# Run "production" version
just deploy

# Api is available at http://localhost:1234
# Swagger is available at http://localhost:1235

curl -v localhost:1234/healthz

# Run tests
just test
```
