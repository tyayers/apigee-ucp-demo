## Apigee UCP, A2A & MCP Demo
This is demo of putting Apigee proxies with authn/authz, quotas & security policies in front of Unified Commerce Protocol (UCP), Agent2Agent (A2A), REST & Model Context Protocol (MCP) services. It is based on the [UCP Samples demo](https://github.com/Universal-Commerce-Protocol/samples), with additional REST, MCP & API proxies added.

## Test a deployed client
Visit https://apigee-ucp.agenticplatform.dev to test.

## Deploy to Google Cloud
To run this demo, you will need a Google Cloud Project with the services Apigee, Cloud Run & Model Armor active and provisioned. A service account (SA_EMAIL) is also needed with the roles **roles/run.invoker**, **roles/aiplatform.user**, and **roles/modelarmor.user**. A **Google AI API key** is needed for the business A2A service.

```sh
# Step 1: create your .env environment variables file. APIGEE_API_KEY can be filled in later after deploying the proxies.
cat > .env <<EOF

export GOOGLE_API_KEY=YOUR_GOOGLE_API_KEY
export MODEL=gemini-2.5-flash
export GOOGLE_CLOUD_PROJECT=YOUR_PROJECT_ID
export GOOGLE_CLOUD_LOCATION=YOUR_GCP_REGION
export SA_EMAIL=YOUR_SA_EMAIL
export APIGEE_ENV=YOUR_APIGEE_ENV
export APIGEE_API_KEY=YOUR_APIGEE_KEY

EOF

# deploy the nutrition MCP service
./1.deploy-nutritionservice.sh
# update the resulting URL in the ./apigee/UCP-MCP-NutritionService/apiproxy/targets/target.xml file.

# deploy the ingredeients REST service
./2.deploy-ingredientsservice.sh
# update the resulting URL in the ./apigee/UCP-REST-IngredientsService/apiproxy/targets/target.xml file.

# deploy the business A2A service
./3.deploy-businessservice.sh
# update the resulting URL in the ./a2a/chat-client/server.ts file.

# deploy the chat client
./4.deploy-chatclient.sh
# open the chat client URL and search for cookies and do a checkout.

```
