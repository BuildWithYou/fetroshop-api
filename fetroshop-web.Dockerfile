# Use an official Node.js image as the base image
FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /app

# Clone the GitHub repository
RUN apk --no-cache add git \
   && git clone https://github.com/BuildWithYou/fetroshop-web.git .

# Install dependencies
RUN npm install

# Build the Next.js app
RUN npm run build

# Expose the port that the app will run on
EXPOSE 3003

# Set the command to run your Next.js app
CMD ["npx", "next", "dev", "--", "-p", "3003"]
