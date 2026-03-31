import express from "express";

const app = express();
// app.use(cors());

app.use(express.static("dist"));
app.use(express.json());

app.post("/api", async function (request, response) {
  const targetUrl =
    process.env.TARGET_URL ||
    "https://api.apigee-bap7.agenticplatform.dev/businessservice";

  const headers = { ...request.headers } as Record<string, string>;
  delete headers.host;
  delete headers.connection;
  delete headers["content-length"];

  headers["x-api-key"] = process.env["APIGEE_API_KEY"];

  try {
    const res = await fetch(targetUrl, {
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

app.listen("8080", () => {
  console.log(`app listening on port 8080`);
});
