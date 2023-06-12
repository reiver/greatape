# GʀᴇᴀᴛAᴘᴇ

⚠️ This project is under heavy development and should not be used in production yet.
<br /><br />

<img style="width:150px" src="https://raw.githubusercontent.com/reiver/greatape/master/assets/media/greatape-logo.png" />

[![image](https://img.shields.io/badge/Go-00A7D0?style=for-the-badge&logo=go&logoColor=white)](https://go.dev) [![image](https://img.shields.io/badge/ActivityPub-DD307D?style=for-the-badge&logoColor=white)](https://www.w3.org/TR/activitypub/) [![image](https://img.shields.io/badge/JSON--LD-FF6600?style=for-the-badge&logo=json&logoColor=white)](https://json-ld.org) [![image](https://img.shields.io/badge/PostgreSQL-40668D?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org)

**greatape** is a free **social audio & video** social-media platform that can be used via an app.

**greatape** is a Fediverse technology that supports Federation via ActivityPub.

## 🏎️ Running the Project

### 🚀 Using Go and Postgres

1. Clone the project repository:
    ```
    git clone https://github.com/reiver/greatape
    ```
2. Navigate to the project directory: 
    ```
    cd greatape
    ```
3. Create an empty Postgres database.
4. Update the `config.yaml` file in the project root directory with the actual values for your database.
5. Download the project dependencies:
    ```
    go mod download
    ```
6. Run the project:
    ```
    go run main.go
    ```

### 🐳 Using Docker

1. Clone the project repository:
    ```
    git clone https://github.com/reiver/greatape
    ```
2. Navigate to the project directory: 
    ```
    cd greatape
    ```
3. Build the Docker image:
    ```
    docker build -t greatape .
    ```
4. Replace the environment variables below with your own and run the Docker container:
    ```
    docker run \
        --name greatape \
        -e PROTOCOL=https \
        -e FQDN=yourdomain.com \
        -e PORT=7080 \
        -e POSTGRES_HOST=127.0.0.1 \
        -e POSTGRES_PORT=5432 \
        -e POSTGRES_DATABASE=greatape \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_PASSWORD=password \
        -p 7080:7080 \
        greatape
    ```

### 🐳 Using docker-compose

1. Clone the project repository:
    ```
    git clone https://github.com/reiver/greatape
    ```
2. Navigate to the project directory: 
    ```
    cd greatape
    ```
3. Run the Docker containers using docker-compose: 
    ```
    docker-compose up
    ```

## 👥 Team

The following is a list of the people who are actively working on Great Ape (in alphabetical order):

| Name                     | Role                       | Online                                                                             |
|--------------------------|----------------------------|------------------------------------------------------------------------------------|
| Charles Iliya Krempeaux  | lead, product, engineering | [🐘](https://mastodon.social/@reiver) [🕸️](http://changelog.ca/)                   |
| Chet Earl Woodside       | illustration               | [🕸️](http://cosmicblend.ca/)                                                       |
| Chris Trottier           | product, qa                | [🐘](https://calckey.social/@atomicpoet) [📷](https://peerverse.space/atomicpoet) |
| Farzaneh Amini           | ux                         | [🕸️](https://www.behance.net/farzanehamini)                                        |
| Massoud Seifi            | engineering                | [🐘](https://mastodon.social/@accesstoken)                                         |
| Meysam Mousavi           | engineering                |                                                                                    |
| Nariman Movaffaghi       | engineering                |                                                                                    |
| Nastaran Ahmadi Bonakdar | engineering                |                                                                                    |
