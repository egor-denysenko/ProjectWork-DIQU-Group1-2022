FROM node:16.15-alpine3.16

COPY package*.json ./

RUN npm install npm@latest -g

COPY ./ .
#RUN npm install
# If you are building your code for production
#RUN npm ci --only=production
RUN npm install --no-optional && npm cache clean --force

EXPOSE 9022

CMD [ "npm","start" ]