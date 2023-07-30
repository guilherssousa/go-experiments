const messageTitle = document.querySelector("#message-title");
const messageDescription = document.querySelector("#message-description");
const contentElement = document.querySelector("#content");

async function main() {
  console.log("fetching message data...");

  const { searchParams } = new URL(window.location.href);
  const id = searchParams.get("id");  

  if (!id) {
    alert("Nao e possivel achar a mensagem.");
    window.location.navigate("/");
  }

  const request = await fetch(`/messages/${id}`);
  const json = await request.json();

  messageTitle.textContent = json.title;
  messageDescription.innerHTML = `
    por ${json.nome} \<<a href="mailto:${json.email}">${json.email}</a>\>
  `;

  contentElement.innerHTML = marked.parse(json.content); 
}

main();
