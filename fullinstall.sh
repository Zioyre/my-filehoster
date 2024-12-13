# looking for a fix run as sudo

mkdir /mnt/storage/

mkdir /mnt/storage/files

cd filehandler/

docker compose -f "compose.yml" up -d --build

cd docker-share-files/

docker compose -f "compose.yml" up -d --build


cp html/ /var/www/

echo "Done!"
