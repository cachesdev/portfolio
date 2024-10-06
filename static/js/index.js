// Check de tema para evitar FOUC
// if (
//     localStorage.theme === "night" ||
//     (!("theme" in localStorage) &&
//         window.matchMedia("(prefers-color-scheme: dark)").matches)
// ) {
//     document.documentElement.setAttribute("data-theme", "night");
// } else {
//     document.documentElement.setAttribute("data-theme", "winter");
// }

function monitorEvents(element) {
  var log = function (e) {
    console.log(e);
  };
  var events = [];

  for (var i in element) {
    if (i.startsWith("on")) events.push(i.substr(2));
  }
  events.forEach(function (eventName) {
    element.addEventListener(eventName, log);
  });
}

function addEventListeners(element, events, handler) {
  events.forEach((event) => element.addEventListener(event, handler));
}

// surreal
function pluginClassToggleList(e) {
  function classToggleList(e, classes, force) {
    if (typeof classes !== "string") return e;

    const classList = classes.split(" ");
    if (surreal.isNodeList(e))
      e.forEach((_) => {
        classToggleList(_, classes, force);
      });
    if (surreal.isNode(e)) {
      classList.forEach((className) => {
        e.classList.toggle(className.trim(), force);
      });
    }
    return e; // Make chainable.
  }
  // Add sugar
  e.classToggleList = (classes, force) => {
    return classToggleList(e, classes, force);
  };
}

surreal.plugins.push(pluginClassToggleList);
