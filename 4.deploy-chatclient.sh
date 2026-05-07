# get service URL
SERVICE_URL=$(gcloud run services describe ucp-chat-service-bap7 --format 'value(status.url)' --region $GOOGLE_CLOUD_LOCATION --project $GOOGLE_CLOUD_PROJECT)

# deploy to Cloud Run if it doesn't exist
if [ -z "${SERVICE_URL}" ]; then
  gcloud run deploy ucp-chat-service-bap7 --source ./a2a/chat-client --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION \
    --port 8080 --env-vars-file .env --min-instances 0 --allow-unauthenticated

  SERVICE_URL=$(gcloud run services describe ucp-chat-service-bap7 --format 'value(status.url)' --region $GOOGLE_CLOUD_LOCATION --project $GOOGLE_CLOUD_PROJECT)
fi

# set url in env file
sed -i "s,^export CHAT_URL=.*,export CHAT_URL=$SERVICE_URL," .env

gcloud run deploy ucp-chat-service-bap7 --source ./a2a/chat-client --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION \
  --port 8080 --env-vars-file .env --min-instances 0 --allow-unauthenticated
