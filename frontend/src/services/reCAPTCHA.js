export function renderRecaptcha(element, sitekey) {
  return new Promise((resolve, reject) => {
    /* eslint-disable no-undef */
    if (
      typeof grecaptcha === "undefined" &&
      typeof grecaptcha.render !== "function"
    ) {
      reject(new Error("reCAPTCHA script not loaded"));
    }
    /* eslint-enable no-undef */

    if (!element.hasChildNodes()) {
      // eslint-disable-next-line no-undef
      grecaptcha.render(element, {
        sitekey: sitekey,
        callback: (response) => resolve(response),
      });
    }
  });
}
