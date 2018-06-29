Tweet Stack = stack overflow with twitter tags

#Install Docker on respective OS and run the following command on your system to setup a mongo container where we will be persisting our data.
docker run -d -v data:/data/db/ --name mongo_cont -p 27017:27017 mongo

#Access Mongo DB
docker exec -it <mongo_cont> /bin/sh
mongo

