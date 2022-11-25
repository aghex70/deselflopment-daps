# Build React static files
FROM node:alpine as build
RUN export NODE_OPTIONS=--openssl-legacy-provider
WORKDIR /daps/frontend/
COPY package.json ./
COPY package-lock.json ./
RUN npm config set fetch-retry-mintimeout 20000
RUN npm config set fetch-retry-maxtimeout 120000
RUN npm ci --force
COPY . .
RUN npm run build
CMD ["npm", "run", "start"]