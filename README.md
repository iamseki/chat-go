<h2 align="center">
  <a href="https://github.com/iamseki?tab=repositories">
    <img alt="Open Weather Logo" src="https://www.pngkey.com/png/full/376-3767568_addthis-sharing-buttons-gopher-with-glasses.png" width="200px" />
  </a>
</h2>
<h2 align="center">
  Simple real time chat App 
</h2>

<p align="center">This simple project upgrades the connection from HTTP to the <strong>WebSocket</strong>.</p>
 <p align="center">An <strong>unique chat room</strong> is created when user is signed in and everyone talks to everyone.
  </br>
  <small>User authentication is made by OAuth protocol</small>
 </p>


<p align="center">
  <a href="https://www.linkedin.com/in/christian-seki/">
    <img alt="Made by Christian Seki" src="https://img.shields.io/badge/made%20by-Christian%20Seki-blue">
  </a>
</p>

---
## :wrench: Running the project

> if you dont have golang in your environment, go and install : https://golang.org/dl/  

Now you're ready to go ;)

#### :floppy_disk: install dependencies
  In go, you don't have to worry about dependencies, because i'm using go modules.

#### :key: Setup Env vars

> create **.env** file in root to set **OAuth credentials**, like this: 
  ```  
        SECURITY_KEY= PUT-ANY-HASH-HERE-YOU-CAN-GENERATE-ONE-IN-MD5-HASH-GENERATOR-ONLINE

        GOOGLE_CLIENT_ID= YOUR-GOOGLE-CLIENT-ID
        GOOGLE_SECRET_KEY= YOUR-GOOGLE-SECRET-KEY

        GITHUB_CLIENT_ID= YOUR-GITHUB-CLIENT-ID
        GITHUB_SECRET_KEY= YOUR-GOOGLE-SECRET-KEY
        
        FACE_CLIENT_ID= YOUR-FACEBOOK-CLIENT-ID
        FACE_SECRET_KEY= YOUR-FACEBOOK-SECRET-KEY
  ```
  **getting OAuth credentials**
  - <a href="https://github.com/settings/developers">github</a>
  - <a href="https://console.cloud.google.com/apis/credentials">google</a> 
  
  #### :running: Running the server
  
  execute `go run *.go` 
  
  or `go build` and `./gochat`
  
  > expected output in terminal: `😎 Server listening on:  :8080`
  
  chat endpoint : **`localhost:8080/chat`**
  



