document.addEventListener("DOMContentLoaded", function () {
  focusOnNewMessage();
  submitChatOnEnter();
  setupMenu();
});

function focusOnNewMessage() {
  document.addEventListener("htmx:afterSwap", function (e) {
    if (e.target.id === "chat-window") {
      e.target.scrollTop = e.target.scrollHeight;
    }
  });
}

function submitChatOnEnter() {
  document
    .getElementById("app-chat-input")
    .addEventListener("keydown", function (e) {
      if (e.key === "Enter" && !e.shiftKey) {
        e.preventDefault();
        document.getElementById("app-chat-submit-btn").click();
      }
    });
}

function onChatSubmit(el) {
  el = document.getElementById("app-chat-submit-btn");
  el.setAttribute("hx-swap", "beforeend");
  document.getElementById("app-chat-input").value = "";
  document.querySelector(".quick-bar-container").classList.add("show");
}

function setupMenu() {
  const menu = document.getElementById("app-menu");
  if (window.innerWidth >= 1280) {
    menu.setAttribute("popover", "manual");
    menu.showPopover();
  } else {
    menu.setAttribute("popover", "auto");
  }
}
