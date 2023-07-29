
const mailingList = document.querySelector("#mailing-list");
const messageList = document.querySelector("#message-list");

async function loadMessages() {
  console.log("fetching messages");
  
  const response = await fetch("/messages");
  const json = await response.json();

  json.forEach((message) => {
    messageList.innerHTML += `
      <li><a href="/message/${message.id}">"${message.title}"</a> por ${message.nome} \<<a href="mailto:${message.email}">${message.email}</a>\></li>
    `;
  });
}

async function loadMailingList() {
  console.log("fetching mailing list members");

  const response = await fetch("/usuarios");
  const json = await response.json();

  json.forEach((member) => {
    mailingList.innerHTML += `<li>${member.nome} (<a href="mailto:${member.email}">${member.email}</a>)</li>`;
  });
}

document.addEventListener("DOMContentLoaded", function () {
  loadMessages();
  loadMailingList();
});
