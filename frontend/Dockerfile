FROM node:alpine
WORKDIR /daps/frontend/
COPY package.json ./
COPY package-lock.json ./
COPY . .
RUN npm install --force
CMD ["npm", "run", "start"]