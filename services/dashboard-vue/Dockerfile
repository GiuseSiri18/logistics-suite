FROM node:16-alpine

WORKDIR /app

# Installiamo Vue CLI globalmente
RUN npm install -g @vue/cli

# Copiamo il package.json e installiamo le dipendenze
COPY package*.json ./
RUN npm install

# Copiamo il resto del codice
COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]