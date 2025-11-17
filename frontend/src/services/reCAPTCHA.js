export function renderRecaptcha(element, sitekey) {
  return new Promise((resolve, reject) => {
    if (typeof grecaptcha === "undefined") {
      reject(new Error("reCAPTCHA script not loaded"));
    }

    // eslint-disable-next-line no-undef
    grecaptcha.render(element, {
      sitekey: sitekey,
      callback: (response) => resolve(response),
    });
  });
}
