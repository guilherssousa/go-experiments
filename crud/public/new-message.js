async function submitMessage(e) {
  e.preventDefault();
  const messageForm = document.querySelector("#message-form");
  const formData = new FormData(messageForm);

  const formObject = [...formData.entries()].reduce((o, [k, v]) => {
    o[k] = v;
    return o;
  }, {});

  if (formObject.title.length < 3 || formObject.content.length < 8) {
    alert("Por favor, preencha o formulario.");
  }

  const request = await fetch("/messages", {
    headers: {
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify(formObject),
  });

  if (request.status === 201) {
    window.location.replace("/");
  } else {
    alert("Request falhou.");
  }
}
