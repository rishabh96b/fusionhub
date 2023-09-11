# fusionhub
A sample web application fabricated with Golang, Go Templates, HTML &amp; CSS. This app serves as a practical demonstration of the deployment patterns like Blue-Green, Canary etc. on k8s.

## How It Works

### Dynamic Content

The application dynamically displays a message on the webpage, and the background color of the page is customizable based on the value of an environment variable (`FUSIONHUB_COLOR`).


The primary feature of this application is its ability to demonstrate blue-green deployment. It offers two distinct versions: "blue" and "green," each represented by a different environment variable (`FUSIONHUB_COLOR=blue` or `FUSIONHUB_COLOR=green`). Depending on the value of this environment variable, the application will display either a blue-themed or green-themed webpage. This enables seamless switching between versions without any downtime or interruptions.

## Running the Application
```bash
docker run -e COLOR=blue -p 8080:80 rishabh96b/fusionhub
```
Above command will run the blue application.


