<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->




<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->


<h3 align="center">Survey</h3>

  <p align="center">
    Creating surveys and functionality for them
    <br />
    <br />
    <a href="https://github.com/Uikola/survey">View Demo</a>
    ·
    <a href="https://t.me/uikola">Report Bug</a>
  </p>


<!-- ABOUT THE PROJECT -->
## About The Project

This project is an application for creating and managing surveys. Here are the methods that this project implements:
- `POST /api/start-survey` start a survey
- `DELETE /api/delete-survey/{survey_id}` delete a survey
- `POST /api/get-result` get the survey result
- `POST /api/add-ans` add an answer to the survey
- `DELETE /api/delete-ans/{survey_id}/{ans_id}`
- `POST /api/vote` vote for the answer


### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.


<!-- GETTING STARTED -->
## Getting Started

To start the project, you need to perform the following actions

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
2. Run go mod tidy
   ```sh
   go mod tidy
   ```
3. Create a folder named envs in the config folder, and then create a dev.env file and enter the following information:
* PORT=:YOUR_PORT
* CONN_STRING=host=your_host port=your_port user=your_user password=your_pass dbname=your_db_name sslmode=disable 
* DRIVER_NAME=postgres

4. Run the app
   ```sh
   go run cmd/main.go 
   ```

<!-- CONTACT -->
## Contact

Yuri - [@telegram](https://t.me/uikola) - ugulaev806@yandex.ru

Project Link: [https://github.com/Uikola/survey](https://github.com/Uikola/survey)


