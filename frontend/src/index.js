import "./styles/application.css";

import Marionette from "backbone.marionette";
import AppView from "src/components/app/AppView";

document.addEventListener("DOMContentLoaded", () => {
  const app = new AppView();
  app.start();
});
