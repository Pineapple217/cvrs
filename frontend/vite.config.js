import { defineConfig } from "vite";
import preact from "@preact/preset-vite";

export default defineConfig(({ command }) => ({
  plugins: [preact()],
  define: {
    __BACKEND_URL__: JSON.stringify(
      command === "serve" ? "http://localhost:3000/api" : "/api"
    ),
    __JWT_LOCALSTORAGE__: JSON.stringify("cvrs_auth_token"),
  },
}));
