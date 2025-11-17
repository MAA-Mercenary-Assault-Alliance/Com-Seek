// cypress.config.js
export default {
  e2e: {
    baseUrl: 'http://localhost:5173', // your URL
    supportFile: 'cypress/support/e2e.js',
    viewportWidth: 1920,
    viewportHeight: 1080,
    defaultCommandTimeout: 10000,  // <-- set here
  }
}