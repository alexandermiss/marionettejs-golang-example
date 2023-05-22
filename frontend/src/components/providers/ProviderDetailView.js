import Marionette from "backbone.marionette";
import provider_detail from "src/templates/providers/provider_detail.jst";

export default Marionette.View.extend({
  template: provider_detail,
  region: "main",
});
