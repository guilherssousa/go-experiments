const messageTitle = document.querySelector("#message-title");
const messageDescription = document.querySelector("#message-description");
const contentElement = document.querySelector("#content");

const { Marked } = globalThis.marked;
const { markedHighlight } = globalThis.markedHighlight;

async function main() {
  console.log("fetching message data...");

  const highlighter = await shiki.getHighlighter({
    theme: "monokai",
    langs: ["js", "sql"],
  });

  const marked = new Marked(
    markedHighlight({
      langPrefix: "code-block language-",
      highlight(code, lang) {
        return highlighter.codeToHtml(code, { lang });
      },
    }),
    {
      mangle: false,
      headerIds: false,
    }
  );

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
  
  console.log("parsing and highlighting code....");
  console.time("parser");
  contentElement.innerHTML = marked.parse(json.content);
  console.timeEnd("parser");
}

main();
