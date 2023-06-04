# Transcribe website using AWS services

I wanted to learn about the services that AWS offers. While i did that i also challenged myself to write some Ruby and GO.


#### How it works

The user uploads a file to an S3 bucket throuh a website. That triggers a lambda function written in Ruby which then uploads the file to an AWS transcript service. 
The transcript service then uploads the file to the same S3 bucket which the user checks and in the end receives the transcription. 
<img width="1191" alt="Skærmbillede 2023-06-03 kl  20 54 35" src="https://github.com/gbalslevn/transcribe/assets/97167089/385f3d4c-154c-4ece-8655-3b6935344a03">

#### Live demo

https://github.com/gbalslevn/transcribe/assets/97167089/39822cc9-6dcf-41fa-a551-1ba2500668d7

Try it here: [Website](https://transcribe-audio-to-danish.onrender.com/)


#### What did i accomplish?
The project works and i learned about the AWS services (S3, AWS Transcribe, Lambda, API Gateway, CloudWatch) and also about Ruby and Go. 
It is maybe not so rational to use so many different services and languages for such a small project, but it is ideal for the learning experience. 

#### What would i do different?
Right now the user uploads the file from the frontend which also exposes the AWS keys to the bucket. Quite bad practice but i just wanted to complete the project as it was a weekend project and i ran out of time. My plan in the beginning (look at pic below) was to upload the file to some Go code ('main.go' which i have included in this repository). I originally thought about using the AWS API Gateway to access the Go code. The go code would be hosted as a lamda function. 
It would also be better practice to have an input and output bucket. Right now the user constantly asks the S3 bucket if there is a file. In the future a more elegant solution could be implemented here. 


<img width="947" alt="Skærmbillede 2023-06-03 kl  21 13 22" src="https://github.com/gbalslevn/transcribe/assets/97167089/8303d502-4f15-4951-9196-0e7753244ac8">






