import express from "express";
import profile_json from "./profile/agent_profile.json" with { type: "json" };

// do a ping to the businessservice proxy to wake it up from cloud run sleep :)
fetch(
  (process.env.TARGET_URL ||
    "https://api.apigee-bap7.agenticplatform.dev/businessservice") +
    "/.well-known/agent-card.json",
);

const app = express();
// app.use(cors());

app.use(express.static("dist"));
app.use(express.json());

app.post("/api", async function (request, response) {
  const businessServiceUrl =
    process.env.TARGET_URL ||
    "https://api.apigee-bap7.agenticplatform.dev/businessservice";

  const headers = { ...request.headers } as Record<string, string>;
  delete headers.host;
  delete headers.connection;
  delete headers["content-length"];

  headers["x-api-key"] =
    headers["x-mode"] == "freemium"
      ? process.env["APIGEE_API_KEY"]
      : process.env["APIGEE_PRO_API_KEY"];
  // headers["ucp-agent"] =
  //   'profile="' + process.env["CHAT_URL"] + '/profile/agent_profile.json"';

  try {
    const res = await fetch(businessServiceUrl, {
      method: "POST",
      headers,
      body: JSON.stringify(request.body),
    });

    res.headers.forEach((value, key) => {
      response.setHeader(key, value);
    });

    const data = await res.text();
    response.status(res.status).send(data);
  } catch (error) {
    console.error("Error forwarding request:", error);
    response.status(500).json({ error: "Failed to forward request" });
  }
});

app.get("/profile/agent_profile.json", async function (request, response) {
  response.json(profile_json);
});

app.listen("8080", () => {
  console.log(`app listening on port 8080`);
});
