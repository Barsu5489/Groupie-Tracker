# Groupie Tracker

## Table of Contents
- [Description](#description)
- [Features](#features)
- [Instalation](#instalation)
- [Implementation](#implementation-details-algorithm)
- [Authors](#authors)
- [References](#references)

## Description
`Groupie Tracker` is a web application that allows users to track and manage their favorite musical artists and bands. The application provides information on artists' biographies, discographies, tour dates, and venues. Users can also receive notifications for upcoming events and new releases from their favorite artists.

### Features

#### 1. Band Profiles

**Band Members:** Display a list of band members with their names and roles in the band.

**Dates:** Show the year the band was formed Display the date when the band released its first album.

#### 2. Discography

**Album History:** Provide a timeline of albums starting from the first album release date, showing the progression of the band's discography over time.
**Album Covers:** Include album cover images alongside release dates and track listings.

#### 3.Image Gallery

**Band Images:** Display images of the band, possibly with album art, concert photos, or promotional material.

### Instalation

1. Clone the repository:
    ```go
    git clone https://https://learn.zone01kisumu.ke/git/kada/groupie-tracker.git
    cd groupie-tracker
    ```
2. Install Go and set up your Go environment if you haven't already. You can download and install Go from the official Go website.

3. Run the application:

    ```go
    go run main.go
    ```
4. Open your web browser and navigate to http://localhost:8080/.


### Implementation Details: Algorithm

`Groupie Tracker` utilizes a series of Go packages to manage different aspects of the application. Here's an overview of the implementation:

#### 1. Main Server Setup:
 The main function sets up an HTTP server that listens on port 8080. It defines several route handlers: 
`/` - for the homepage
`/endpoint` - for handling different URL paths

#### 2. Frontend HTML:

The HTML files serve as the user interface for the Groupie Tracker application. These files provide the structure and layout for the different pages of the application, including:

**Band Profiles:** Displays detailed information about each band, including member details, creation date, first album release date, and images.

**Tour Dates:** Shows a list of upcoming tour dates and venues for the selected band.

**Image Gallery:** Displays images related to the band, including album covers and members. This gallery is often integrated within the band profile page, enhancing the visual appeal and providing users with a richer experience.

Each HTML file is designed to be responsive and user-friendly, ensuring a smooth experience across different devices. The HTML files work in conjunction with CSS for styling and JavaScript for interactivity, creating a seamless and engaging user experience..
#### 3. Backend Handlers:
The backend handlers manage the routing and processing of user requests in the Groupie Tracker application. These handlers are responsible for routing request, process search queries, serving band profiles, handling tour dates, and data retreaval and API intergration.

With these details, you should be able to understand and run Groupie Tracker application effectively. Feel free to explore the code and modify it according to your needs.

## Authors

- [kada](https://learn.zone01kisumu.ke/git/kada)
- [ebarsula](https://learn.zone01kisumu.ke/git/ebarsula)
- [Malika](https://learn.zone01kisumu.ke/git/masman)


## References

- JSON File and HTML

## License
This project is licensed under the MIT protection act. [LICENSE](./LICENSE)
