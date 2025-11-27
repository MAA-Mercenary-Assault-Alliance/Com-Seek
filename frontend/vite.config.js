import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    headers: {
      "X-Frame-Options": "DENY",
      "X-Content-Type-Options": "nosniff",
      "Referrer-Policy": "strict-origin-when-cross-origin",
      "Content-Security-Policy":
        "default-src 'self';" +
        "script-src 'self' https://www.google.com https://www.gstatic.com; " +
        "style-src 'self' https://fonts.googleapis.com 'unsafe-inline'; " +
        "font-src 'self' https://fonts.gstatic.com; " +
        "img-src 'self' data: https://i.pravatar.cc blob: http://localhost:8000;" +
        "connect-src 'self' http://localhost:8000 https://www.google.com;" + // Use 'self' for dev, or customize as needed
        "frame-src 'self' https://www.google.com;" +
        "frame-ancestors 'none';",
      "Cross-Origin-Opener-Policy": "same-origin",
      "Cross-Origin-Resource-Policy": "same-site",
    },
    proxy: {
      "/auth": "http://localhost:8000",
      "/company": "http://localhost:8000",
      "/job": "http://localhost:8000",
    },
  },
});
