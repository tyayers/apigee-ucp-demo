/*
 * Copyright 2026 UCP Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import react from "@vitejs/plugin-react";
import path from "node:path";
import { defineConfig } from "vite";

export default defineConfig(() => {
  return {
    server: {
      port: 3000,
      host: "0.0.0.0",
      allowedHosts: [
        "apigee-ucp.agenticplatform.dev",
        "ucpdemo-chat-609874082793.europe-west1.run.app",
        "localhost",
      ],
      proxy: {
        "/api": {
          target: "https://api.apigee-bap7.agenticplatform.dev/businessservice", // https://dev.34-107-170-170.nip.io/agent_demo - https://a2a-business-agent-bap7-323709580283.europe-west1.run.app - http://localhost:10999 or http://35.210.153.88:10999
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ""),
          secure: false,
        },
      },
    },
    plugins: [react()],
    define: {},
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "."),
      },
    },
  };
});
