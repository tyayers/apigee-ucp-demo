apigeecli apis create bundle -f ./apigee/UCP-REST-GeminiService/apiproxy --name UCP-REST-GeminiService -o $GOOGLE_CLOUD_PROJECT \
  -e $APIGEE_ENV -s $SA_EMAIL --ovr --default-token
