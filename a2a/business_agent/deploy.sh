gcloud run deploy a2a-business-agent-bap7 --source . --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION --env-vars-file .env --min-instances 0 --port 10999 --no-allow-unauthenticated
