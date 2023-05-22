import Marionette from "backbone.marionette";
import client_list_item from "src/templates/clients/client_list_item.jst";

export default Marionette.View.extend({
  tagName: "tr",
  template: client_list_item,
});
