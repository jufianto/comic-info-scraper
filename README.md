# Comic Info Scraper

This project is a fun tool designed to fetch the latest updates from comic websites.

**NOTE**: This project is solely focused on collecting data.

## What It Does:
- Fetches the homepage of comic sites (currently supporting **AsuraScans**).
- Retrieves the latest updated **Title** and **Chapter** information.
- Saves the data to **YAML** files.
- Optionally stores the data in **Firestore** for persistence.

## Installation Instructions:
1. Clone this project to your local machine.
2. Run `go mod tidy`, or `go mod vendor` if you prefer vendoring dependencies.
3. Ensure that **Chrome** or **Chromium** is installed on your system.
4. To store data in Firestore:
   - Update the config in `cmd/config/scraper.yaml` to enable Firestore storage.
   - Set your **project ID** and **collection name** in the config file.
5. For Firestore integration, create a **service account key** in the **Google Cloud Console**.
6. Run the application with:
   ```sh
   go run cmd/main.go
   ```
   **NOTE:** You must run the command from the project root directory.

## Notes
* Future updates will focus on improving Firestore persistence and keeping track of the latest comic that was read.
* Currently, Firestore configuration is optional. You can use the tool without configuring Firestore.