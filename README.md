# Currency Exchange Rates

## Description
This project consists of a Go backend and a React frontend that provides currency exchange rates. It efficiently handles exchange rates fetched from an external API and includes a fallback to hardcoded rates if the API is unavailable. The frontend enables users to interact with the system to view exchange rates and perform currency conversions.

## Features
Fetch exchange rates from an external API.
Fallback to hardcoded rates when the API is unavailable.
Interactive UI to display and convert exchange rates.
Multi-currency support with dynamic selection.

## Technology Stack
Backend: Go (Golang)
Frontend: ReactJS
External API: ExchangeRatesAPI.io

## Getting Started
Prerequisites
Go (Version 1.15 or higher recommended)
Node.js and npm
Git

## Installation
Clone the repository:

bash
Copy code
git clone https://github.com/kyhinds/app_exchange_rates.git
cd app_exchange_rates
Set up the backend:

bash
Copy code
go build ./...
go run cmd/server/main.go
Build the frontend:

bash
Copy code
cd client
npm install
npm run build
The Go server will serve the static files from the client/build directory.

## Usage
The backend server starts on http://localhost:8080.
Access the application at http://localhost:8080 since the Go server serves the built React files.
API Endpoints
GET /rates - Retrieves exchange rates for specified base and target currencies.
GET /convert - Converts an amount from one currency to another using real-time or fallback rates.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contact
Kyle Hinds - kyhinds@gmail.com
Project Link: https://github.com/kyhinds/app_exchange_rates