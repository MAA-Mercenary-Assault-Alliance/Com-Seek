export function renderRecaptcha(element, sitekey) {
  return new Promise((resolve, reject) => {
    // eslint-disable-next-line no-undef
    if (typeof grecaptcha === "undefined" && !grecaptcha.render) {
      reject(new Error("reCAPTCHA script not loaded"));
    }

    if (!element.hasChildNodes()) {
      // eslint-disable-next-line no-undef
      grecaptcha.render(element, {
        sitekey: sitekey,
        callback: (response) => resolve(response),
      });
    }
  });
}
