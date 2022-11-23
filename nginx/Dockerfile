# Build React static files
FROM node:alpine as build
WORKDIR /daps/frontend/
COPY ./frontend/package.json ./
COPY ./frontend/package-lock.json ./
RUN npm install -g npm@9.1.2
RUN npm ci --silent --force
COPY ./frontend ./
RUN npm run build
CMD ["npm", "run", "start"]


# Copy React static files and start nginx
FROM nginx:stable-alpine
COPY --from=build /daps/frontend/build /usr/share/nginx/html

COPY --from=build /app/nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]