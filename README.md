# Backend for development

It's a backend for frontend development or mobile application development. This program allows you to generate random text for your applications.

## Routes

|           **Route**          |         **Description**         |
|:----------------------------:|:-------------------------------:|
| `/text/word`                 | Returns a random word           |
| `/text/sentence`             | Returns a random sentence       |
| `/text/paragraph`            | Returns a random paragraph      |
| `/text/words/{count}`        | Returns count random words      |
| `/text/paragraphs/{count}`   | Returns count random paragraphs |
| `/text/sentences/{count}`    | Returns count random sentences  |
| `/img/{width}x{height}.png`  | Returns a png image             |
| `/img/{width}x{height}.jpg`  | Returns a jpg image             |
| `/img/{width}x{height}.jpeg` | Returns a jpeg image            |
| `/img/{width}x{height}.gif`  | Returns a gif image             |

### Text routes

`http://127.0.0.1:8000/text/word`

![word](https://imgur.com/iAHbQMA.png)

`http://127.0.0.1:8000/text/sentence`

![sentence](https://imgur.com/g4UyvKL.png)

`http://127.0.0.1:8000/text/paragraph`

![paragraph](https://imgur.com/xQWqyJo.png)

### Image routes

Query params

When setting fg or bg options you can use alpha channel: `1234` or `12345678`

`bg` - background color
`fg` - foreground color
`border` - border and diagonal width

`http://127.0.0.1:8000/img/256x256.png`

![256x256 image](https://imgur.com/E2jD5rm.png)

`http://127.0.0.1:8000/img/256x256.png?bg=00FFFF&fg=ff0&border=10`

![256x256 image](https://imgur.com/BVdtFre.png)
