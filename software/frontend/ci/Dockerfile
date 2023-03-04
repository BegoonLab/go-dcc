FROM node:19 as base

USER node

WORKDIR /var/www/app

COPY package*.json ./

FROM base as development

EXPOSE 8080

CMD [ "npm run dev" ]