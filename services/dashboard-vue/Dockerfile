FROM node:14

WORKDIR /app

# Installiamo Vue CLI per poter inizializzare il progetto
RUN npm install -g @vue/cli

# Esponiamo la porta di Vue
EXPOSE 8080

# Il comando di avvio per lo sviluppo
CMD ["npm", "run", "serve"]