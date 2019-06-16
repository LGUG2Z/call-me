# Call Me
A ci-agnostic orchestrator for triggering tests on a shared environment pool across multiple pipelines.

# Background
At a certain scale, user acceptance and integration tests need to be run against live environments that are
either mirrors of a production environment or scaled down mirrors of a production environment. Scheduling
these tests can become complicated when they need to be triggered from multiple pipelines and repositories:

`Push to Repo A -> Deploy to Test Environment -> Run Tests`

`Push to Repo B -> Deploy to Test Environment -> Run Tests`

If the push to `Repo B` triggers a deployment and test run while the test run for `Repo A` is still in progress,
there is a high likelihood that this will produce side effects on the tests running for `Repo A` and produce
inaccurate results and false-failures.

If an unique test environment takes a short time to provision and destroy after every test, this would be
the preferred way to handle such scenarios, but if the environment provisioning takes longer than is acceptable
and is significantly detrimental to the development feedback loop, pre-provisioned environments become the
more attractive alternative.

Pre-provisioned environments can be managed at a small scale with manual triggers and a bit of discipline, but
this usually can't be maintained as the development team and number of repositories or projects grows.

# API Overview
`call-me` provides a RESTful API for managing slots and locks on pre-provisioned testing environments. Under the
hood, `call-me` uses [bbolt](https://github.com/etcd-io/bbolt) as a fast and reliable key/value store that is
optimised for read-intensive workloads As `bbolt` uses an exclusive write lock on the database, this nicely
handles the potential problem of multiple concurrent updates.

`call-me` exposes the following routes:
## `/maybe`
### `GET`
Checks if a slot on the desired environment is free.

If the slot is free:
```bash
curl -I -X GET 'http://call.me/maybe?environment=uat' -H 'X-API-KEY: E-MO-TION'

HTTP/1.1 200 OK
Date: Sun, 16 Jun 2019 19:29:16 GMT
Content-Length: 0
```

If the slot is locked:
```bash
curl -I -X GET 'http://call.me/maybe?environment=uat' -H 'X-API-KEY: E-MO-TION'

HTTP/1.1 403 Forbidden
Date: Sun, 16 Jun 2019 19:31:16 GMT
Content-Length: 0
```
### `POST`
Requests a slot on the desired environment.

If the slot is free and given to the requester:
```bash
curl -I -X POST 'http://call.me/maybe?environment=uat' -H 'X-API-KEY: E-MO-TION'

HTTP/1.1 201 Created
Date: Sun, 16 Jun 2019 19:29:44 GMT
Content-Length: 0
```

If the slot is locked at the time of the request:
```bash
curl -I -X POST 'http://call.me/maybe?environment=uat' -H 'X-API-KEY: E-MO-TION'

HTTP/1.1 403 Forbidden
Date: Sun, 16 Jun 2019 19:33:14 GMT
Content-Length: 0
```

### `DELETE`
Releases a slot on the specified environment.

Get the latest manifest for a specific hostname.
```bash
curl -I -X POST 'http://call.me/maybe?environment=uat' -H 'X-API-KEY: E-MO-TION'

HTTP/1.1 204 No Content
Date: Sun, 16 Jun 2019 19:35:39 GMT
```

# Server
The server takes two environment variables as configuration:
* `PORT`
* `API_KEY`

```bash
# docker
docker build -t call-me .
docker run -e PORT=8000 -e API_KEY=crj -p 8000:8000 call-me

# binary
GO111MODULE=on go install ./cmd/call-me-server
PORT=8000 API_KEY=crj call-me-server
```

# Client
The client requires two environment variables:
* `HOST`
* `API_KEY`

```
GO111MODULE=on go install ./cmd/call-me-client

export HOST=call-me-server:8000
export API_KEY=crj
```

Request a free slot:
```
# request a free slot
call-me-client request uat

acquired slot for environment: uat
waited 57.134996ms
```

Wait until a request is free and then lock the slot:
```
call-me-client request uat

waiting for free slot on environment: uat
waiting for free slot on environment: uat
acquired slot for environment: uat
waited 10.089027168s
```

Release a slot when finished:

```
call-me-client release uat

slot has been freed for environment: uat
```
