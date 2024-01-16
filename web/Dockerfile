# Node image to build React app
FROM node:14 AS build
WORKDIR /app

# Declare arguments passed from docker-compose
ARG NGINX_ENV
ARG NGINX_HOST
ARG NGINX_PORT

# Set environment variables
ENV REACT_APP_NGINX_ENV=$NGINX_ENV
ENV REACT_APP_NGINX_HOST=$NGINX_HOST
ENV REACT_APP_NGINX_PORT=$NGINX_PORT

COPY package.json package-lock.json ./
RUN npm install --verbose
COPY . ./
RUN npm run build

# The final image uses the static files from the build step
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx.frontend.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
