FROM node:14-alpine
WORKDIR /usr/src/app

ENV PORT 3000

COPY package.json package-lock.json ./ 
RUN npm i

COPY . ./
RUN npm run build

EXPOSE 3000
# Running the app
CMD [ "npm", "start" ]