run_mysql_and_service_and_migration() {
    run_mysql_and_service_without_migration
    run_migration_migrate
}

run_mysql_and_service_without_migration() {
    echo "Start Running Mysql and Service"
    docker-compose up -d
    echo "Mysql and Service Started"
}


run_migration_info() {
    echo "Starting migration info"
    cd migration
    docker run --network ina-adm-regions_net --rm -v "$(pwd)":/opt/maven -w /opt/maven maven:3.6-openjdk-8 mvn flyway:info
    echo "Finished migration"
}

run_migration_migrate() {
    echo "Starting migration migrate"
    cd migration
    docker run --network ina-adm-regions_net --rm -v "$(pwd)":/opt/maven -w /opt/maven maven:3.6-openjdk-8 mvn flyway:migrate
    echo "Finished migration"
}

echo "\n===============Indonesia Administrative Regions===============\n"

echo "1. Run All (Mysql, Service and Migration)"
echo "2. Run All (Mysql, Service without Migration)"
echo "3. Run Migration Info"
echo "4. Run Migration Migrate"

read -p "Pick options : " option

if [ $option == 2 ]
then
    run_mysql_and_service_without_migration
elif [ $option == 3 ]
then
    run_migration_info
elif [ $option == 4 ]
then
    run_migration_migrate
else
    run_mysql_and_service_and_migration
fi