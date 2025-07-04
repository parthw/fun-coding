let userId = null;

window.onload = () => {
  userId = prompt("Enter your User ID:");
  if (!userId || !userId.trim()) {
    alert("User ID is required to continue.");
    location.reload();
  }
  document
    .getElementById("user-input")
    .addEventListener("keydown", function (event) {
      if (
        event.key === "Enter" &&
        !event.shiftKey &&
        !event.ctrlKey &&
        !event.altKey &&
        !event.metaKey
      ) {
        sendMessage();
      }
    });
};

async function sendMessage() {
  const inputField = document.getElementById("user-input");
  const message = inputField.value;
  if (!message.trim()) return;
  inputField.value = "";
  const chatBox = document.getElementById("chat-box");
  chatBox.innerHTML += `<div><strong>You:</strong> ${message}</div>`;
  const response = await fetch("/ui/chat", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      message: message,
      userId: userId,
    }),
  });

  const data = await response.json();
  chatBox.innerHTML += `<div><strong>Odin:</strong> ${data.response}</div>`;
  chatBox.scrollTop = chatBox.scrollHeight;
}
