# **URL Shortner**
It is a URL shortener written in golang. Shorten URLs, manage your links!

##  <u>**Introduction**</u> 👋🏻

A URL shortener is a simple tool that takes a long URL and turns it into a shorter URL for you. As social media emerged, lengthy URLs started to present a problem. Originally, Twitter limited their messages to a maximum of 140 characters, and they counted every character in a link. This meant that a lengthy URL could take up your entire tweet. URL shortener tools emerged as a solution to this sharing problem.
An URL shortener ensures that you get the right messages out to your audience without taking up too much space in your social posts.

## <u>**Getting started**</u> ▶️

To start using the project you need  to install Golang (I tried it for Go 1.19).   
1. Run these commands:
```
git clone https://github.com/kitarp29/urlShortner.git
cd Infra
```

2. Make sure you have Golang installed on your system.
```
go --version 

```

3. Now, we install the pkgs
```
go mod tidy
```

4. Finally run the server using:

```
go run main.go
```

> API will be running at <a href="localhost:8000/"> localohost:8000 </a> now.
  ## **The project is up and running🔥!**
  <hr>

## <u>**Running in a Docker container**</u> 📦
> Ensure you have docker installed in your system

1. Make sure you are inside the folder of the project.
2. Run this command to build the Docker Image:
```
docker build . -t urlshort
```
3. You can check the image by running this:
```
docker images | grep urlshort
```
4. Run a container using this:
```
docker run -p 8000:8000 urlshort:latest  
```
> API will be running at <a href="localhost:8000/"> localohost:8000 </a> now.
  ## **The project is up and running🔥!**
 

Refer PostMan Collection: [Here](https://www.getpostman.com/collections/a1f2466abe2ea9c6a977)
### Thanks for the interest in my API :)
<hr>
