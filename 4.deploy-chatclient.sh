gcloud run deploy ucpdemo-chat --source ./a2a/chat-client --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION \
  --port 8080 --env-vars-file .env --min-instances 0 --allow-unauthenticated
