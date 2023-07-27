
const mailingList = document.querySelector("#mailing-list");

async function loadMailingList() {
  console.log("fetching mailing list members");

  const response = await fetch("/usuarios");
  const json = await response.json();

  json.forEach((member) => {
    mailingList.innerHTML += `<li>${member.nome} (<a href="mailto:${member.email}">${member.email}</a>)</li>`;
  });
}

document.addEventListener("DOMContentLoaded", function () {
  loadMailingList();
});
