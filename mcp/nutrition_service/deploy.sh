gcloud run deploy nutrition-service --source . --project $GOOGLE_CLOUD_PROJECT --region $GOOGLE_CLOUD_LOCATION --no-allow-unauthenticated --env-vars-file .env
