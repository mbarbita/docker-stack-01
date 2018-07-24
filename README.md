# docker-stack-alpha

    cd docker=stack=alpha
    mkdir mysql-data portainer-data grafana-data
    cd web-files
    ./build-app.sh
    cd ..
    
edit **db_root_password.txt** and replace **"password"** with a new one

    docker swarm init --advertise-addr <ip>
    docker stack deploy -c docker-compose.yml servers
