# Python image to use.
FROM python:3.13-alpine

# Set the working directory to /app
WORKDIR /app

# copy the requirements file used for dependencies
COPY . .

# Install any needed packages specified in requirements.txt
RUN pip install uv
RUN uv sync

EXPOSE 10999

# Run app.py when the container launches
ENTRYPOINT ["uv", "run", "business_agent"]