import Marionette from "backbone.marionette";
import header from "src/templates/header.jst";

export default Marionette.View.extend({
  template: header,
  ui: {
    title: "h1",
  },
});
