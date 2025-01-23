# Go Get Info

This Go project retrieves information from a specified URL and saves the data into a CSV file. The project is designed to be functional and modular, making it easy to extend and maintain.

## Project Structure

- `main.go`: The entry point of the application. It orchestrates the fetching and saving of data.
- `fetch.go`: Contains the logic for fetching data from the specified URL.
- `save.go`: Contains the logic for saving the fetched data into a CSV file.
- `go.mod`: The Go module file.

## Functionality

1. **Fetching Data**: The application fetches data from the URL `https://carinx.ir/api/fund-raises?page=<page>&platforms=halalfund,dongi,karencrowd,zeema,ifund,hamafarin,novincrowd,charisma,ibcrowd,hamashena,sepehrino,pulsar&activeOpportunitiesOnly=false` for pages 1 to 1000.
2. **Saving Data**: The fetched data is saved into a CSV file named `items.csv`. If the file does not exist, it is created with a header row containing the property names.

## How to Run

1. **Clone the repository**:
    ```sh
    git clone <repository-url>
    cd go-get-info
    ```

2. **Run the application**:
    ```sh
    go run main.go
    ```

## Dependencies

This project does not have any external dependencies. It uses the standard Go library.

## Data Structure

The data fetched from the URL is expected to be in JSON format with the following structure:

```json
{
  "items": [
    {
      "_id": "string",
      "title": "string",
      "url": "string",
      "companyName": "string",
      "fundRaiseAmount": "int64",
      "raisedFundAmount": "int64",
      "investorsCount": "int",
      "expectedReturnRate": "int",
      "collateralType": "string",
      "duration": "string",
      "fundRaisingEndDate": "string",
      "bannerURL": "string",
      "platform": {
        "_id": "string",
        "name": "string",
        "code": "string",
        "website": "string"
      }
    }
  ],
  "count": "int",
  "total": "int"
}
