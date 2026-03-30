gcloud run deploy ucpdemo-nutritionservice --source ./mcp/nutrition_service --project $GOOGLE_CLOUD_PROJECT \
  --region $GOOGLE_CLOUD_LOCATION --no-allow-unauthenticated --env-vars-file .env

apigeecli apis create bundle -f ./apigee/UCP-MCP-NutritionService/apiproxy --name UCP-MCP-NutritionService -o $GOOGLE_CLOUD_PROJECT \
  -e $APIGEE_ENV -s $SA_EMAIL --ovr --default-token
