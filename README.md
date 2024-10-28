# HackGlobal Project: Lyf Guide

This repository is a project for hackglobal 2024 where we create an app for lyf Funan and focus on enhancing the customer satisfaction through their experience.

## Table Of Contents

- [Description](#description)
- [Prerequisites](#prerequisite)
- [Installing](#installing)
- [Executing the Program](#executing-the-program)
- [Authors](#authors)

## Description

Our project is called Lyf Guide, where is designed to transform the travel experience for young solo travelers. By automating itinerary creation, the app personalizes daily travel plans based on user preferences, allowing travelers to share and explore their journeys with companions. With features like Travel Buddy Matchmaking, users can connect with like-minded travelers for shared experiences, enhancing their trips through social engagement. Additionally, the app offers a destination-specific marketplace for essentials, enabling users to pre-purchase items or access buddy discounts, simplifying logistics and reducing costs. Our ultimate goal is to create a seamless, enriched travel ecosystem that fosters satisfaction, loyalty, and a stronger sense of community among travelers.

## Getting Started

### Prerequisites

These are the prerequisites that you need to prepare to run the program:
* Go (with the minimum version 1.2.3)
* PostgreSQL
* Docker

### Installing

On the client side, you need to first install the npm packages by doing:
1. Change directory to client
  ```sh
  cd client
  ```

2. Install npm packages
  ```sh
  npm install
  ```

And on the server side, you need to prepare the database:
1. Migrate database scheme
  ```sh
  go run /cmd/migration/main.go
  ```

2. Create seeder for the created database
  ```sh
  go run /cmd/seeder/main.go
  ```

### Executing the Program

* How to run the program
* Step-by-step bullets
To run the program, you need to follow these steps:
1. Open a terminal and run the server
  ```
  go run /cmd/server/main.go
  ```

2. Open another terminal, and move to client directory
  ```sh
  cd client
  ```

3. Run the client
  ```sh
  npm run dev
  ```

## Authors

Below are the people who contributed to this project:
* [@syahba](https://github.com/syahba)
* [@W-ptra](https://github.com/w-ptra)
* [@mad-af](https://github.com/mad-af)
