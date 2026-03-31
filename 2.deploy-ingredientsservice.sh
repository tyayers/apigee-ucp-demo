gcloud run deploy ucpdemo-ingredients-service --source ./rest/ingredients_service --project $GOOGLE_CLOUD_PROJECT \
  --region $GOOGLE_CLOUD_LOCATION --no-allow-unauthenticated --env-vars-file .env

apigeecli apis create bundle -f ./apigee/UCP-REST-IngredientsService/apiproxy --name UCP-REST-IngredientsService -o $GOOGLE_CLOUD_PROJECT \
  -e $APIGEE_ENV -s $SA_EMAIL --ovr --default-token

apigeecli apis create bundle -f ./apigee/UCP-MCP-IngredientsService/apiproxy --name UCP-MCP-IngredientsService -o $GOOGLE_CLOUD_PROJECT \
  -e $APIGEE_ENV -s $SA_EMAIL --ovr --default-token
