#!/bin/bash

# Set the login URL and the credentials
login_url=$LOGIN_URL
username=$LOGIN_USERNAME
password=$LOGIN_PASSWORD

# Set the email parameters
from_email=$FROM_EMAIL
to_email=$TO_EMAIL
smtp_server='smtp.gmail.com'  # Replace with your SMTP server
smtp_port=587  # Replace with your SMTP port
smtp_username=$FROM_EMAIL
smtp_password=$SENDER_PASSWORD

# Try to login to the website
response=$(curl -d "username=$username&password=$password" "$login_url")

# If the login was successful, print a message
if [ $? -eq 0 ]; then
    echo 'Login successful'
else
    # Send an email with the error message
    echo "$response" | mail -s 'Login error' -a "From: $from_email" "$to_email" +7 -delete