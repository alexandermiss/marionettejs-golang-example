import Marionette from "backbone.marionette";
import client_detail from "src/templates/clients/client_detail.jst";

export default Marionette.View.extend({
  template: client_detail,
  region: "main",
});
