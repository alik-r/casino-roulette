FROM nginx:1.21.3-alpine

COPY nginx.conf /etc/nginx/nginx.conf

COPY ./index.html /usr/share/nginx/html/index.html
COPY ./register.html /usr/share/nginx/html/register.html
COPY ./roulette.html /usr/share/nginx/html/roulette.html
COPY ./leaderboard.html /usr/share/nginx/html/leaderboard.html
COPY ./profile.html /usr/share/nginx/html/profile.html
COPY ./reset_password.html /usr/share/nginx/html/reset_password.html

COPY ./style.css /usr/share/nginx/html/style.css
COPY ./styles/ /usr/share/nginx/html/styles/
COPY ./fonts/ /usr/share/nginx/html/fonts/
COPY ./images/ /usr/share/nginx/html/images/
COPY ./scripts/ /usr/share/nginx/html/scripts/

EXPOSE 80