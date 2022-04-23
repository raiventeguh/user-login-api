FROM mongo

COPY users.json /users.json

CMD mongoimport --uri mongodb://mongodb:27017/mongodb -u mongodb -p mongodb --authenticationDatabase admin --collection user --type json --file /users.json --jsonArray