# docker-stack-alpha

    cd docker=stack=alpha
    mkdir mysql-data portainer-data grafana-data
    cd web-files
    ./build-app.sh
    cd ..
    
edit **db_root_password.txt** and replace **"password"** with a new one

    docker swarm init [--advertise-addr <ip>]
    docker stack deploy -c docker-compose.yml <servers>

## mysql
    CREATE DATABASE grafana;
    CREATE USER 'grafanaro' IDENTIFIED BY 'password';
    GRANT SELECT ON grafana.* TO 'grafanaro';

    CREATE USER 'grafanarw' IDENTIFIED BY 'password';
    GRANT ALL ON grafana.* TO 'grafanarw';
    
### misc
    ALTER USER user IDENTIFIED BY 'auth_string';
    ALTER TABLE table_name AUTO_INCREMENT = 1;
