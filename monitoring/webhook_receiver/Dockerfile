FROM python:3.12-slim

WORKDIR /app

COPY webhook_script.py .

RUN pip install --no-cache-dir Flask

# Expose the port your app runs on
EXPOSE 5000

# Run the Flask app
CMD ["python", "webhook_script.py"]
