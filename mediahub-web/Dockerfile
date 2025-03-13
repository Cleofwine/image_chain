# build front-end
FROM node:lts-alpine AS frontend
RUN npm config set registry https://registry.npmmirror.com
WORKDIR /app
COPY . /app

RUN npm install
RUN npm run build


FROM nginx:latest AS web 
COPY --from=frontend /app/dist/ /usr/share/nginx/html/mediahub-web/
ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx","-g","daemon off;" ]