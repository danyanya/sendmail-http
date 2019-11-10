# sendmail http

  

Simple http server with Email sending. Now support Sendmail and Mutt as backend.

  

## Build

 

Binary can be built by *make*:

```bash
make build
# this creates binary in *bin* directory
ls bin/
# sendmail-http
```

To make Docker image:

```bash
make docker
# or building by your own image
docker build -t your-own/sendmail-http .
```

## Running

```bash
# Running directly with Docker
docker run -d -t --name sendmail \
    -p 1001:1001 \
    danyanya/sendmail-http:latest
# or by docker-compose
docker-compose up -d
```

## Usage

For send email use **/api/sendmutt** endpoint:

```bash
# Get params are from, to, subject and body:  
curl -sSl localhost:1001/sendmutt?from=bill.gates@microsoft.com&to=steve.jobs@apple.com&subject=[job%20offer]&body=Confirmation,$%20Steve!

# Also it accepts file, but it needed to be in container!
#!! &file=/tmp/some.xls
```  

## Copyright

Repo created by Daniil Sliusar at 2019.
