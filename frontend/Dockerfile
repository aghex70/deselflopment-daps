# Build React static files
FROM node:alpine as build
RUN export NODE_OPTIONS=--openssl-legacy-provider
WORKDIR /daps/frontend/
COPY package.json ./
COPY package-lock.json ./
RUN npm config set fetch-retry-mintimeout 200000
RUN npm config set fetch-retry-maxtimeout 1200000
RUN npm  cache clear --force
RUN npm config rm proxy
RUN npm config rm https-proxy
RUN npm ci --force
COPY . .
RUN #npm run build
CMD ["npm", "run", "start"]