# Microsoft Teams Webhook Message Action

This action sends a message to a specific MSTeams Webhook.

## Inputs

### `WEBHOOK_URL`
**Required** 

### `TITLE`
**Required** Title of the message

### `MESSAGE`
**Required** The message body

## Example usage
```yaml
 - name: MSTeams Webhook message
   uses: charmitro/msteams-webhook-message-action@v0.5
   env:
     WEBHOOK_URL: <url>
     TITLE: "Test"
     MESSAGE: "Test message from GitHub action."
```
