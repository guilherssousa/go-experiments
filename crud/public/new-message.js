function submitMessage(e) {
  e.preventDefault()
  const messageForm = document.querySelector("#message-form");
  const formData = new FormData(messageForm);
  
  const readable = [...formData.entries()].reduce((o, [k,v]) => {
    o[k] = v
    return o
  }, {})

  console.log(readable)
}
