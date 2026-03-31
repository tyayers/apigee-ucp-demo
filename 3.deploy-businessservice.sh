# deploy cloud run
gcloud run deploy ucpdemo-businessservice --source ./a2a/business_agent --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION \
  --env-vars-file .env --min-instances 0 --port 10999 --no-allow-unauthenticated

# deploy apigee proxy
apigeecli apis create bundle -f ./apigee/UCP-A2A-BusinessService/apiproxy --name UCP-A2A-BusinessService -o $GOOGLE_CLOUD_PROJECT \
  -e $APIGEE_ENV -s $SA_EMAIL --ovr --default-token
