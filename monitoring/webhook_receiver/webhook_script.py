from flask import Flask, request, jsonify, render_template_string
from datetime import datetime

app = Flask(__name__)

# In-memory store for alerts
alerts_store = []

# Simple HTML template to display alerts
HTML_TEMPLATE = """
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Received Alerts</title>
  <style>
    body { font-family: Arial, sans-serif; margin: 20px; }
    .alert { border: 1px solid #ccc; padding: 10px; margin-bottom: 10px; }
    .status-firing { background-color: #ffdddd; }
    .status-resolved { background-color: #ddffdd; }
    h2 { margin: 0; }
  </style>
</head>
<body>
  <h1>Received Alerts</h1>
  {% for alert in alerts %}
    <div class="alert status-{{ alert['status'] }}">
      <h2>{{ alert['labels']['alertname'] }} ({{ alert['status'] }})</h2>
      <p><strong>Severity:</strong> {{ alert['labels'].get('severity', 'N/A') }}</p>
      <p><strong>Summary:</strong> {{ alert['annotations'].get('summary', '') }}</p>
      <p><strong>Description:</strong> {{ alert['annotations'].get('description', '') }}</p>
      <p><em>Received at {{ alert['received_at'] }}</em></p>
    </div>
  {% endfor %}
</body>
</html>
"""

@app.route("/alerts", methods=["POST"])
def alerts():
    data = request.json
    if data:
        for alert in data.get("alerts", []):
            alert["received_at"] = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            alerts_store.append(alert)
            print(f"Received alert: {alert['labels'].get('alertname')}, status: {alert['status']}")
    return jsonify({"message": "Alert received"}), 200

@app.route("/", methods=["GET"])
def view_alerts():
    return render_template_string(HTML_TEMPLATE, alerts=reversed(alerts_store))

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
