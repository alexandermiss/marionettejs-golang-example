import Marionette from "backbone.marionette";
import provider_list_item from "src/templates/providers/provider_list_item.jst";

export default Marionette.View.extend({
  tagName: "tr",
  template: provider_list_item,
});
