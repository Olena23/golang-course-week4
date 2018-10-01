# Неделя 4


## Задания

К выполнению 2 задания

1. Echo
2. Города

### 1. Эхо

Необходимо сделать эхо-сервер, который принимает любой http запрос и в ответ возваращает по прообразу http://httpbin.org

Список:

- http://httpbin.org/ip Returns Origin IP.
- http://httpbin.org/user-agent Returns user-agent.
- http://httpbin.org/headers Returns header dict.
- http://httpbin.org/get Returns GET data.
- http://httpbin.org/post Returns POST data.
- http://httpbin.org/put Returns PUT data.
- http://httpbin.org/delete Returns DELETE data
- http://httpbin.org/status/:code Returns given HTTP Status code.
- http://httpbin.org/response-headers?key=val Returns given response headers.
- http://httpbin.org/cookies Returns cookie data.
- http://httpbin.org/cookies/set/:name/:value Sets a simple cookie.
- http://httpbin.org/stream/:n Streams n–100 lines.
- http://httpbin.org/delay/:n Delays responding for n–10 seconds

### 2. Города

За основу взять предыдущее решение городов. Реализовать веб-сервер

1) /init - начинает игру, возвращает 1й город 
2) /get-answer/{city_name} - получаете вариант ответа
3) /submit - отправляете город, получаете новый
4) возвращаетесь к пункту 2 


## Cookie
- https://ru.wikipedia.org/wiki/Cookie
- https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.1.html