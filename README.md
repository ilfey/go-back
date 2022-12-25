# Backend for development

It's a backend for frontend development or mobile application development. This program allows you to generate random text for your applications.

## How to use it?

You need to download the server and create a `.env` file or set flags when starting the server.

What should be the file?

```env
ADDRESS=0.0.0.0
PORT=8000
LOGLEVEL=debug
DATABASE_URL="postgresql://ilfey:QWEasd123@localhost:5432/go-back"
DATABASE_FILE=go-back.db
JWT_KEY=secret
JWT_LIFE_SPAN=24
```

What variables are needed?

* `ADDRESS` - The IP address that the server will listen on. Default: `0.0.0.0`
* `PORT` - The port that the server will listen on. Dafault: `8000`
* `LOGLEVEL` - Server logging level. Default: `info`
* `DATABASE_URL` - The URL of the PostgeSQL database where users will be stored. Default: `postgresql://ilfey:QWEasd123@localhost:5432/go-back`
* `DATABASE_FILE` - The SQLite database file. Used when it is not possible to connect to the PostgreSQL database. Default: `go-back.db`
* `JWT_KEY` - Secret key to generate JWT. Default: `secret`
* `JWT_LIFE_SPAN` - JWT life span in hours. Default: `24`

What flags to set?

You can also start the server with the -h flag to see what flags exist.
Flags will take precedence if you use flags and environment file at the same time.

* `-a` - The IP address that the server will listen on. Default: `0.0.0.0`
* `-p` - The port that the server will listen on. Dafault: `8000`
* `-ll` - Server logging level. Default: `info`
* `-du` - The URL of the PostgeSQL database where users will be stored. Default: `postgresql://ilfey:QWEasd123@localhost:5432/go-back`
* `-df` - The SQLite database file. Used when it is not possible to connect to the PostgreSQL database. Default: `go-back.db`
* `-jk` - Secret key to generate JWT. Default: `secret`
* `-jls` - JWT life span in hours. Default: `24`

## Routes

|               **Route**                  |             **Description**             |
|:-----------------------------------------|:----------------------------------------|
| `/ping`                                  | Check authorization                     |
| `/text/word?amount={count}`              | Returns the amount of random words      |
| `/text/paragraph?amount={count}`         | Returns the amount of random paragraphs |
| `/text/sentence?amount={count}`          | Returns the amount of random sentences  |
| `/img/{width}x{height}.png`              | Returns a png image                     |
| `/img/{width}x{height}.jpg`              | Returns a jpg image                     |
| `/img/{width}x{height}.gif`              | Returns a gif image                     |
| `/jwt/register`                          | Creates a new user in database          |
| `/jwt/login`                             | Authorizes the user                     |
| `/private/text/word?amount={count}`      | Returns the amount of random words      |
| `/private/text/paragraph?amount={count}` | Returns the amount of random paragraphs |
| `/private/text/sentence?amount={count}`  | Returns the amount of random sentences  |
| `/private/img/{width}x{height}.png`      | Returns a png image                     |
| `/private/img/{width}x{height}.jpg`      | Returns a jpg image                     |
| `/private/img/{width}x{height}.gif`      | Returns a gif image                     |

### Text endpoints

`http://127.0.0.1:8000/text/word?amount=10`

![word](https://imgur.com/juXNLSY.png)

`http://127.0.0.1:8000/text/sentence?amount=2`

![sentence](https://imgur.com/bbQa0ui.png)

`http://127.0.0.1:8000/text/paragraph?amount=1`

![paragraph](https://imgur.com/uolMMVx.png)

### Image endpoints

Query params

* `bg` - background color
* `fg` - foreground color
* `border` - border and diagonal width

When setting fg or bg options you can use alpha channel: `1234` or `12345678`

`http://127.0.0.1:8000/img/256x256.png?bg=121D32&fg=BF3284`

![256x256 image](https://imgur.com/xFDdOyE.png)

### JWT endpoints

`http://127.0.0.1:8000/jwt/register`

![register request](https://imgur.com/myrfpJ7.png)

`http://127.0.0.1:8000/jwt/login`

![login request](https://imgur.com/MPj569q.png)

### Ping endpoint

`http://localhost:8000/ping`

![ping](https://imgur.com/EGgeY0G.png)
