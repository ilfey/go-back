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
JWT_KEY=secret
JWT_LIFE_SPAN=24
```

What variables are needed?

* `ADDRESS` - The IP address that the server will listen on. Default: `0.0.0.0`
* `PORT` - The port that the server will listen on. Dafault: `8000`
* `LOGLEVEL` - Server logging level. Default: `info`
* `DATABASE_URL` - The URL of the PostgeSQL database where users will be stored. Default: `postgresql://ilfey:QWEasd123@localhost:5432/go-back`
* `JWT_KEY` - Secret key to generate JWT. Default: `secret`
* `JWT_LIFE_SPAN` - JWT life span in hours. Default: `24`

What flags to set?

You can also start the server with the -h flag to see what flags exist.
Flags will take precedence if you use flags and environment file at the same time.

* `-a` - The IP address that the server will listen on. Default: `0.0.0.0`
* `-p` - The port that the server will listen on. Dafault: `8000`
* `-ll` - Server logging level. Default: `info`
* `-du` - The URL of the PostgeSQL database where users will be stored. Default: `postgresql://ilfey:QWEasd123@localhost:5432/go-back`
* `-jk` - Secret key to generate JWT. Default: `secret`
* `-jls` - JWT life span in hours. Default: `24`

## Routes

|               **Route**              |         **Description**         |
|:------------------------------------:|:-------------------------------:|
| `/text/word`                         | Returns a random word           |
| `/text/sentence`                     | Returns a random sentence       |
| `/text/paragraph`                    | Returns a random paragraph      |
| `/text/words/{count}`                | Returns count random words      |
| `/text/paragraphs/{count}`           | Returns count random paragraphs |
| `/text/sentences/{count}`            | Returns count random sentences  |
| `/img/{width}x{height}.png`          | Returns a png image             |
| `/img/{width}x{height}.jpg`          | Returns a jpg image             |
| `/img/{width}x{height}.jpeg`         | Returns a jpeg image            |
| `/img/{width}x{height}.gif`          | Returns a gif image             |
| `/jwt/register`                      | Creates a new user in database  |
| `/jwt/login`                         | Authorizes the user             |
| `/private/text/word`                 | Returns a random word           |
| `/private/text/sentence`             | Returns a random sentence       |
| `/private/text/paragraph`            | Returns a random paragraph      |
| `/private/text/words/{count}`        | Returns count random words      |
| `/private/text/paragraphs/{count}`   | Returns count random paragraphs |
| `/private/text/sentences/{count}`    | Returns count random sentences  |
| `/private/img/{width}x{height}.png`  | Returns a png image             |
| `/private/img/{width}x{height}.jpg`  | Returns a jpg image             |
| `/private/img/{width}x{height}.jpeg` | Returns a jpeg image            |
| `/private/img/{width}x{height}.gif`  | Returns a gif image             |

### Text routes

`http://127.0.0.1:8000/text/word`

![word](https://imgur.com/iAHbQMA.png)

`http://127.0.0.1:8000/text/sentence`

![sentence](https://imgur.com/g4UyvKL.png)

`http://127.0.0.1:8000/text/paragraph`

![paragraph](https://imgur.com/xQWqyJo.png)

### Image routes

Query params

* `bg` - background color
* `fg` - foreground color
* `border` - border and diagonal width

When setting fg or bg options you can use alpha channel: `1234` or `12345678`

`http://127.0.0.1:8000/img/256x256.png`

![256x256 image](https://imgur.com/j97nzA5.png)

`http://127.0.0.1:8000/img/256x256.png?bg=00000000&fg=215&border=10`

![256x256 image with transparent bg](https://imgur.com/8qi3U6z.png)

### JWT routes

`http://127.0.0.1:8000/jwt/register`

`http://127.0.0.1:8000/jwt/login`

