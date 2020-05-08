# Golang SendGrid Email Cloud Function

This is a Google Cloud Function in Go to send an email powered by SendGrid by simply making a POST HTTP request.

## Usage

Your request body should have the following:

```json
{
  "from": "your email",
  "fromName": "your name",
  "to": "recipient email",
  "toName": "recipient name",
  "subject": "email subject line",
  "html": "<div>some html</div>",
  "plainText": "some plain text."
}
```
