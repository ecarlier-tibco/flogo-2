package sendWSMessage

var jsonMetadata = `{
  "name": "sendSlackMessage",
  "version": "0.0.1",
  "title": "Send Slack Message",
  "description": "This activity sends a message to a Slack Channel",
  "inputs":[
    {
      "name": "Webhook",
      "type": "string",
      "required": true
    },
    {
      "name": "Channel",
      "type": "string",
      "required": false
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    },
    {
      "name": "Username",
      "type": "string",
      "required": false
    },
    {
      "name": "Iconemoji",
      "type": "string",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}`
