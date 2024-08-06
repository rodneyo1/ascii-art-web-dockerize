# ascii-art-web-dockerize
Ascii-art-web entails containerizing of the ascii-art-web project using Docker. The project includes a Dockerfile to build the Docker image and run the web server in a container.

## Features
- Converts text to ASCII art
- Displays the ASCII art on a HTML template on a web browser.
- Utilizes specific graphical templates for ASCII representation
- Containerization of the application using Docker.

## Prerequisites

- Docker: Ensure Docker is installed on your system. You can download it from [Docker's official website](https://www.docker.com/get-started).
- Go: Make sure you have Go installed. Instructions can be found at [Go's official website](https://golang.org/doc/install).


## Installation

1. Clone the repository:

    git clone <a>https://learn.zone01kisumu.ke/git/rodnochieng/ascii-art-web-dockerize</a>


2. Navigate to the project directory:

    ```bash
    cd ascii-art-web-dockerize
    ```
 
## Usage
To be able to containerize the application, you first need to build the docker image: 
```
docker build -t <name_of_the_image> .
```

or

```
docker image build -f Dockerfile -t <name_of_the_image> .
```
The second step is to run the Docker Container
```
docker run -d -p <port_you_what_to_run> --name <name_of_the_container>  <name_of_the_image>
```

or  
```
docker container run -p <port_you_what_to_run> --detach --name <name_of_the_container> <name_of_the_image>
```

You can also build a docker image and run a docker container by running a shell script as it can can streamline the process and make it repeatable. 

Once you have created your shell script, you can run it ising the following command:
```
sh <name_of_script>.sh
```
For example:
```
sh run.sh
```


When the container is running you go to your browser and type the link:
```
'http://localhost:<port_you_want_to_run>'
```

For example: 
```
'http://localhost:8080'
```

You should see the main page where you can input text and select a banner. After submitting, you will be able to see the generted ASCII art in the specified format.

## Testing 
To run the tests present do the following:

Run the test using this command:

```
go test ./server && go test ./asciiart
```

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by:

[Thadeus Ogondola](https://learn.zone01kisumu.ke/git/togondol/)

[Rodney Ochieng](https://learn.zone01kisumu.ke/git/rodnochieng)

